package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	mysqlwrapper "openai-svr/mysql_wrapper"
	wrapper "openai-svr/openai_wrapper"

	"github.com/sashabaranov/go-openai"
	"k8s.io/klog/v2"
)

type chatCompletionReq struct {
	ChatID  string `json:"chat_id"`
	UserUID string `json:"user_uid"`
	Prompt  string `json:"prompt"`
}

type chatCompletionData struct {
	ID     string `json:"id"`
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func handleChatChatCompletion(http_resp http.ResponseWriter, http_req *http.Request) {
	resp := UniversalResp{}
	defer response(&resp, http_resp)

	buf, err := ioutil.ReadAll(http_req.Body)
	defer http_req.Body.Close()
	if err != nil {
		http_resp.WriteHeader(http.StatusBadGateway)
		resp.QuickSetup(NetworkError, fmt.Sprintf("ReadAll error: %s", err.Error()))
		return
	}

	http_resp.WriteHeader(http.StatusOK)

	internal_req := chatCompletionReq{}
	err = json.Unmarshal(buf, &internal_req)
	if err != nil {
		resp.QuickSetup(UnmarshalJsonError, fmt.Sprintf("Unmarshal error: %s", err.Error()))
		return
	}

	if len(internal_req.ChatID) == 0 || len(internal_req.Prompt) == 0 || len(internal_req.UserUID) == 0 {
		resp.QuickSetup(ParamaterError, fmt.Sprintf("Parameter error: %s", "some field miss"))
		return
	}

	uuid, err := mysqlwrapper.GetUserUIDByChatID(internal_req.ChatID)
	if err != nil {
		if err == sql.ErrNoRows {
			resp.QuickSetup(ParamaterError, fmt.Sprintf("invalid request: chat id not found, please create new chat first"))
			return
		}
		resp.QuickSetup(DBError, fmt.Sprintf("internal error, try again later"))
		return
	}
	if uuid != internal_req.UserUID {
		resp.QuickSetup(NoAuthError, fmt.Sprintf("UserUUID not consistent with chatID"))
		return
	}

	historyMessages, err := mysqlwrapper.GetHistoryMessageByChatID(internal_req.ChatID)
	if err != nil {
		klog.Errorf("GetHistoryMessageByChatID failed: %s", err.Error())
		resp.QuickSetup(DBError, fmt.Sprintf("access histroy failed"))
		return
	}

	openai_resp, err := wrapper.ChatCompletion(historyMessages, internal_req.Prompt)
	if err != nil {
		resp.QuickSetup(OpenAIError, fmt.Sprintf("OpenAIError error: %s", err.Error()))
		return
	}
	resp.Data = chatCompletionData{
		ID:   openai_resp.ID,
		Text: openai_resp.Choices[0].Message.Content,
	}

	// TODO: record user usage in DB

	// TODO:: record user prompt and open ai answer in DB
	newMessages := make([]openai.ChatCompletionMessage, 2, 2)
	newMessages[0] = openai.ChatCompletionMessage{
		Role:    "user",
		Content: internal_req.Prompt,
	}
	newMessages[1] = openai_resp.Choices[0].Message
	err = mysqlwrapper.UpdateHistoryMessageByChatID(internal_req.ChatID, newMessages)
	if err != nil {
		klog.Errorf("UpdateHistroyMessageByChatID failed: %s", err.Error())
	}

	resp.QuickSetup(Ok, "ok")
	return
}

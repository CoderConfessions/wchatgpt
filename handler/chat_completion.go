package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	wrapper "openai-svr/openai_wrapper"

	"github.com/sashabaranov/go-openai"
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
	if err != nil {
		http_resp.WriteHeader(http.StatusBadGateway)
		resp.QuickSetup(NetworkError, fmt.Sprintf("ReadAll error: %s", err.Error()))
		return
	}

	internal_req := chatCompletionReq{}
	err = json.Unmarshal(buf, &internal_req)
	if err != nil {
		http_resp.WriteHeader(http.StatusBadRequest)
		resp.QuickSetup(UnmarshalJsonError, fmt.Sprintf("Unmarshal error: %s", err.Error()))
		return
	}

	if len(internal_req.ChatID) == 0 || len(internal_req.Prompt) == 0 || len(internal_req.UserUID) == 0 {
		http_resp.WriteHeader(http.StatusBadRequest)
		resp.QuickSetup(ParamaterError, fmt.Sprintf("Parameter error: %s", "some field miss"))
		return
	}
	// TODO: access ChatID , error if not exists
	// TODO: acccess history messages in DB, join here ?

	// TODO: validate chatID has expired or not, error if expired

	historyMessages := make([]openai.ChatCompletionMessage, 0)
	openai_resp, err := wrapper.ChatCompletion(historyMessages, internal_req.Prompt)
	if err != nil {
		http_resp.WriteHeader(http.StatusInternalServerError)
		resp.QuickSetup(OpenAIError, fmt.Sprintf("OpenAIError error: %s", err.Error()))
		return
	}
	resp.Data = chatCompletionData{
		ID:   openai_resp.ID,
		Text: openai_resp.Choices[0].Message.Content,
	}

	// TODO: record user usage in DB

	// TODO:: record user prompt and open ai answer in DB

	resp.QuickSetup(Ok, "ok")
	return
}

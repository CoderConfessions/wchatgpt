package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	wrapper "openai-svr/openai_wrapper"

	"github.com/sashabaranov/go-openai"
)

type chatCompletionStatelessReq struct {
	Messages []openai.ChatCompletionMessage `json:"messages"`
	Prompt   string                         `json:"prompt"`
}

type chatCompletionStatelessData struct {
	Text string `json:"text"`
}

func handleChatChatCompletionStateless(http_resp http.ResponseWriter, http_req *http.Request) {
	resp := UniversalResp{}
	defer response(&resp, http_resp)

	buf, err := ioutil.ReadAll(http_req.Body)
	defer http_req.Body.Close()
	if err != nil {
		http_resp.WriteHeader(http.StatusBadGateway)
		resp.quickSetup(NetworkError, fmt.Sprintf("ReadAll error: %s", err.Error()))
		return
	}

	http_resp.WriteHeader(http.StatusOK)

	internal_req := chatCompletionStatelessReq{}
	err = json.Unmarshal(buf, &internal_req)
	if err != nil {
		resp.quickSetup(UnmarshalJsonError, fmt.Sprintf("Unmarshal error: %s", err.Error()))
		return
	}

	if len(internal_req.Prompt) == 0 {
		resp.quickSetup(ParamaterError, fmt.Sprintf("Parameter error: %s", "prompt is empty"))
		return
	}

	openai_resp, err := wrapper.ChatCompletion(internal_req.Messages, internal_req.Prompt)
	if err != nil {
		resp.quickSetup(OpenAIError, fmt.Sprintf("OpenAIError error: %s", err.Error()))
		return
	}
	resp.Data = chatCompletionStatelessData{
		Text: openai_resp.Choices[0].Message.Content,
	}

	resp.quickSetup(Ok, "ok")
	return
}

package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	wrapper "openai-svr/openai_wrapper"
)

type SingleCompletionReq struct {
	Prompt    string `json:"prompt"`
	UserToken string `json:"user_token"`
}

func HandleSingleCompletion(http_resp http.ResponseWriter, http_req *http.Request) {
	resp := UniversalResp{}
	defer func() {
		body, _ := json.Marshal(resp)
		http_resp.Header().Set("Content-Type", "application/json")
		http_resp.Write(body)
	}()

	buf, err := ioutil.ReadAll(http_req.Body)
	if err != nil {
		http_resp.WriteHeader(http.StatusBadGateway)
		resp.QuickSetup(NetworkError, fmt.Sprintf("ReadAll error: %s", err.Error()))
		return
	}

	internal_req := SingleCompletionReq{}
	err = json.Unmarshal(buf, &internal_req)
	if err != nil {
		http_resp.WriteHeader(http.StatusBadRequest)
		resp.QuickSetup(UnmarshalJsonError, fmt.Sprintf("Unmarshal error: %s", err.Error()))
		return
	}

	if len(internal_req.Prompt) == 0 || len(internal_req.UserToken) == 0 {
		http_resp.WriteHeader(http.StatusBadRequest)
		resp.QuickSetup(ParamaterError, fmt.Sprintf("Parameter error: %s", "some field miss"))
		return
	}

	if internal_req.UserToken != "xilan" {
		http_resp.WriteHeader(http.StatusUnauthorized)
		resp.QuickSetup(NoAuthError, fmt.Sprintf("NoAuth error: token is invalid"))
		return
	}

	if false {
		openai_resp, err := wrapper.SingleCompletion(internal_req.Prompt)
		if err != nil {
			http_resp.WriteHeader(http.StatusInternalServerError)
			resp.QuickSetup(OpenAIError, fmt.Sprintf("OpenAIError error: %s", err.Error()))
			return
		}
		resp.Data = openai_resp
	}

	resp.QuickSetup(Ok, "ok")
	return
}

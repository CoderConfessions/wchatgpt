package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	wrapper "openai-svr/openai_wrapper"

	"k8s.io/klog/v2"
)

type singleCompletionReq struct {
	Prompt  string `json:"prompt"`
	UserUID string `json:"user_uid"`
}

type singleCompletionData struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func handleSingleCompletion(http_resp http.ResponseWriter, http_req *http.Request) {
	resp := UniversalResp{}
	defer response(&resp, http_resp)

	buf, err := ioutil.ReadAll(http_req.Body)
	defer http_req.Body.Close()
	if err != nil {
		http_resp.WriteHeader(http.StatusBadGateway)
		resp.QuickSetup(NetworkError, fmt.Sprintf("ReadAll error: %s", err.Error()))
		return
	}
	klog.V(8).Infof("read request ok")
	http_resp.WriteHeader(http.StatusOK)

	internal_req := singleCompletionReq{}
	err = json.Unmarshal(buf, &internal_req)
	if err != nil {
		resp.QuickSetup(UnmarshalJsonError, fmt.Sprintf("Unmarshal error: %s", err.Error()))
		return
	}

	if len(internal_req.Prompt) == 0 || len(internal_req.UserUID) == 0 {
		resp.QuickSetup(ParamaterError, fmt.Sprintf("Parameter error: %s", "some field miss"))
		return
	}
	// TODO: access UserUID in db, insert if not exists

	openai_resp, err := wrapper.SingleCompletion(internal_req.Prompt)
	if err != nil {
		klog.Errorf("OpenAIError error: %v", err.Error())
		resp.QuickSetup(OpenAIError, fmt.Sprintf("OpenAIError error: %s", err.Error()))
		return
	}
	resp.Data = singleCompletionData{
		ID:   openai_resp.ID,
		Text: openai_resp.Choices[0].Text,
	}

	// TODO: record user usage in DB

	resp.QuickSetup(Ok, "ok")
	return
}

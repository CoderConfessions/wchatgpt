package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	mysqlwrapper "openai-svr/mysql_wrapper"

	"github.com/google/uuid"
	"k8s.io/klog/v2"
)

type createChatIdReq struct {
	UserUID string `json:"user_uid"`
}

type createChatIdData struct {
	ChatID string `json:"chat_id"`
}

func createChatId(http_resp http.ResponseWriter, http_req *http.Request) {
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

	internal_req := createChatIdReq{}
	err = json.Unmarshal(buf, &internal_req)
	if err != nil {
		resp.QuickSetup(UnmarshalJsonError, fmt.Sprintf("Unmarshal error: %s", err.Error()))
		return
	}

	if internal_req.UserUID == "" {
		resp.QuickSetup(ParamaterError, fmt.Sprintf("Parameter error: %s", "some field miss"))
		return
	}
	chatID := uuid.New().String()
	resp.Data = createChatIdData{ChatID: chatID}

	err = mysqlwrapper.UpdateChatID(internal_req.UserUID, chatID)
	if err != nil {
		klog.Errorf("UpdateChatID failed: %s", err.Error())
		resp.QuickSetup(DBError, fmt.Sprintf("Internal error, try again later"))
		return
	}

	resp.QuickSetup(Ok, "ok")
	return
}

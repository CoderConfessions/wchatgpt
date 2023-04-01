package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
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
	if err != nil {
		http_resp.WriteHeader(http.StatusBadGateway)
		resp.QuickSetup(NetworkError, fmt.Sprintf("ReadAll error: %s", err.Error()))
		return
	}

	internal_req := createChatIdReq{}
	err = json.Unmarshal(buf, &internal_req)
	if err != nil {
		http_resp.WriteHeader(http.StatusBadRequest)
		resp.QuickSetup(UnmarshalJsonError, fmt.Sprintf("Unmarshal error: %s", err.Error()))
		return
	}

	if internal_req.UserUID == "" {
		http_resp.WriteHeader(http.StatusBadRequest)
		resp.QuickSetup(ParamaterError, fmt.Sprintf("Parameter error: %s", "some field miss"))
		return
	}
	chatID := uuid.New().String()
	resp.Data = createChatIdData{ChatID: chatID}

	// TODO: write db record user uid in chat id

	resp.QuickSetup(Ok, "ok")
	return
}

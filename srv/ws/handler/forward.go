package handler

import (
	"bytes"
	"encoding/json"
	"gateway/constants"
	"io/ioutil"
	"log"
	"net/http"
)

type forward struct {
}

var Forwarding *forward

func init() {
	Forwarding = &forward{}
}

func (receiver forward) Send(msg []byte) {
	host := receiver.getAppUrl(msg)
	resp, err := http.Post(host, "application/json", bytes.NewReader(msg))

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	response := &responseGenerated{}
	if json.Unmarshal(body, response) != nil {
		log.Println(string(body))
	}

	if response.Code != 0 {
		log.Println(response.Code)
		log.Println(response.Msg)
	}
}

func (receiver *forward) getAppUrl(msg []byte) string {
	msgStruct := &msgGenerated{}
	err := json.Unmarshal(msg, msgStruct)
	if err != nil {
		log.Println(err.Error())
	}

	if msgStruct.AppID == 0 {
		msgStruct.AppID = 1
	}

	return constants.GetAppAgent(msgStruct.AppID)
}

type msgGenerated struct {
	AppID int `json:"app_id"`
}

type responseGenerated struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

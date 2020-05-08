package net

import (
	"encoding/json"
)

var netDataBufChan chan []byte
var msgMap map[string]chan []byte

type tmpField struct {
	MsgParam string
}

//申请一个消息字典，用于消息解析
func CreateMsgField(field string, ch chan []byte) {
	msgMap[field] = ch
}

//按api规定的格式解析
func waitNetData() {

	for {
		buf := <-netDataBufChan
		tmpField := &tmpField{}

		if !json.Valid(buf) {
			log.PrintlnWarring("json.Valid return false:", string(buf))
			continue
		}

		if err := json.Unmarshal(buf, tmpField); err != nil {
			log.PrintlnWarring("err:", err, "data:", string(buf))
			continue
		}

		if neTmpChan, ok := msgMap[tmpField.MsgParam]; ok {
			neTmpChan <- buf
		} else {
			log.PrintlnErr("msgMap no find field:", tmpField.MsgParam)
		}

	}
}

func APIInit() {
	msgMap = make(map[string]chan []byte)
	netDataBufChan = make(chan []byte, 1000)

	go waitNetData()
}

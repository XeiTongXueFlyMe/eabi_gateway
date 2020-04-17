package main

import (
	"encoding/json"
	"fmt"
)

var netDataBufChan chan []byte
var msgMap map[string]chan []byte

type tmpField struct {
	MsgParam string
}

//申请一个消息字典，用于消息解析
func createMsgField(field string, ch chan []byte) {
	msgMap[field] = ch
}

//按api规定的格式解析
func waitNetData() {

	for {
		buf := <-netDataBufChan
		tmpField := &tmpField{}

		if !json.Valid(buf) {
			fmt.Println("json.Valid return false:", string(buf))
			continue
		}

		if err := json.Unmarshal(buf, tmpField); err != nil {
			fmt.Println(err)
			continue
		}

		if neTmpChan, ok := msgMap[tmpField.MsgParam]; ok {
			neTmpChan <- buf
		} else {
			fmt.Println("msgMap no find field:", tmpField.MsgParam)
		}

	}
}

func apiInit() {
	msgMap = make(map[string]chan []byte)
	netDataBufChan = make(chan []byte, 1000)

	go waitNetData()
}

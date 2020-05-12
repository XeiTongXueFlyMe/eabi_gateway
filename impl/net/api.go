package net

import (
	model "eabi_gateway/impl"
	"eabi_gateway/impl/config"
	"encoding/json"
	"time"
)

var netDataBufChan chan []byte
var msgMap map[string]chan []byte

type tmpField struct {
	MsgGwID  string
	MsgParam string
}

//申请一个消息字典，用于消息解析
func CreateMsgField(field string, ch chan []byte) {
	msgMap[field] = ch
}

func gwIdErr(buf []byte) {
	var msgID, msgGwID string

	m := make(map[string]interface{})
	if err := json.Unmarshal(buf, &m); err != nil {
		log.PrintlnErr(err)
	}

	for k, v := range m {
		switch k {
		case "msgId":
			if str, ok := v.(string); ok {
				msgID = str
			} else {
				log.PrintfErr("json msgId is no string")
			}
		case "msgGwId":
			if str, ok := v.(string); ok {
				msgGwID = str
			} else {
				log.PrintfErr("json msgId is no string")
			}
		}
	}

	param := &model.StdResp{
		MsgType:      "PUT",
		MsgID:        msgID,
		MsgGwID:      config.SysParamGwId(),
		MsgTimeStamp: time.Now().Unix(),
		MsgParam:     msgGwID,
		MsgResp:      "errorIdentity",
	}
	if buf, err := json.Marshal(param); err != nil {
		log.PrintlnErr(err)
	} else {
		SendData(buf)
	}
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

		if tmpField.MsgGwID != config.SysParamGwId() {
			gwIdErr(buf)
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

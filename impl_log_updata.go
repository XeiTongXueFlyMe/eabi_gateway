package main

import (
	config "eabi_gateway/impl/config"
	net "eabi_gateway/impl/net"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var AUTO_FEEDING_HISTORY_DATA_FILE_NAME string = "autoFeedingAlgoHistoryData.log"

var logParamChannel chan []byte
var feedingLogParamChannel chan []byte

type logResp struct {
	MsgType      string `json:"msgType"`
	MsgID        string `json:"msgId"`
	MsgGwID      string `json:"msgGwId"`
	MsgTimeStamp int64  `json:"msgTimeStamp"`
	MsgParam     string `json:"msgParam"`
	MsgResp      string `json:"msgResp"`
	LogData      string `json:"logData"`
}

func updataLogFileToServer(msgID interface{}) {
	param := &logResp{
		MsgType:      "GET",
		MsgGwID:      config.SysParamGwId(),
		MsgTimeStamp: time.Now().Unix(),
		MsgParam:     "eabiLog",
		MsgResp:      "ok",
	}

	year, month, _ := time.Now().Date()

	fn := fmt.Sprintf("%d-%d.log", year, month)
	f, err := os.OpenFile(fn, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fstat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fstat.Size())

	n, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	param.LogData = string(buf[:n])

	if id, ok := msgID.(string); ok {
		param.MsgID = id
	} else {
		return
	}

	b, err := json.Marshal(param)
	if err != nil {
		log.PrintlnErr(err)
		goto _exit
	}
	if _, err := net.SendData(b); err != nil {
		log.PrintlnErr(err)
		goto _exit
	}

_exit:
	return

}

func deleteLofFile() {
	year, month, _ := time.Now().Date()
	fn := fmt.Sprintf("%d-%d.log", year, month)

	os.Remove(fn)
}

func waitLogParamReq() {

	for {
		buf := <-logParamChannel

		m := make(map[string]interface{})
		if err := json.Unmarshal(buf, &m); err != nil {
			log.PrintlnErr(err)
			respToServer(m["msgId"], err.Error(), "eabiLog")
			continue
		}
		if v, ok := m["msgType"]; ok {
			if str, ok := v.(string); ok {
				switch str {
				case "GET":
					updataLogFileToServer(m["msgId"])
				case "DELETE":
					deleteLofFile()
					respDeleteToServer(m["msgId"], "ok", "eabiLog")
				}
			} else {
				log.PrintfErr("json msgType no is string")
			}
		} else {
			log.PrintlnErr("no find msgType")
		}
	}
}

func algoLogFileToServer(msgID interface{}) {
	param := &logResp{
		MsgType:      "GET",
		MsgGwID:      config.SysParamGwId(),
		MsgTimeStamp: time.Now().Unix(),
		MsgParam:     "feedingAlgo",
		MsgResp:      "ok",
	}

	f, err := os.OpenFile(AUTO_FEEDING_HISTORY_DATA_FILE_NAME, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fstat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fstat.Size())

	n, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	param.LogData = string(buf[:n])

	if id, ok := msgID.(string); ok {
		param.MsgID = id
	} else {
		return
	}

	b, err := json.Marshal(param)
	if err != nil {
		log.PrintlnErr(err)
		goto _exit
	}
	if _, err := net.SendData(b); err != nil {
		log.PrintlnErr(err)
		goto _exit
	}

_exit:
	return

}

func waitAlgoLogParamReq() {

	for {
		buf := <-feedingLogParamChannel

		m := make(map[string]interface{})
		if err := json.Unmarshal(buf, &m); err != nil {
			log.PrintlnErr(err)
			respToServer(m["msgId"], err.Error(), "feedingAlgo")
			continue
		}
		if v, ok := m["msgType"]; ok {
			if str, ok := v.(string); ok {
				switch str {
				case "GET":
					algoLogFileToServer(m["msgId"])
				}
			} else {
				log.PrintfErr("json msgType no is string")
			}
		} else {
			log.PrintlnErr("no find msgType")
		}
	}
}

func logUpdataInit() {

	//系统日志
	logParamChannel = make(chan []byte, 1)
	net.CreateMsgField("eabiLog", logParamChannel)
	go waitLogParamReq()

	//算法日志
	feedingLogParamChannel = make(chan []byte, 1)
	net.CreateMsgField("feedingAlgo", feedingLogParamChannel)
	go waitAlgoLogParamReq()
}

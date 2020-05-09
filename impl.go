package main

import (
	modle "eabi_gateway/impl"
	"eabi_gateway/impl/config"
	"eabi_gateway/impl/net"
	rfNet "eabi_gateway/impl/rf_net"
	"encoding/json"
	"time"
)

var gatewayParamChannel chan []byte
var rfNetInfoChannel chan []byte

func waitGatewayParamConfig() {
	for {
		buf := <-gatewayParamChannel

		m := make(map[string]interface{})
		if err := json.Unmarshal(buf, &m); err != nil {
			log.PrintlnErr(err)
		}
		if v, ok := m["msgType"]; ok {
			if str, ok := v.(string); ok {
				switch str {
				case "GET":
					gwParmToServer(m)
				case "PUT":
					config.ConfigTofile(m)
				}
			} else {
				log.PrintfErr("json msgType no is string")
			}
		} else {
			log.PrintlnErr("no find msgType")
		}
	}
}

func gwParmToServer(m map[string]interface{}) {
	var msgID string

	for k, v := range m {
		switch k {
		case "msgId":
			if str, ok := v.(string); ok {
				msgID = str
			} else {
				log.PrintfErr("json msgId is no string")
			}
		}
	}

	ip, port := config.SysParamServerIPAndPort()
	rfid, rfchan, rfnetid := config.SysParamRf()

	param := &modle.GatewayParmResp{
		MsgType:      "GET",
		MsgID:        msgID,
		MsgGwID:      config.SysParamGwId(),
		MsgTimeStamp: time.Now().Unix(),
		MsgParam:     "gatewayParam",
		MsgResp:      "ok",
		GwID:         config.SysParamGwId(),
		GwIP:         config.SysParamGwIp(),
		ServerIP:     ip,
		ServerPort:   port,
		RfID:         rfid,
		RfChannel:    rfchan,
		RfNetID:      rfnetid,
		DataUpCycle:  config.SysParamDataUpCycle(),
		HeartCycle:   config.SysParamHeartCycle(),
	}

	buf, err := json.Marshal(param)
	if err != nil {
		log.PrintlnErr(err)
		goto _exit
	}
	if _, err := net.SendData(buf); err != nil {
		log.PrintlnErr(err)
		goto _exit
	}

_exit:
	return
}

func waitRfNetInfoConfig() {
	for {
		buf := <-rfNetInfoChannel

		m := make(map[string]interface{})
		if err := json.Unmarshal(buf, &m); err != nil {
			log.PrintlnErr(err)
		}
		if v, ok := m["msgType"]; ok {
			if str, ok := v.(string); ok {
				switch str {
				case "GET":
					rfNetInfoToServer(m)
				case "DELETE":
					rfNet.CleanInfo()
				}
			} else {
				log.PrintfErr("json msgType no is string")
			}
		} else {
			log.PrintlnErr("no find msgType")
		}
	}
}

func rfNetInfoToServer(m map[string]interface{}) {
	var msgID string

	for k, v := range m {
		switch k {
		case "msgId":
			if str, ok := v.(string); ok {
				msgID = str
			} else {
				log.PrintfErr("json msgId is no string")
			}
		}
	}

	param := &modle.RfNetInfoResp{
		MsgType:      "GET",
		MsgID:        msgID,
		MsgGwID:      config.SysParamGwId(),
		MsgTimeStamp: time.Now().Unix(),
		MsgParam:     "rfNetInfo",
		MsgResp:      "ok",
	}

	num := 0
	rfNetList := rfNet.ReadInfo()
	for _, v := range rfNetList {
		num++
		param.RfNetInfo = append(param.RfNetInfo, v)
	}
	param.RfNetNum = num

	buf, err := json.Marshal(param)
	if err != nil {
		log.PrintlnErr(err)
		goto _exit
	}
	if _, err := net.SendData(buf); err != nil {
		log.PrintlnErr(err)
		goto _exit
	}

_exit:
	return
}

func implInit() {
	gatewayParamChannel = make(chan []byte, 1)
	net.CreateMsgField("gatewayParam", gatewayParamChannel)
	go waitGatewayParamConfig()

	rfNetInfoChannel = make(chan []byte, 1)
	net.CreateMsgField("rfNetInfo", rfNetInfoChannel)
	go waitRfNetInfoConfig()
}

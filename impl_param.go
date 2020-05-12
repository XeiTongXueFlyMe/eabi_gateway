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
var sensorInfoCfgChannel chan []byte

func waitGatewayParamConfig() {
	for {
		buf := <-gatewayParamChannel

		m := make(map[string]interface{})
		if err := json.Unmarshal(buf, &m); err != nil {
			log.PrintlnErr(err)
			continue
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
			continue
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
func sensorCfgToServer(req modle.SensorInfoReq, sInfo []modle.SensorInfo) {

	param := &modle.SensorInfoResp{
		MsgType:      "GET",
		MsgID:        req.MsgID,
		MsgGwID:      config.SysParamGwId(),
		MsgTimeStamp: time.Now().Unix(),
		MsgParam:     "sensorInfo",
		MsgResp:      "ok",
	}
	param.SensorListNum = config.ReadSensorConfigNum()
	param.SensorList = sInfo

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

func waitSensorCfgInfoConfig() {
	var sensorInfo modle.SensorInfoReq

	for {
		buf := <-sensorInfoCfgChannel

		if err := json.Unmarshal(buf, &sensorInfo); err != nil {
			log.PrintlnErr(err)
			continue
		}

		switch sensorInfo.MsgType {
		case "GET":
			sensorCfgToServer(sensorInfo, config.ReadSensorConfig())
		case "PUT":
			config.WriteSensorConfig(sensorInfo.SensorList)
		default:
			log.PrintfErr("json msgType:%s no support ", sensorInfo.MsgType)
		}
	}
}

func implInit() {
	//网关参数的增删改查
	gatewayParamChannel = make(chan []byte, 1)
	net.CreateMsgField("gatewayParam", gatewayParamChannel)
	go waitGatewayParamConfig()

	//射频网络的信息的删查
	rfNetInfoChannel = make(chan []byte, 1)
	net.CreateMsgField("rfNetInfo", rfNetInfoChannel)
	go waitRfNetInfoConfig()

	//传感器modbus配置
	//TODO:本地存储或读取配置文件，
	sensorInfoCfgChannel = make(chan []byte, 1)
	net.CreateMsgField("sensorInfo", sensorInfoCfgChannel)
	go waitSensorCfgInfoConfig()
}

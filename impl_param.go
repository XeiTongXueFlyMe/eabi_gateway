package main

import (
	modle "eabi_gateway/impl"
	"eabi_gateway/impl/config"
	"eabi_gateway/impl/net"
	rfNet "eabi_gateway/impl/rf_net"
	"encoding/json"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

var gatewayParamChannel chan []byte
var rfNetInfoChannel chan []byte
var sensorInfoCfgChannel chan []byte
var alarmInfoCfgChannel chan []byte
var AdapterInfoCfgChannel chan []byte

func waitGatewayParamConfig() {
	for {
		buf := <-gatewayParamChannel

		m := make(map[string]interface{})
		if err := json.Unmarshal(buf, &m); err != nil {
			log.PrintlnErr(err)
			respToServer(m["msgId"], err.Error(), "gatewayParam")
			continue
		}
		if v, ok := m["msgType"]; ok {
			if str, ok := v.(string); ok {
				switch str {
				case "GET":
					gwParmToServer(m)
				case "PUT":
					msgResp := config.ConfigTofile(m)
					respToServer(m["msgId"], msgResp, "gatewayParam")
				case "DELETE":
					respDeleteToServer(m["msgId"], "ok", "gatewayParam")
				}
			} else {
				log.PrintfErr("json msgType no is string")
			}
		} else {
			log.PrintlnErr("no find msgType")
		}
	}
}

func respToServer(msgID interface{}, msgResp string, msgParam string) {
	param := &modle.StdResp{
		MsgType:      "PUT",
		MsgGwID:      config.SysParamGwId(),
		MsgTimeStamp: time.Now().Unix(),
		MsgParam:     msgParam,
		MsgResp:      msgResp,
	}

	if id, ok := msgID.(string); ok {
		param.MsgID = id
	} else {
		return
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

func respDeleteToServer(msgID interface{}, msgResp string, msgParam string) {
	param := &modle.StdResp{
		MsgType:      "DELETE",
		MsgGwID:      config.SysParamGwId(),
		MsgTimeStamp: time.Now().Unix(),
		MsgParam:     msgParam,
		MsgResp:      msgResp,
	}

	if id, ok := msgID.(string); ok {
		param.MsgID = id
	} else {
		return
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
			respToServer(m["msgId"], err.Error(), "rfNetInfo")
			continue
		}
		if v, ok := m["msgType"]; ok {
			if str, ok := v.(string); ok {
				switch str {
				case "GET":
					rfNetInfoToServer(m)
				case "DELETE":
					rfNet.CleanInfo()
					respDeleteToServer(m["msgId"], "ok", "rfNetInfo")
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

type sensorCfgFile struct {
	SensorList []modle.SensorInfo `json:"sensorList"`
}

func writeSensorCfgToFile(sList []modle.SensorInfo) error {
	cfg := sensorCfgFile{SensorList: sList}

	if b, err := json.Marshal(cfg); err == nil {
		f, er := os.OpenFile("sensorCfg.json", os.O_RDWR|os.O_CREATE, 0777)
		if er != nil {
			log.Printlntml(er)
			return err
		}
		defer f.Close()

		if _, err = f.Write(b); err != nil {
			log.Printlntml(err)
			return err
		}
	}

	if b, err := yaml.Marshal(cfg); err == nil {
		f, er := os.OpenFile("sensorCfg.yaml", os.O_RDWR|os.O_CREATE, 0777)
		if er != nil {
			log.Printlntml(er)
			return err
		}
		defer f.Close()

		if _, err = f.Write(b); err != nil {
			log.Printlntml(err)
			return err
		}
	}

	return nil
}

func readSensorCfgTofile(cfg *sensorCfgFile) {
	var err error
	var n int
	var f *os.File

	buf := make([]byte, 1024*100)

	f, err = os.OpenFile("sensorCfg.json", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		goto _exit
	}
	defer f.Close()

	if n, err = f.Read(buf); err != nil {
		goto _exit
	}
	if err = json.Unmarshal(buf[0:n], cfg); err != nil {
		goto _exit
	}
_exit:
	return
}

func waitSensorCfgInfoConfig() {
	var sensorInfo modle.SensorInfoReq
	var cfg sensorCfgFile
	readSensorCfgTofile(&cfg)
	config.WriteSensorConfig(cfg.SensorList)

	for {
		buf := <-sensorInfoCfgChannel

		if err := json.Unmarshal(buf, &sensorInfo); err != nil {
			log.PrintlnErr(err)
			respToServer(sensorInfo.MsgID, err.Error(), "sensorInfo")
			continue
		}

		if (len(sensorInfo.SensorList) == 0) && (sensorInfo.MsgType == "PUT") {
			respToServer(sensorInfo.MsgID, "SensorList is null", "sensorInfo")
		}

		switch sensorInfo.MsgType {
		case "GET":
			sensorCfgToServer(sensorInfo, config.ReadSensorConfig())
		case "PUT":
			config.WriteSensorConfig(sensorInfo.SensorList)
			writeSensorCfgToFile(sensorInfo.SensorList)
			respToServer(sensorInfo.MsgID, "ok", "sensorInfo")
		default:
			log.PrintfErr("json msgType:%s no support ", sensorInfo.MsgType)
		}
	}
}

func alarmCfgToServer(req modle.AlarmInfoReq, aInfo []modle.AlarmInfo) {

	param := &modle.AlarmInfoResp{
		MsgType:      "GET",
		MsgID:        req.MsgID,
		MsgGwID:      config.SysParamGwId(),
		MsgTimeStamp: time.Now().Unix(),
		MsgParam:     "alarmConfig",
		MsgResp:      "ok",
	}

	param.AlarmListNum = config.ReadAlarmCfgNum()
	param.AlarmList = aInfo

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

type alarmCfgFile struct {
	AlarmList []modle.AlarmInfo `json:"AlarmList"`
}

func writeAlarmCfgToFile(l []modle.AlarmInfo) {
	cfg := alarmCfgFile{AlarmList: l}

	if b, err := json.Marshal(cfg); err == nil {
		f, er := os.OpenFile("alarmCfg.json", os.O_RDWR|os.O_CREATE, 0777)
		if er != nil {
			log.Printlntml(er)
			return
		}
		defer f.Close()

		if _, err = f.Write(b); err != nil {
			log.Printlntml(err)
			return
		}
	}

	if b, err := yaml.Marshal(cfg); err == nil {
		f, er := os.OpenFile("alarmCfg.yaml", os.O_RDWR|os.O_CREATE, 0777)
		if er != nil {
			log.Printlntml(er)
			return
		}
		defer f.Close()

		if _, err = f.Write(b); err != nil {
			log.Printlntml(err)
			return
		}
	}

	return
}

func readAlarmCfgTofile(cfg *alarmCfgFile) {
	var err error
	var n int
	var f *os.File

	buf := make([]byte, 1024*100)

	f, err = os.OpenFile("alarmCfg.json", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		goto _exit
	}
	defer f.Close()

	if n, err = f.Read(buf); err != nil {
		goto _exit
	}
	if err = json.Unmarshal(buf[0:n], cfg); err != nil {
		goto _exit
	}

_exit:
	return
}

func waitAlarmCfgInfoConfig() {
	var alarmInfo modle.AlarmInfoReq
	var cfg alarmCfgFile
	readAlarmCfgTofile(&cfg)
	config.WriteAlarmCfg(cfg.AlarmList)

	for {
		buf := <-alarmInfoCfgChannel

		if err := json.Unmarshal(buf, &alarmInfo); err != nil {
			log.PrintlnErr(err)
			respToServer(alarmInfo.MsgID, err.Error(), "alarmConfig")
			continue
		}

		switch alarmInfo.MsgType {
		case "GET":
			alarmCfgToServer(alarmInfo, config.ReadAlarmCfg())
		case "PUT":
			config.WriteAlarmCfg(alarmInfo.AlarmList)
			writeAlarmCfgToFile(alarmInfo.AlarmList)
			respToServer(alarmInfo.MsgID, "ok", "alarmConfig")
		default:
			log.PrintfErr("json msgType:%s no support ", alarmInfo.MsgType)
		}
	}
}

func adapterCfgToServer(req modle.AdapterInfoReq, info modle.AdapterInfo) {

	param := &modle.AdapterInfoResp{
		MsgType:      "GET",
		MsgID:        req.MsgID,
		MsgGwID:      config.SysParamGwId(),
		MsgTimeStamp: time.Now().Unix(),
		MsgParam:     "adapter",
		MsgResp:      "ok",
	}

	param.AdapterInfo = info

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

func readAdapterCfgFromFile() {
	var err error
	var n int
	var f *os.File
	m := make(map[string]modle.AdapterInfo)

	buf := make([]byte, 1024*100)

	f, err = os.OpenFile("AdapterCfg.json", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		goto _exit
	}
	defer f.Close()

	if n, err = f.Read(buf); err != nil {
		goto _exit
	}
	if err = json.Unmarshal(buf[0:n], m); err != nil {
		goto _exit
	}

	config.InitAdapterInfo(m)

_exit:
	return
}

func writeAdapterCfgToFile() {
	m := config.ReadAdapterMapInfo()

	if b, err := json.Marshal(m); err == nil {
		f, er := os.OpenFile("AdapterCfg.json", os.O_RDWR|os.O_CREATE, 0777)
		if er != nil {
			log.Printlntml(er)
			return
		}
		defer f.Close()

		if _, err = f.Write(b); err != nil {
			log.Printlntml(err)
			return
		}
	}

	if b, err := yaml.Marshal(m); err == nil {
		f, er := os.OpenFile("AdapterCfg.yaml", os.O_RDWR|os.O_CREATE, 0777)
		if er != nil {
			log.Printlntml(er)
			return
		}
		defer f.Close()

		if _, err = f.Write(b); err != nil {
			log.Printlntml(err)
			return
		}
	}

	return
}

func waitAdapterInfoCfgInfoConfig() {
	var adapterInfo modle.AdapterInfoReq

	//读取本地设配器配置
	readAdapterCfgFromFile()

	for {
		buf := <-AdapterInfoCfgChannel
		if err := json.Unmarshal(buf, &adapterInfo); err != nil {
			log.PrintlnErr(err)
			respToServer(adapterInfo.MsgID, err.Error(), "adapter")
			continue
		}

		switch adapterInfo.MsgType {
		case "GET":
			if info, err := config.ReadAdapterInfo(adapterInfo.AdapterInfo.SensorID); err != nil {
				respToServer(adapterInfo.MsgID, err.Error(), "adapter")
			} else {
				adapterCfgToServer(adapterInfo, info)
			}
		case "PUT":
			config.WriteAdapterInfo(adapterInfo.AdapterInfo)
			writeAdapterCfgToFile()
		default:
			log.PrintfErr("json msgType:%s no support ", adapterInfo.MsgType)
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
	sensorInfoCfgChannel = make(chan []byte, 1)
	net.CreateMsgField("sensorInfo", sensorInfoCfgChannel)
	go waitSensorCfgInfoConfig()

	//报警参数配置
	//本地存储或读取配置文件，
	alarmInfoCfgChannel = make(chan []byte, 1)
	net.CreateMsgField("alarmConfig", alarmInfoCfgChannel)
	go waitAlarmCfgInfoConfig()

	//设配器参数配置
	AdapterInfoCfgChannel = make(chan []byte, 1)
	net.CreateMsgField("adapter", AdapterInfoCfgChannel)
	go waitAdapterInfoCfgInfoConfig()
	//TODO:需要开辟一个线程队列发送服务器的下发的配置数据

}

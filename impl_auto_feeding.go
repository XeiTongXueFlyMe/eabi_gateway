package main

import (
	modle "eabi_gateway/impl"
	busNet "eabi_gateway/impl/Industrial_bus"
	config "eabi_gateway/impl/config"
	"eabi_gateway/impl/modbus"
	net "eabi_gateway/impl/net"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

var feedingParamChannel chan []byte
var materialNumChannel chan []byte

func feedingParamToServer(req modle.AutoFeedingReq, feedingParam modle.AutoFeeding) {

	param := &modle.AutoFeedingResp{
		MsgType:      "GET",
		MsgID:        req.MsgID,
		MsgGwID:      config.SysParamGwId(),
		MsgTimeStamp: time.Now().Unix(),
		MsgParam:     "autoFeeding",
		MsgResp:      "ok",
	}
	param.AutoFeeding = feedingParam

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

func readFeedingParamCfgTofile(m *modle.AutoFeeding) {
	var err error
	var n int
	var f *os.File

	buf := make([]byte, 1024*10)

	f, err = os.OpenFile("FeedingParam.json", os.O_RDWR|os.O_CREATE, 0777)
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

_exit:
	return
}

func writeFeedingParamToFile(m modle.AutoFeeding) {

	if b, err := json.Marshal(m); err == nil {
		os.Remove("FeedingParam.json")
		f, er := os.OpenFile("FeedingParam.json", os.O_RDWR|os.O_CREATE, 0777)
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
		os.Remove("FeedingParam.yaml")
		f, er := os.OpenFile("FeedingParam.yaml", os.O_RDWR|os.O_CREATE, 0777)
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

func waitFeedingParamConfig() {
	var req modle.AutoFeedingReq
	var feedingParam modle.AutoFeeding
	readFeedingParamCfgTofile(&feedingParam)

	for {
		buf := <-feedingParamChannel

		if err := json.Unmarshal(buf, &req); err != nil {
			log.PrintlnErr(err)
			respToServer(req.MsgID, err.Error(), "autoFeeding")
			continue
		}

		switch req.MsgType {
		case "GET": //后面实时获取，或则系统重启后获取
			feedingParamToServer(req, feedingParam)
		case "PUT":
			feedingParam = req.AutoFeeding
			writeFeedingParamToFile(feedingParam)
			adapterAutoFeedingChannel <- feedingParam
			respToServer(req.MsgID, "ok", "autoFeeding")
		default:
			log.PrintfErr("json msgType:%s no support ", req.MsgType)
		}
	}
}

func materialNumParamToServer(req modle.MaterialNumReq, materialNum modle.MaterialNum) {

	param := &modle.MaterialNumResp{
		MsgType:      "GET",
		MsgID:        req.MsgID,
		MsgGwID:      config.SysParamGwId(),
		MsgTimeStamp: time.Now().Unix(),
		MsgParam:     "materialNum",
		MsgResp:      "ok",
	}
	param.MaterialNum = materialNum

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

func readMaterialNumCfgTofile(m *modle.MaterialNum) {
	var err error
	var n int
	var f *os.File

	buf := make([]byte, 1024*10)

	f, err = os.OpenFile("MaterialNum.json", os.O_RDWR|os.O_CREATE, 0777)
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

_exit:
	return
}

func writeMaterialNumToFile(m modle.MaterialNum) {

	if b, err := json.Marshal(m); err == nil {
		os.Remove("MaterialNum.json")
		f, er := os.OpenFile("MaterialNum.json", os.O_RDWR|os.O_CREATE, 0777)
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
		os.Remove("MaterialNum.yaml")
		f, er := os.OpenFile("MaterialNum.yaml", os.O_RDWR|os.O_CREATE, 0777)
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

func waitMaterialNumParamConfig() {
	var req modle.MaterialNumReq
	var materialNum modle.MaterialNum
	readMaterialNumCfgTofile(&materialNum)

	for {
		buf := <-materialNumChannel

		if err := json.Unmarshal(buf, &req); err != nil {
			log.PrintlnErr(err)
			respToServer(req.MsgID, err.Error(), "materialNum")
			continue
		}

		switch req.MsgType {
		case "GET":
			materialNumParamToServer(req, materialNum)
		case "PUT":
			materialNum = req.MaterialNum
			writeMaterialNumToFile(materialNum)
			respToServer(req.MsgID, "ok", "materialNum")
		default:
			log.PrintfErr("json msgType:%s no support ", req.MsgType)
		}
	}
}

func writeAlgoLog(s string) {
	f, err := os.OpenFile(AUTO_FEEDING_HISTORY_DATA_FILE_NAME, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(s)
}

//Pt油压 Pc套压
//oldPt oldPc Ygjpd　上位机可以配置，默认分别是0,0,2.5
var newPt, newPc float32

func autoFeedingAlgo() {
	var OpenFlag bool

	//等待设备第一次传参
	for newPt == 0 && newPc == 0 {
		time.Sleep(time.Second * 10)
	}

	for {
		var materialNum modle.MaterialNum
		readMaterialNumCfgTofile(&materialNum)

		//判断是否是485通信
		if config.SysHardware() != "485" {
			time.Sleep(time.Minute * 1)
			continue
		}

		//判断是否到时间点
		hour, min, _ := time.Now().Clock()
		if hour == int(materialNum.Hour) && min == int(materialNum.Minutes) {
			//判断开井，关井标志
			//油压&&套压　设配器2，第１通道，第２通道。　new - old > 2.5, old默认为０，上位机可配　Ygjpd默认为２.5上位机可配
			if (newPt-materialNum.OldPt > materialNum.Ygjpd) && (newPc-materialNum.OldPc > materialNum.Ygjpd) {
				OpenFlag = false //关井标志
			} else {
				OpenFlag = true //开井标志
			}
			//记录上次油压和套压
			materialNum.OldPt = newPt
			materialNum.OldPc = newPc
			writeMaterialNumToFile(materialNum)
			//判断是否处于开井,套压大于油压判断
			if OpenFlag && ((newPc - newPt) > 1) {
				//开始计算,
				d_ygjygd, d_tgjygd, d_jyl, d_jiayl, err := feedingCalculate(newPt, newPc)
				if err != nil {
					writeAlgoLog(err.Error())
					goto _continue
				}

				jy := int(d_jiayl)
				buf := make([]byte, 0)
				buf = append(buf, byte((jy>>8)&0x000000ff))
				buf = append(buf, byte(jy&0x000000ff))
				switch config.SysHardware() {
				case "485":
					busNet.Send(modbus.WriteDeviceReg(1, 26023, 1, 2, buf[:]))
				default:
					return
				}

				//记录
				s := fmt.Sprintf("当前套压=%d,当前油压=%d,油管积液高度=%d,套管积液高度=%d,积液量=%d,加液量=%d\n", newPc, newPt, d_ygjygd, d_tgjygd, d_jyl, d_jiayl)
				writeAlgoLog(s)
			}
		}

	_continue:
		time.Sleep(time.Minute * 1)
	}
}

func autoFeedingInit() {
	//网关参数的增删改查
	feedingParamChannel = make(chan []byte, 1)
	net.CreateMsgField("autoFeeding", feedingParamChannel)
	go waitFeedingParamConfig()

	//加液量参数
	materialNumChannel = make(chan []byte, 1)
	net.CreateMsgField("materialNum", materialNumChannel)
	go waitMaterialNumParamConfig()

	//定时自动投料
	go autoFeedingAlgo()
}

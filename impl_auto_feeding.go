package main

import (
	modle "eabi_gateway/impl"
	config "eabi_gateway/impl/config"
	net "eabi_gateway/impl/net"
	"encoding/json"
	"time"
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

func waitFeedingParamConfig() {
	var req modle.AutoFeedingReq
	var feedingParam modle.AutoFeeding

	for {
		buf := <-feedingParamChannel

		if err := json.Unmarshal(buf, &req); err != nil {
			log.PrintlnErr(err)
			respToServer(req.MsgID, err.Error(), "autoFeeding")
			continue
		}

		switch req.MsgType {
		case "GET": //TODO：后面实时获取，或则系统重启后获取
			feedingParamToServer(req, feedingParam)
		case "PUT":
			//TODO:需要通过４８５下发
			feedingParam = req.AutoFeeding
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

func waitMaterialNumParamConfig() {
	var req modle.MaterialNumReq
	var materialNum modle.MaterialNum

	for {
		buf := <-materialNumChannel

		if err := json.Unmarshal(buf, &req); err != nil {
			log.PrintlnErr(err)
			respToServer(req.MsgID, err.Error(), "materialNum")
			continue
		}

		switch req.MsgType {
		case "GET": //TODO：需要本地保存
			materialNumParamToServer(req, materialNum)
		case "PUT":
			materialNum = req.MaterialNum
			respToServer(req.MsgID, "ok", "materialNum")
		default:
			log.PrintfErr("json msgType:%s no support ", req.MsgType)
		}
	}
}

func autoFeedingAlgo() {
	for {
		//TODO:判断是否到时间点
		//hour, min, _ := time.Now().Clock()

		//TODO:判断开井，关井标志

		//TODO:记录上次油压和套压
		//TODO:判断是否处于开井
		//TODO:套压大于油压判断

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

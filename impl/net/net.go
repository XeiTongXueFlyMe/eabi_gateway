package net

import (
	model "eabi_gateway/impl"
	"eabi_gateway/impl/config"
	module "eabi_gateway/module"
	myLog "eabi_gateway/module/my_log"
	webs "eabi_gateway/module/websocket"
	"encoding/json"

	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	mu       sync.RWMutex
	net      module.NetInterfase
	netHeart chan []byte
	netPing  chan []byte
)

func SendData(buf []byte) (int, error) {
	mu.Lock()
	defer mu.Unlock()

	return net.Write(buf)
}

func waitNetReceive() {
	defer net.Close()

	for {
		buf := make([]byte, 1024*1024)
		mu.Lock()
		mu.Unlock()
		if n, err := net.Read(buf); err != nil {
			log.PrintlnWarring("websocker close :", err)
			break
		} else {
			if n != 0 {
				log.Printlntml(string(buf))
				netDataBufChan <- buf[0:n]
			}
		}
	}
}

func rebootNetConnet(host, path string) {
	mu.Lock()
	defer mu.Unlock()
	t := time.Now().Unix()

	net = &webs.Conn{}
	if err := net.Open(host, path); err != nil {

		log.PrintlnErr(err)

		panic(err)
	}

	log.PrintfInfo("Reconnect to the server %d Second after ", time.Now().Unix()-t)
	go waitNetReceive()
}

func ping() {
	for {
		mu.Lock()
		mu.Unlock()
		param := &model.StdReq{
			MsgType:      "GET",
			MsgID:        uuid.New().String(),
			MsgGwID:      config.SysParamGwId(),
			MsgTimeStamp: time.Now().Unix(),
			MsgParam:     "ping",
		}
		if buf, err := json.Marshal(param); err != nil {
			log.PrintlnErr(err)
		} else {
			SendData(buf)
		}
		time.Sleep(time.Duration(config.SysParamHeartCycle()) * time.Second)
	}
}

func waitHeart() {
	for {
		select {
		case <-netHeart:
		case <-time.After(time.Duration(config.SysParamHeartCycle()) * time.Second * 3):
			log.PrintfWarring("心跳超时")
			net.Close()
			ip, p := config.SysParamServerIPAndPort()
			rebootNetConnet(ip+":"+p, config.SysParamPath())
		}
	}
}

func pong() {
	for {
		<-netPing

		param := &model.StdResp{
			MsgType:      "GET",
			MsgID:        uuid.New().String(),
			MsgGwID:      config.SysParamGwId(),
			MsgTimeStamp: time.Now().Unix(),
			MsgParam:     "pong",
			MsgResp:      "ok",
		}
		if buf, err := json.Marshal(param); err != nil {
			log.PrintlnErr(err)
		} else {
			SendData(buf)
		}
	}
}

var log module.LogInterfase

func NetInit() {
	log = &myLog.L{}

	netHeart = make(chan []byte, 1)
	netPing = make(chan []byte, 1)

	CreateMsgField("pong", netHeart)
	CreateMsgField("ping", netPing)

	ip, p := config.SysParamServerIPAndPort()
	rebootNetConnet(ip+":"+p, config.SysParamPath())

	go ping()
	go pong()
	go waitHeart()
}

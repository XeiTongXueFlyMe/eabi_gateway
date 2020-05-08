package net

import (
	"eabi_gateway/impl/config"
	module "eabi_gateway/module"
	myLog "eabi_gateway/module/my_log"
	webs "eabi_gateway/module/websocket"

	"sync"
	"time"
)

var (
	mu       sync.RWMutex
	net      module.NetInterfase
	netHeart chan []byte
	netPing  chan []byte
)

func sendNetData(buf []byte) (int, error) {
	mu.Lock()
	defer mu.Unlock()

	return net.Write(buf)
}

func waitNetReceive() {
	defer net.Close()

	for {
		buf := make([]byte, 1024)
		mu.Lock()
		mu.Unlock()
		if n, err := net.Read(buf); err != nil {
			log.PrintlnWarring("websocker close :", err)
			break
		} else {
			log.Printlntml(string(buf))
			netDataBufChan <- buf[0:n]
		}
	}
}

func rebootNetConnet(host, path string) {
	mu.Lock()
	defer mu.Unlock()
	for {
		net = &webs.Conn{}
		if err := net.Open(host, path); err != nil {
			log.PrintlnErr(err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	go waitNetReceive()
}

func ping() {
	for {
		//TODO
		sendNetData([]byte("{\"msgType\":\"GET\",\"msgId\":\"a7356eac-71ae-4862-b66c-a212cd292baf\",\"msgGwId\":\"AFAF73EADCF5\",\"msgTimeStamp\":1586162503,\"msgParam\":\"ping\"}"))
		time.Sleep(1 * time.Second)
	}
}

func waitHeart() {
	for {
		select {
		case <-netHeart:
		case <-time.After(5 * time.Second):
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
		//TODO
		sendNetData([]byte("{\"msgType\":\"XXX\",\"msgId\":\"\",\"msgGwId\":\"XXXXXXXXXXXX\",\"msgTimeStamp\":1586162503,\"msgParam\":\"pong\"}"))
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

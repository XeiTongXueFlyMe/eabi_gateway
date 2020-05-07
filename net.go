package main

import (
	modle "eabi_gateway/model"
	webs "eabi_gateway/model/websocket"
	"sync"
	"time"
)

var (
	mu       sync.RWMutex
	net      modle.NetInterfase
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
			ip, p := sysParamServerIPAndPort()
			rebootNetConnet(ip+":"+p, sysParamPath())
		}
	}
}

func pong() {
	for {
		<-netPing
		sendNetData([]byte("{\"msgType\":\"XXX\",\"msgId\":\"\",\"msgGwId\":\"XXXXXXXXXXXX\",\"msgTimeStamp\":1586162503,\"msgParam\":\"pong\"}"))
	}
}

func netInit() {
	netHeart = make(chan []byte, 1)
	netPing = make(chan []byte, 1)

	createMsgField("pong", netHeart)
	createMsgField("ping", netPing)

	ip, p := sysParamServerIPAndPort()
	rebootNetConnet(ip+":"+p, sysParamPath())

	go ping()
	go pong()
	go waitHeart()
}

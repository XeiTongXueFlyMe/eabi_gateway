package main

import (
	modle "eabi_gateway/model"
	webs "eabi_gateway/model/websocket"
	"fmt"
	"sync"
	"time"
)

var (
	mu       sync.RWMutex
	net      modle.NetInterfase
	netHeart chan []byte
	netPing  chan []byte
)

func sendData(buf []byte) (int, error) {
	mu.Lock()
	defer mu.Unlock()

	return net.Write(buf)
}

func waitReceive() {
	defer net.Close()

	for {
		buf := make([]byte, 1024)
		mu.Lock()
		mu.Unlock()
		if n, err := net.Read(buf); err != nil {
			fmt.Println("websocker close :", err)
			break
		} else {
			fmt.Println(string(buf))
			netDataBufChan <- buf[0:n]
		}
	}
}

func rebootConnet(host, path string) {
	mu.Lock()
	defer mu.Unlock()
	for {
		net = &webs.Conn{}
		if err := net.Open(host, path); err != nil {
			fmt.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	go waitReceive()
}

func ping() {
	for {
		sendData([]byte("{\"msgType\":\"GET\",\"msgId\":\"a7356eac-71ae-4862-b66c-a212cd292baf\",\"msgGwId\":\"AFAF73EADCF5\",\"msgTimeStamp\":1586162503,\"msgParam\":\"ping\"}"))
		time.Sleep(1 * time.Second)
	}
}

func waitHeart() {
	for {
		select {
		case <-netHeart:
		case <-time.After(5 * time.Second):
			fmt.Println("心跳超时")
			net.Close()
			//TODO
			rebootConnet("192.168.0.168:8286", "/")
			//rebootConnet("120.55.191.153:8286", "/")
		}
	}
}

func pong() {
	for {
		<-netPing
		sendData([]byte("{\"msgType\":\"XXX\",\"msgId\":\"\",\"msgGwId\":\"XXXXXXXXXXXX\",\"msgTimeStamp\":1586162503,\"msgParam\":\"pong\"}"))
	}
}

func netInit(host, path string) {
	netHeart = make(chan []byte, 1)
	netPing = make(chan []byte, 1)

	createMsgField("pong", netHeart)
	createMsgField("ping", netPing)

	rebootConnet(host, path)
	go ping()
	go pong()
	go waitHeart()
}

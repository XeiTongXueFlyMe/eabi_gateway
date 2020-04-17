package main

import (
	modle "eabi_gateway/model"
	webs "eabi_gateway/model/websocket"
	"fmt"
	"sync"
	"time"
)

var net modle.NetInterfase
var netHeart chan bool
var mu sync.RWMutex

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
		if _, err := net.Read(buf); err != nil {
			fmt.Println("websocker close :", err)
			break
		}
		netDataBufChan <- buf
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
		sendData([]byte("{\"msgType\":\"GET\",\"msgId\":\"a7356eac-71ae-4862-b66c-a212cd292baf\",\"msgGwId\":\"AFAF73BADCF6\",\"msgTimeStamp\":1586162503,\"msgParam\":\"ping\"}"))

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
			rebootConnet("120.55.191.153:8286", "/")
		}
	}
}

func netInit(host, path string) {
	netHeart = make(chan bool, 1)

	rebootConnet(host, path)
	go ping()
	go waitHeart()
}

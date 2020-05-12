package main

import (
	webs "eabi_gateway/module/websocket"
	"fmt"
	"time"
)

func main() {
	msg := make([]byte, 1000)

	ws := &webs.Conn{}

	ws.Open("120.55.191.153:8286", "/")

	ws.Write([]byte("{\"type\":\"LOGIN\"}"))
	_, err := ws.Read(msg)
	if err != nil {
		panic(err)
	}
	count := 0
	for {
		count++
		msg := make([]byte, 1000)
		ws.Write([]byte("{\"type\":\"HEART\"}"))
		_, err := ws.Read(msg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("count = %d:", count)
		fmt.Println(string(msg))
		time.Sleep(1 * time.Second)
	}
}

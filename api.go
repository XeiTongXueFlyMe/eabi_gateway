package main

import (
	"encoding/json"
	"fmt"
)

var netDataBufChan chan []byte

//按api规定的格式解析
func waitNetData() {
	for {
		buf := <-netDataBufChan
		if !json.Valid(buf) {
			fmt.Println("json.Valid return false:", string(buf))
		}
		// json.Unmarshal(buf)
	}
}

func apiInit() {
	netDataBufChan = make(chan []byte, 1000)

	go waitNetData()
}

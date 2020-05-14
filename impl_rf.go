package main

import rfNet "eabi_gateway/impl/rf_net"

func waitRfData() {
	b := make([]byte, 1024)
	for {
		n := rfNet.Read(b)
		log.Printlntml("read rfdata num = ", n)
		log.Printlntml("read rfdata  = ", b[:n])
	}
}

func rfInit() {
	//TODO:lora网络状态维护
	//lora控制发送接口

	go waitRfData()
}

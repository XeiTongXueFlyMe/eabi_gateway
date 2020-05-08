package main

import (
	"eabi_gateway/impl"
	config "eabi_gateway/impl/config"
	net "eabi_gateway/impl/net"
	modle "eabi_gateway/model"
	myLog "eabi_gateway/model/my_log"

	"fmt"
	"time"
)

var log modle.LogInterfase

func hello() {
	log = &myLog.L{}
	log.PrintfInfo("eabi_gateway start runing :)")
}

func main() {
	hello()

	//初始化业务逻辑
	net.APIInit()

	//读取本地配置
	config.SysParamInit()

	impl.ImplInit()

	//TODO:初始化射频网络

	//lora网络信息统计
	//rfNetInfoInit()

	//初始化网络链接
	net.NetInit()

	fmt.Println("hi i am eabi_gateway :)")
	for {
		time.Sleep(1 * time.Hour)
	}
}

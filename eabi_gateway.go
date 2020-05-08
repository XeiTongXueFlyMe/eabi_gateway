package main

import (
	config "eabi_gateway/impl/config"
	net "eabi_gateway/impl/net"
	rfNet "eabi_gateway/impl/rf_net"
	module "eabi_gateway/module"
	myLog "eabi_gateway/module/my_log"

	"fmt"
	"time"
)

var log module.LogInterfase

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

	implInit()

	//TODO:初始化射频网络
	rfNet.RfNetInfoInit()

	//lora网络信息统计
	//rfNetInfoInit()

	//初始化网络链接
	net.NetInit()

	fmt.Println("hi i am eabi_gateway :)")
	for {
		time.Sleep(1 * time.Hour)
	}
}

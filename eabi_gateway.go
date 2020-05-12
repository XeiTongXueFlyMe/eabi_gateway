package main

import (
	config "eabi_gateway/impl/config"
	net "eabi_gateway/impl/net"
	rfNet "eabi_gateway/impl/rf_net"
	module "eabi_gateway/module"
	myLog "eabi_gateway/module/my_log"

	"time"
)

var log module.LogInterfase

func hello() {
	log = &myLog.L{}
	log.PrintfInfo("eabi_gateway start runing :)")
}

func main() {
	hello()
	//初始化业务路由
	net.APIInit()

	//读取本地配置
	config.SysParamInit()

	//初始化一些业务
	implInit()

	//TODO:初始化射频网络
	rfNet.RfNetInfoInit()

	//初始化网络链接
	net.NetInit()

	for {
		time.Sleep(1 * time.Hour)
	}
}

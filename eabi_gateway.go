package main

import (
	config "eabi_gateway/impl/config"
	net "eabi_gateway/impl/net"
	rfNet "eabi_gateway/impl/rf_net"
	module "eabi_gateway/module"
	"eabi_gateway/module/lora"
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

	//初始化射频网络
	defer lora.Close()
	rfNet.RfNetInfoInit()
	rfNet.LoraInit()
	rfInit()

	//初始化网络链接
	net.NetInit()

	//TODO:上传本地调试日志到服务器，每天上传一次，或则服务器主动获取
	for {
		time.Sleep(1 * time.Hour)
	}
}

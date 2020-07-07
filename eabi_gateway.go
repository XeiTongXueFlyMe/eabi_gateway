package main

import (
	config "eabi_gateway/impl/config"
	net "eabi_gateway/impl/net"
	rfNet "eabi_gateway/impl/rf_net"
	"eabi_gateway/impl/updata"
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

var foobar map[string]string

func main() {
	hello()
	//初始化业务路由
	net.APIInit()

	//读取本地配置
	config.SysParamInit()

	//初始化一些业务
	implInit()
	updata.Init()

	//初始化射频网络
	defer lora.Close()
	rfNet.RfNetInfoInit()
	rfNet.LoraInit()
	rfInit()

	//初始化自动喂料
	autoFeedingInit()

	//初始化网络链接
	net.NetInit()

	//上传本地调试日志到服务器，服务器主动获取
	logUpdataInit()

	for {
		time.Sleep(1 * time.Hour)
	}
}

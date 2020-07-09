package main

import (
	busNet "eabi_gateway/impl/Industrial_bus"
	config "eabi_gateway/impl/config"
	net "eabi_gateway/impl/net"
	rfNet "eabi_gateway/impl/rf_net"
	"eabi_gateway/impl/updata"
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
	//读取本地配置
	config.SysParamInit()

	//初始化业务路由
	net.APIInit()
	//初始化网络链接
	net.NetInit()
	//上传本地调试日志到服务器，服务器主动获取
	logUpdataInit()

	//初始化一些业务
	implInit()
	updata.Init()

	//初始化射频网络
	rfNet.LoraInit()
	//defer lora.Close()
	busNet.Init()
	//defer rs.Close()
	rfNet.RfNetInfoInit()

	rfInit()

	//初始化自动喂料
	autoFeedingInit()

	for {
		time.Sleep(1 * time.Hour)
	}
}

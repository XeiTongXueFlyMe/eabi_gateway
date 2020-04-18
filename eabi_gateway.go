package main

import (
	modle "eabi_gateway/model"
	myLog "eabi_gateway/model/my_log"
	"fmt"
	"time"
)

var log modle.LogInterfase

func main() {
	fmt.Println("eabi_gateway start runing")

	log = &myLog.L{}
	log.PrintfInfo("eabi_gateway start runing")

	//读取本地配置
	sysParamInit()

	//初始化业务逻辑
	apiInit()

	//初始化网络链接
	//TODO:配置参数，不从这里带入
	netInit(sysParamHost(), sysParamPath())

	//TODO:初始化射频网络

	for {
		time.Sleep(1 * time.Hour)
	}
}

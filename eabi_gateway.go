package main

import (
	modle "eabi_gateway/model"
	myLog "eabi_gateway/model/my_log"
	"fmt"
	"time"
)

var log modle.LogInterfase

func main() {
	log = &myLog.L{}
	log.PrintfInfo("eabi_gateway start runing")

	//初始化业务逻辑
	apiInit()

	//读取本地配置
	sysParamInit()

	//初始化网络链接
	//TODO:配置参数，不从这里带入
	netInit()

	//TODO:初始化射频网络

	fmt.Println("hi i am eabi_gateway :)")
	for {
		time.Sleep(1 * time.Hour)
	}
}

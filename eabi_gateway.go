package main

import (
	config "eabi_gateway/impl/config"
	modle "eabi_gateway/model"
	myLog "eabi_gateway/model/my_log"

	"fmt"
	"time"
)

var log modle.LogInterfase

func main() {
	log = &myLog.L{}
	log.PrintfInfo("eabi_gateway start runing :)")

	//初始化业务逻辑
	//apiInit()

	//读取本地配置
	config.SysParamInit()
	//TODO:初始化射频网络

	//lora网络信息统计
	//rfNetInfoInit()

	//初始化网络链接
	//TODO：为了测试，暂时关闭
	//netInit()

	fmt.Println("hi i am eabi_gateway :)")
	for {
		time.Sleep(1 * time.Hour)
	}
}

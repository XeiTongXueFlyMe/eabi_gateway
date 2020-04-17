package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("eabi_gateway start runing")
	//TODO:读取本地配置

	//TODO:初始化射频网络

	//初始化网络链接
	//TODO:配置参数，不从这里带入
	netInit("120.55.191.153:8286", "/")

	for {
		time.Sleep(1 * time.Hour)
	}
}

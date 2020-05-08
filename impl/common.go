package impl

import (
	"eabi_gateway/impl/config"
	"eabi_gateway/impl/net"
	modle "eabi_gateway/model"
	myLog "eabi_gateway/model/my_log"
	"encoding/json"
)

var gatewayParamChannel chan []byte
var log modle.LogInterfase

func waitGatewayParamConfig() {
	for {
		buf := <-gatewayParamChannel

		m := make(map[string]interface{})
		if err := json.Unmarshal(buf, &m); err != nil {
			log.PrintlnErr(err)
		}
		if v, ok := m["msgType"]; ok {
			if str, ok := v.(string); ok {
				switch str {
				case "GET":
					sendConfigToServer()
				case "PUT":
					config.ConfigTofile(m)
				}
			} else {
				log.PrintfErr("json msgType no is string")
			}
		} else {
			log.PrintlnErr("no find msgType")
		}
	}
}

func sendConfigToServer() {
	//TODO
}

func ImplInit() {
	log = &myLog.L{}
	gatewayParamChannel = make(chan []byte, 1)

	net.CreateMsgField("gatewayParam", gatewayParamChannel)
	go waitGatewayParamConfig()
}

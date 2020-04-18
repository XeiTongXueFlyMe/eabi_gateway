package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type SysParam struct {
	Myself struct {
		GwId          string `yaml:"gwId"`
		GwIP          string `yaml:"gwIP"`
		ServerIP      string `yaml:"serverIP"`
		ServerPort    string `yaml:"serverPort"`
		DataUpCycle   int    `yaml:"dataUpCycle"`
		HeartCycle    int    `yaml:"heartCycle"`
		DataReadCycle int    `yaml:"dataReadCycle"`
	}

	Websocket struct {
		Path string `yaml:"path"`
	}

	Rf struct {
		Id      string `yaml:"id"`
		Channel string `yaml:"channel"`
		NetId   string `yaml:"netId"`
	}
}

var cfgName = `./config.yaml`

var sysParam SysParam
var gatewayParamChannel chan []byte

//初始化，读取配置文件到缓存
func sysParamInit() {
	var err error
	var n int
	var f *os.File

	log.Printlntml("loading...")
	buf := make([]byte, 1000)

	f, err = os.OpenFile(cfgName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		goto _exit
	}
	defer f.Close()

	if n, err = f.Read(buf); err != nil {
		goto _exit
	}
	if err = yaml.Unmarshal(buf[0:n], &sysParam); err != nil {
		goto _exit
	}

	log.Printlntml(string(buf[0:n]))

	gatewayParamChannel = make(chan []byte, 1)
	createMsgField("gatewayParam", gatewayParamChannel)

	go waitGatewayParamConfig()
	return

_exit:
	panic(err)
}

func waitGatewayParamConfig() {
	for {
		buf := <-gatewayParamChannel

		m := make(map[string]interface{})
		if err := json.Unmarshal(buf, &m); err != nil {
			log.PrintlnErr(err)
		}
		for k, v := range m {
			switch k {
			case "gwId":
				if str, ok := v.(string); ok {
					sysParam.Myself.GwId = str
				} else {
					log.PrintfErr("json gwId no is string")
				}
			case "serverIP":
				if str, ok := v.(string); ok {
					sysParam.Myself.ServerIP = str
				} else {
					log.PrintfErr("json serverIP no is string")
				}
			case "serverPort":
				if str, ok := v.(string); ok {
					sysParam.Myself.ServerPort = str
				} else {
					log.PrintfErr("json serverPort no is string")
				}
			case "rfId":
				if str, ok := v.(string); ok {
					sysParam.Rf.Id = str
				} else {
					log.PrintfErr("json rfId no is string")
				}
			case "rfChannel":
				if n, ok := v.(int); ok {
					sysParam.Rf.Channel = fmt.Sprintln(n)
				} else {
					log.PrintfErr("json Channel no is int")
				}
			case "rfNetId":
				if str, ok := v.(string); ok {
					sysParam.Rf.NetId = str
				} else {
					log.PrintfErr("json rfNetId no is string")
				}
			case "dataUpCycle":
				if n, ok := v.(int); ok {
					sysParam.Myself.DataUpCycle = n
				} else {
					log.PrintfErr("json dataUpCycle no is int")
				}
			case "heartCycle":
				if n, ok := v.(int); ok {
					sysParam.Myself.HeartCycle = n
				} else {
					log.PrintfErr("json heartCycle no is int")
				}
			}
		}

		writeSysParamToFile()
	}
}

//写配置文件到缓存
func writeSysParamToFile() error {
	if b, err := yaml.Marshal(sysParam); err == nil {
		f, er := os.OpenFile(cfgName, os.O_RDWR|os.O_CREATE, 0777)
		if er != nil {
			log.PrintlnErr(er)
			return er
		}
		defer f.Close()

		if _, err = f.Write(b); err != nil {
			log.PrintlnErr(er)
			return err
		}
	}

	return nil
}

//各种参数读取与写入
func sysParamGwId(v ...interface{}) string {
	for _, arg := range v {
		if str, ok := arg.(string); ok {
			sysParam.Myself.GwId = str
			writeSysParamToFile()
		}
	}

	return sysParam.Myself.GwId
}

func sysParamGwIp(v ...interface{}) string {
	for _, arg := range v {
		if str, ok := arg.(string); ok {
			sysParam.Myself.GwIP = str
			writeSysParamToFile()
		}
	}

	return sysParam.Myself.GwIP
}

func sysParamPath(v ...interface{}) string {
	for _, arg := range v {
		if str, ok := arg.(string); ok {
			sysParam.Websocket.Path = str
			writeSysParamToFile()
		}
	}

	return sysParam.Websocket.Path
}

func sysParamServerIPAndPort(v ...interface{}) (string, string) {
	for n, arg := range v {
		if str, ok := arg.(string); ok {
			defer writeSysParamToFile()

			switch n {
			case 0:
				sysParam.Myself.ServerIP = str
			case 1:
				sysParam.Myself.ServerPort = str
			}

		}
	}

	return sysParam.Myself.ServerIP, sysParam.Myself.ServerPort
}

func sysParamDataUpCycle(v ...interface{}) int {
	for _, arg := range v {
		if value, ok := arg.(int); ok {
			sysParam.Myself.DataUpCycle = value
			writeSysParamToFile()
		}
	}

	return sysParam.Myself.DataUpCycle
}

func sysParamHeartCycle(v ...interface{}) int {
	for _, arg := range v {
		if value, ok := arg.(int); ok {
			sysParam.Myself.HeartCycle = value
			writeSysParamToFile()
		}
	}

	return sysParam.Myself.HeartCycle
}

func sysParamDataReadCycle(v ...interface{}) int {
	for _, arg := range v {
		if value, ok := arg.(int); ok {
			sysParam.Myself.DataReadCycle = value
			writeSysParamToFile()
		}
	}

	return sysParam.Myself.DataReadCycle
}

// id channel netid
func sysParamRf(v ...interface{}) (string, string, string) {
	for n, arg := range v {
		if str, ok := arg.(string); ok {
			defer writeSysParamToFile()

			switch n {
			case 0:
				sysParam.Rf.Id = str
			case 1:
				sysParam.Rf.Channel = str
			case 2:
				sysParam.Rf.NetId = str
			}

		}
	}

	return sysParam.Rf.Id, sysParam.Rf.Channel, sysParam.Rf.NetId
}

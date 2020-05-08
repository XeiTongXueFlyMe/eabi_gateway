package config

import (
	modle "eabi_gateway/model"
	myLog "eabi_gateway/model/my_log"
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type SysParam struct {
	Myself struct {
		GwID          string `yaml:"gwId"`
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
		ID      string `yaml:"id"`
		Channel string `yaml:"channel"`
		NetID   string `yaml:"netId"`
	}
}

var cfgName = `./config.yaml`

var sysParam SysParam
var gatewayParamChannel chan []byte
var log modle.LogInterfase

//初始化，读取配置文件到缓存
func SysParamInit() {
	var err error
	var n int
	var f *os.File

	log = &myLog.L{}
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
	//TODO
	//createMsgField("gatewayParam", gatewayParamChannel)

	go waitGatewayParamConfig()
	return

_exit:
	panic(err)
}

//TODO　放在你这里不合适
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
					configTofile(m)
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

func configTofile(m map[string]interface{}) {
	defer writeSysParamToFile()

	for k, v := range m {
		switch k {
		case "gwId":
			if str, ok := v.(string); ok {
				sysParam.Myself.GwID = str
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
				sysParam.Rf.ID = str
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
				sysParam.Rf.NetID = str
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
func SysParamGwId(v ...interface{}) string {
	for _, arg := range v {
		if str, ok := arg.(string); ok {
			sysParam.Myself.GwID = str
			writeSysParamToFile()
		}
	}

	return sysParam.Myself.GwID
}

func SysParamGwIp(v ...interface{}) string {
	for _, arg := range v {
		if str, ok := arg.(string); ok {
			sysParam.Myself.GwIP = str
			writeSysParamToFile()
		}
	}

	return sysParam.Myself.GwIP
}

func SysParamPath(v ...interface{}) string {
	for _, arg := range v {
		if str, ok := arg.(string); ok {
			sysParam.Websocket.Path = str
			writeSysParamToFile()
		}
	}

	return sysParam.Websocket.Path
}

func SysParamServerIPAndPort(v ...interface{}) (string, string) {
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

func SysParamDataUpCycle(v ...interface{}) int {
	for _, arg := range v {
		if value, ok := arg.(int); ok {
			sysParam.Myself.DataUpCycle = value
			writeSysParamToFile()
		}
	}

	return sysParam.Myself.DataUpCycle
}

func SysParamHeartCycle(v ...interface{}) int {
	for _, arg := range v {
		if value, ok := arg.(int); ok {
			sysParam.Myself.HeartCycle = value
			writeSysParamToFile()
		}
	}

	return sysParam.Myself.HeartCycle
}

func SysParamDataReadCycle(v ...interface{}) int {
	for _, arg := range v {
		if value, ok := arg.(int); ok {
			sysParam.Myself.DataReadCycle = value
			writeSysParamToFile()
		}
	}

	return sysParam.Myself.DataReadCycle
}

// id channel netid
func SysParamRf(v ...interface{}) (string, string, string) {
	for n, arg := range v {
		if str, ok := arg.(string); ok {
			defer writeSysParamToFile()

			switch n {
			case 0:
				sysParam.Rf.ID = str
			case 1:
				sysParam.Rf.Channel = str
			case 2:
				sysParam.Rf.NetID = str
			}

		}
	}

	return sysParam.Rf.ID, sysParam.Rf.Channel, sysParam.Rf.NetID
}

package config

import (
	module "eabi_gateway/module"
	myLog "eabi_gateway/module/my_log"
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
		Hardware string `yaml:"hardware"` //485,lora
		ID       string `yaml:"id"`
		Channel  string `yaml:"channel"`
		NetID    string `yaml:"netId"`
	}
}

var cfgName = `./config.yaml`

var sysParam SysParam
var log module.LogInterfase

//SysParamInit 初始化，读取配置文件到缓存
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

	return

_exit:
	panic(err)
}

func ConfigTofile(m map[string]interface{}) string {
	resp := "ok"
	defer writeSysParamToFile()

	for k, v := range m {
		switch k {
		case "gwId":
			if str, ok := v.(string); ok {
				sysParam.Myself.GwID = str
			} else {
				log.PrintfErr("json gwId no is string")
				resp = "json gwId no is string"
			}
		case "serverIP":
			if str, ok := v.(string); ok {
				sysParam.Myself.ServerIP = str
			} else {
				log.PrintfErr("json serverIP no is string")
				resp = "json serverIP no is string"
			}
		case "serverPort":
			if str, ok := v.(string); ok {
				sysParam.Myself.ServerPort = str
			} else {
				log.PrintfErr("json serverPort no is string")
				resp = "json serverPort no is string"
			}
		case "rfId":
			if str, ok := v.(string); ok {
				sysParam.Rf.ID = str
			} else {
				log.PrintfErr("json rfId no is string")
				resp = "json rfId no is string"
			}
		case "rfChannel":
			if n, ok := v.(string); ok {
				sysParam.Rf.Channel = n
			} else {
				log.PrintfErr("json Channel no is int")
				resp = "json Channel no is int"
			}
		case "rfNetId":
			if str, ok := v.(string); ok {
				sysParam.Rf.NetID = str
			} else {
				log.PrintfErr("json rfNetId no is string")
				resp = "json rfNetId no is string"
			}
		case "dataUpCycle":
			if n, ok := v.(float64); ok {
				sysParam.Myself.DataUpCycle = int(n)
			} else {
				log.PrintfErr("json dataUpCycle no is int")
				resp = "json dataUpCycle no is int"
			}
		case "heartCycle":
			if n, ok := v.(float64); ok {
				sysParam.Myself.HeartCycle = int(n)
			} else {
				log.PrintfErr("json heartCycle no is int")
				resp = "json heartCycle no is int"
			}
		case "dataReadCycle":
			if n, ok := v.(float64); ok {
				sysParam.Myself.DataReadCycle = int(n)
			} else {
				log.PrintfErr("json dataReadCycle no is int")
				resp = "json dataReadCycle no is int"
			}
		case "hardware":
			if str, ok := v.(string); ok {
				sysParam.Rf.Hardware = str
			} else {
				log.PrintfErr("json Hardware no is srting")
				resp = "json Hardware no is srting"
			}
		}

	}

	return resp
}

//写缓存到配置文件
func writeSysParamToFile() error {
	if b, err := yaml.Marshal(sysParam); err == nil {
		os.Remove(cfgName)
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

func SysHardware() string { return sysParam.Rf.Hardware }

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

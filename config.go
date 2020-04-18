package main

import (
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

//初始化，读取配置文件到缓存
func sysParamInit() {
	var err error
	var n int
	var f *os.File

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
	return

_exit:
	panic(err)
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

package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type SysParam struct {
	Myself struct {
		MsgGwId string `yaml:"msgGwId"`
	}

	Websocket struct {
		Host string `yaml:"host"`
		Path string `yaml:"path"`
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
			sysParam.Myself.MsgGwId = str
			writeSysParamToFile()
		}
	}

	return sysParam.Myself.MsgGwId
}

func sysParamHost(v ...interface{}) string {
	for _, arg := range v {
		if str, ok := arg.(string); ok {
			sysParam.Websocket.Host = str
			writeSysParamToFile()
		}
	}

	return sysParam.Websocket.Host
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

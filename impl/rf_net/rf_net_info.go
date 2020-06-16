package rfNet

import (
	"os"

	modle "eabi_gateway/impl"

	"gopkg.in/yaml.v2"
)

var chanMap map[string]string
var rfNetInfoMap map[string]modle.RfNetInfo
var rfNetInfoFileName = "./rf_net_info.yaml"

func RfNetInfoInit() {
	rfNetInfoMap = make(map[string]modle.RfNetInfo)
	chanMap = make(map[string]string)
}

func CleanInfo() {
	defer writeSysParamToFile()

	for k := range rfNetInfoMap {
		delete(rfNetInfoMap, k)
	}

	for k := range chanMap {
		delete(chanMap, k)
	}
}

func ReadInfo() map[string]modle.RfNetInfo {
	return rfNetInfoMap
}

func WriteInfo(id, name, sVer, hVer, channel string) {
	info := modle.RfNetInfo{
		ID:          id,
		Name:        name,
		SoftwareVer: sVer,
		HardwareVer: hVer,
	}

	//读取旧的的数据
	if v, ok := rfNetInfoMap[info.ID]; ok {
		info.Channel = append(info.Channel, v.Channel...)
	}

	//查询是否新的数据
	if _, ok := chanMap[info.ID+channel]; !ok {
		info.Channel = append(info.Channel, "Channel_"+channel)
		chanMap[info.ID+channel] = "Channel_" + channel
	}

	defer writeSysParamToFile()
	rfNetInfoMap[info.ID] = info
}

//写缓存到配置文件
func writeSysParamToFile() error {
	if b, err := yaml.Marshal(rfNetInfoMap); err == nil {
		os.Remove(rfNetInfoFileName)
		f, er := os.OpenFile(rfNetInfoFileName, os.O_RDWR|os.O_CREATE, 0777)
		if er != nil {
			return er
		}
		defer f.Close()

		if _, err = f.Write(b); err != nil {
			return err
		}
	}

	return nil
}

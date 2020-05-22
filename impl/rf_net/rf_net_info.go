package rfNet

import (
	"os"

	modle "eabi_gateway/impl"

	"gopkg.in/yaml.v2"
)

var rfNetInfoMap map[string]modle.RfNetInfo
var rfNetInfoFileName = "./rf_net_info.yaml"

func RfNetInfoInit() {
	rfNetInfoMap = make(map[string]modle.RfNetInfo)
}

func CleanInfo() {
	defer writeSysParamToFile()

	for k := range rfNetInfoMap {
		delete(rfNetInfoMap, k)
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
		Channel1:    "no",
		Channel2:    "no",
		Channel3:    "no",
		Channel4:    "no",
		Channel5:    "no",
		Channel6:    "no",
		Channel7:    "no",
		Channel8:    "no",
	}

	if v, ok := rfNetInfoMap[info.ID]; ok {
		info.Channel1 = v.Channel1
		info.Channel2 = v.Channel2
		info.Channel3 = v.Channel3
		info.Channel4 = v.Channel4
		info.Channel5 = v.Channel5
		info.Channel6 = v.Channel6
		info.Channel7 = v.Channel7
		info.Channel8 = v.Channel8
	}

	switch channel {
	case "1":
		info.Channel1 = "yes"
	case "2":
		info.Channel2 = "yes"
	case "3":
		info.Channel3 = "yes"
	case "4":
		info.Channel4 = "yes"
	case "5":
		info.Channel5 = "yes"
	case "6":
		info.Channel6 = "yes"
	case "7":
		info.Channel7 = "yes"
	case "8":
		info.Channel8 = "yes"

	}

	defer writeSysParamToFile()
	rfNetInfoMap[info.ID] = info
}

//写缓存到配置文件
func writeSysParamToFile() error {
	if b, err := yaml.Marshal(rfNetInfoMap); err == nil {
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

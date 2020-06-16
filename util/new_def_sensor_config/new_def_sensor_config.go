package main

import (
	modle "eabi_gateway/impl"
	"encoding/json"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	var cList []modle.ChanneInfo
	var sList []modle.SensorInfo

	cList = append(cList, modle.ChanneInfo{
		Channel:    1,
		AAdder:     11,
		ASize:      12,
		BAdder:     13,
		BSize:      14,
		WorkAdder:  15,
		WorkSize:   16,
		ValueAdder: 17,
		ValueSize:  4,
		ValueType:  "温度(℃)",
	})
	cList = append(cList, modle.ChanneInfo{
		Channel:    2,
		AAdder:     21,
		ASize:      22,
		BAdder:     23,
		BSize:      24,
		WorkAdder:  25,
		WorkSize:   26,
		ValueAdder: 27,
		ValueSize:  4,
		ValueType:  "压力(pa)",
	})

	sList = append(sList, modle.SensorInfo{
		SensorID:    "111111111111",
		SensorName:  "东1层1号温度气压表",
		SensorAdder: 0x01,
		DataAdder:   0,
		DataSize:    31,
		ChannelList: cList,
	})
	sList = append(sList, modle.SensorInfo{
		SensorID:    "222222222222",
		SensorName:  "东1层１号温度气压表",
		SensorAdder: 0x02,
		DataAdder:   0,
		DataSize:    31,
		ChannelList: cList,
	})
	sList = append(sList, modle.SensorInfo{
		SensorID:    "333333333333",
		SensorName:  "东1层3号气表",
		SensorAdder: 0x03,
		DataAdder:   0,
		DataSize:    31,
		ChannelList: cList,
	})
	writeSensorCfgToFile(sList)
}

type sensorCfgFile struct {
	SensorList []modle.SensorInfo `json:"sensorList"`
}

func writeSensorCfgToFile(sList []modle.SensorInfo) {
	cfg := sensorCfgFile{SensorList: sList}

	if b, err := json.Marshal(cfg); err == nil {
		os.Remove("sensorCfg.json")
		f, er := os.OpenFile("sensorCfg.json", os.O_RDWR|os.O_CREATE, 0777)
		if er != nil {
			return
		}
		defer f.Close()

		if _, err = f.Write(b); err != nil {
			return
		}
	}

	if b, err := yaml.Marshal(cfg); err == nil {
		os.Remove("sensorCfg.yaml")
		f, er := os.OpenFile("sensorCfg.yaml", os.O_RDWR|os.O_CREATE, 0777)
		if er != nil {
			return
		}
		defer f.Close()

		if _, err = f.Write(b); err != nil {
			return
		}
	}

	return
}

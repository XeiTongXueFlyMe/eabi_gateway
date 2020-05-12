package config

import (
	modle "eabi_gateway/impl"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteSensorCfg(t *testing.T) {
	var sList []modle.SensorInfo
	var cList []modle.ChanneInfo

	cList = append(cList, modle.ChanneInfo{
		Channel:    1,
		AAdder:     11,
		ASize:      12,
		BAdder:     13,
		BSize:      14,
		WorkAdder:  15,
		WorkSize:   16,
		ValueAdder: 17,
		ValueSize:  18,
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
		ValueSize:  28,
		ValueType:  "温度(℃)",
	})

	sList = append(sList, modle.SensorInfo{
		SensorID:    "123456",
		SensorName:  "东1层1号气表",
		SensorAdder: 0x5,
		DataAdder:   1,
		DataSize:    4,
	})
	sList = append(sList, modle.SensorInfo{
		SensorID:    "234567",
		SensorName:  "东1层2号气表",
		ChannelList: cList,
	})
	sList = append(sList, modle.SensorInfo{
		SensorID:   "3456789",
		SensorName: "东1层3号气表",
	})

	WriteSensorConfig(sList)

	if v, err := SensorNameFromID("3456789"); err != nil {
		panic(err)
	} else {
		assert.EqualValues(t, "东1层3号气表", v)
	}
	if v, err := SensorAdderFromID("123456"); err != nil {
		panic(err)
	} else {
		assert.EqualValues(t, 5, v)
	}
	if adder, vsize, err := SensorDataFromID("123456"); err != nil {
		panic(err)
	} else {
		assert.EqualValues(t, 1, adder)
		assert.EqualValues(t, 4, vsize)
	}
	if adder, vsize, err := ChannelAFromID("234567", 1); err != nil {
		panic(err)
	} else {
		assert.EqualValues(t, 11, adder)
		assert.EqualValues(t, 12, vsize)
	}
	if adder, vsize, err := ChannelBFromID("234567", 2); err != nil {
		panic(err)
	} else {
		assert.EqualValues(t, 23, adder)
		assert.EqualValues(t, 24, vsize)
	}
	if adder, vsize, err := ChannelWorkFromID("234567", 1); err != nil {
		panic(err)
	} else {
		assert.EqualValues(t, 15, adder)
		assert.EqualValues(t, 16, vsize)
	}
	if adder, vsize, typeV, err := ChannelValueFromID("234567", 2); err != nil {
		panic(err)
	} else {
		assert.EqualValues(t, 27, adder)
		assert.EqualValues(t, 28, vsize)
		assert.EqualValues(t, "温度(℃)", typeV)
	}
	assert.EqualValues(t, 3, ReadSensorConfigNum())
}

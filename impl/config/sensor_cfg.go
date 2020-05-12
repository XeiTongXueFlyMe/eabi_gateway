package config

import (
	modle "eabi_gateway/impl"
	"fmt"
)

//SensorList 包含传感器modbus配置信息
type SensorList []modle.SensorInfo

func (t *SensorList) findSensorFromID(sensorID string) (*modle.SensorInfo, error) {
	for n, v := range *t {
		if v.SensorID == sensorID {
			return &(*t)[n], nil
		}
	}

	return &modle.SensorInfo{}, fmt.Errorf("no find sensorId:%s in SensorList", sensorID)
}

//SensorNameFromID 获取传感器名称
func (t *SensorList) SensorNameFromID(sensorID string) (string, error) {
	if sensor, err := t.findSensorFromID(sensorID); err != nil {
		return "", err
	} else {
		return sensor.SensorName, nil
	}
}

//SensorAdderFromID 获取传感器modbus地址
func (t *SensorList) SensorAdderFromID(sensorID string) (int, error) {
	if sensor, err := t.findSensorFromID(sensorID); err != nil {
		return 0, err
	} else {
		return sensor.SensorAdder, nil
	}
}

//SensorDataFromID 获取传感数据地址及其数据总长度，方便一次性获取全部通道的数据
func (t *SensorList) SensorDataFromID(sensorID string) (int, int, error) {
	if sensor, err := t.findSensorFromID(sensorID); err != nil {
		return 0, 0, err
	} else {
		return sensor.DataAdder, sensor.DataSize, nil
	}
}

func (t *SensorList) findChanneInfoFromID(sensorID string, channel int) (*modle.ChanneInfo, error) {
	for n, v := range *t {
		if v.SensorID == sensorID {
			for cnt, value := range (*t)[n].ChannelList {
				if value.Channel == channel {
					return &((*t)[n].ChannelList)[cnt], nil
				}
			}
			goto _exit
		}
	}

_exit:
	return &modle.ChanneInfo{}, fmt.Errorf("no find sensorId:%s and channel %d in SensorList", sensorID, channel)
}

//ChannelAFromID 获取传感器某个通道的标定数据a的modbus地址和大小
func (t *SensorList) ChannelAFromID(sensorID string, channel int) (int, int, error) {
	var cInfo *modle.ChanneInfo
	var err error

	if cInfo, err = t.findChanneInfoFromID(sensorID, channel); err != nil {
		return 0, 0, err
	}

	return cInfo.AAdder, cInfo.ASize, nil
}

//ChannelBFromID 获取传感器某个通道的标定数据b的modbus地址和大小
func (t *SensorList) ChannelBFromID(sensorID string, channel int) (int, int, error) {
	var cInfo *modle.ChanneInfo
	var err error

	if cInfo, err = t.findChanneInfoFromID(sensorID, channel); err != nil {
		return 0, 0, err
	}

	return cInfo.BAdder, cInfo.BSize, nil
}

//ChannelWorkFromID 获取传感器某个通道的工作状态的地址和大小
func (t *SensorList) ChannelWorkFromID(sensorID string, channel int) (int, int, error) {
	var cInfo *modle.ChanneInfo
	var err error

	if cInfo, err = t.findChanneInfoFromID(sensorID, channel); err != nil {
		return 0, 0, err
	}

	return cInfo.WorkAdder, cInfo.WorkSize, nil
}

//ChannelValueFromID 获取传感器某个通道的值的地址和大小，值类型
func (t *SensorList) ChannelValueFromID(sensorID string, channel int) (int, int, string, error) {
	var cInfo *modle.ChanneInfo
	var err error

	if cInfo, err = t.findChanneInfoFromID(sensorID, channel); err != nil {
		return 0, 0, "", err
	}

	return cInfo.ValueAdder, cInfo.ValueSize, cInfo.ValueType, nil
}

//WriteSensorConfig 写传感器配置
func (t *SensorList) WriteSensorConfig(sList []modle.SensorInfo) {
	*t = sList
}

//ReadSensorConfig 读传感器配置
func (t *SensorList) ReadSensorConfig() []modle.SensorInfo {
	return *t
}

//ReadSensorConfigNum 读传感器配置数量
func (t *SensorList) ReadSensorConfigNum() int {
	return len(*t)
}

//NewSensorConfig 初始化传感器配置
func NewSensorConfig() *SensorList {
	return &SensorList{}
}

var common SensorList

//SensorNameFromID 获取传感器名称
func SensorNameFromID(sensorID string) (string, error) {
	return common.SensorNameFromID(sensorID)
}

//SensorAdderFromID 获取传感器modbus地址
func SensorAdderFromID(sensorID string) (int, error) {
	return common.SensorAdderFromID(sensorID)
}

//SensorDataFromID 获取传感数据地址及其数据总长度，方便一次性获取全部通道的数据
func SensorDataFromID(sensorID string) (int, int, error) {
	return common.SensorDataFromID(sensorID)
}

//ChannelAFromID 获取传感器某个通道的标定数据a的modbus地址和大小
func ChannelAFromID(sensorID string, channel int) (int, int, error) {
	return common.ChannelAFromID(sensorID, channel)
}

//ChannelBFromID 获取传感器某个通道的标定数据b的modbus地址和大小
func ChannelBFromID(sensorID string, channel int) (int, int, error) {
	return common.ChannelBFromID(sensorID, channel)
}

//ChannelWorkFromID 获取传感器某个通道的工作状态的地址和大小
func ChannelWorkFromID(sensorID string, channel int) (int, int, error) {
	return common.ChannelWorkFromID(sensorID, channel)
}

//ChannelValueFromID 获取传感器某个通道的值的地址和大小，值类型
func ChannelValueFromID(sensorID string, channel int) (int, int, string, error) {
	return common.ChannelValueFromID(sensorID, channel)
}

//WriteSensorConfig 写传感器配置
func WriteSensorConfig(sList []modle.SensorInfo) {
	common.WriteSensorConfig(sList)
}

//ReadSensorConfig 读传感器配置
func ReadSensorConfig() []modle.SensorInfo {
	return common.ReadSensorConfig()
}

//ReadSensorConfigNum 读传感器配置数量
func ReadSensorConfigNum() int {
	return common.ReadSensorConfigNum()
}

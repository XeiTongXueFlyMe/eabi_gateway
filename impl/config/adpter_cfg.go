package config

import (
	modle "eabi_gateway/impl"
	"errors"
)

//AdapterMap 适配器与仪表配置信息
var AdapterMap map[string]modle.AdapterInfo

//ReadAdapterInfo 读取某一个适配器数据
func ReadAdapterInfo(SensorID string) (modle.AdapterInfo, error) {
	if AdapterMap == nil {
		goto _end
	}

	if v, ok := AdapterMap[SensorID]; ok {
		return v, nil
	}

_end:
	return modle.AdapterInfo{}, errors.New("no find")
}

//WriteAdapterInfo 写入某一个适配器数据
func WriteAdapterInfo(info modle.AdapterInfo) {
	if AdapterMap == nil {
		AdapterMap = make(map[string]modle.AdapterInfo)
	}
	AdapterMap[info.SensorID] = info
}

//InitAdapterInfo 初始化设配信息
func InitAdapterInfo(m map[string]modle.AdapterInfo) {
	AdapterMap = m
}

//ReadAdapterMapInfo 初始化设配信息
func ReadAdapterMapInfo() map[string]modle.AdapterInfo {
	if AdapterMap == nil {
		AdapterMap = make(map[string]modle.AdapterInfo)
	}

	return AdapterMap
}

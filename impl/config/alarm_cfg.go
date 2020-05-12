package config

import (
	modle "eabi_gateway/impl"
	"errors"
)

//AlarmList 报警参数配置信息
type AlarmList []modle.AlarmInfo

func (t *AlarmList) findAlarmInfo(s string, c int) (modle.AlarmInfo, error) {
	for n, v := range *t {
		if (v.SensorID == s) && (v.Channel == c) {
			return (*t)[n], nil
		}
	}
	return modle.AlarmInfo{}, errors.New("no find")
}

//ReadAlarmCfgNum 获取报警参数配置数量
func (t *AlarmList) ReadAlarmCfgNum() int {
	return len(*t)
}

//IsAlarm 是否到达报警预值
func (t *AlarmList) IsAlarm(sensorID string, channel int, value float32) bool {
	v, err := t.findAlarmInfo(sensorID, channel)
	if err != nil {
		return false
	}
	if (value < v.AlarmValueL) || (value > v.AlarmValueH) {
		return true
	}
	return false
}

//ReadAlarmParamLH 返回报警高低预值
//v.AlarmValueL, v.AlarmValueH, nil
func (t *AlarmList) ReadAlarmParamLH(sensorID string, channel int) (float32, float32, error) {
	v, err := t.findAlarmInfo(sensorID, channel)
	if err != nil {
		return 0, 0, err
	}
	return v.AlarmValueL, v.AlarmValueH, nil
}

//WriteAlarmCfg 写报警参数配置
func (t *AlarmList) WriteAlarmCfg(alarmList []modle.AlarmInfo) {
	*t = alarmList
}

//ReadAlarmCfg 读报警参数配置
func (t *AlarmList) ReadAlarmCfg() []modle.AlarmInfo {
	return *t
}

var def AlarmList

//IsAlarm 是否到达报警预值
func IsAlarm(sensorID string, channel int, value float32) bool {
	return def.IsAlarm(sensorID, channel, value)
}

//ReadAlarmParamLH 返回报警高低预值
//v.AlarmValueL, v.AlarmValueH, nil
func ReadAlarmParamLH(sensorID string, channel int) (float32, float32, error) {
	return def.ReadAlarmParamLH(sensorID, channel)
}

//WriteAlarmCfg 写报警参数配置
func WriteAlarmCfg(alarmList []modle.AlarmInfo) {
	def.WriteAlarmCfg(alarmList)
}

//ReadAlarmCfg 读报警参数配置
func ReadAlarmCfg() []modle.AlarmInfo {
	return def.ReadAlarmCfg()
}

//ReadAlarmCfgNum 获取报警参数配置数量
func ReadAlarmCfgNum() int {
	return def.ReadAlarmCfgNum()
}

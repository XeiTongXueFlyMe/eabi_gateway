package updata

import (
	modle "eabi_gateway/impl"
	"eabi_gateway/impl/config"
	"eabi_gateway/module"
	myLog "eabi_gateway/module/my_log"
	"time"

	"github.com/google/uuid"
)

//临时上传文件使用uuid命名，并且存在内存文件系统中，不长期保存，上传到服务器后直接销毁
//在sd卡中存储所有采集的数据，没有sd卡，不存储数据，按每天一个文件进行存储。
var log module.LogInterfase
var dataCsv *DataCsv
var alarmCsv *DataCsv

//Init 初始化数据和报警数据上传
func Init() {
	log = &myLog.L{}

	dataCsv = NewDataCsv("updata", "~/sdb/")
	alarmCsv = NewDataCsv("alarm", "~/sdb/")
	go dataUpCycle()

	return
}

//WriteUpdata 写入需要上传的数据
func WriteUpdata(SensorName, SensorID, Unit string, Channel uint32, Value float32) {
	d := modle.UpDataMetaInfo{
		SourceID:  uuid.New().String(),
		GwID:      config.SysParamGwId(),
		TimeStamp: time.Now().Unix(),
	}
	d.SensorName = SensorName
	d.SensorID = SensorID
	d.Unit = Unit
	d.Channel = Channel
	d.Value = Value

	dataCsv.Write(d)
}

//WriteAlarmdata 写入需要上传的报警数据
//Isok "ok" "alarm"
func WriteAlarmdata(SensorName, SensorID, Isok string, Channel uint32, AlarmParamH, AlarmParamL, Value float32) {
	d := modle.AlarmMetaInfo{
		AlarmID:   uuid.New().String(),
		GwID:      config.SysParamGwId(),
		TimeStamp: time.Now().Unix(),
	}
	d.SensorName = SensorName
	d.SensorID = SensorID
	d.Channel = Channel
	d.Param = Value
	d.AlarmParamH = AlarmParamH
	d.AlarmParamL = AlarmParamL
	d.Isok = Isok

	alarmCsv.Write(d)
}

//TODO：按上传周期主动上传。
func dataUpCycle() {
	for {
		<-time.After(time.Duration(config.SysParamDataUpCycle()) * time.Second)
		//TODO：读取需要上传的文件，
		if fp := dataCsv.ReadUpFilePath(); fp != "" {
			//TODO：
			dataCsv.DeleteUpFile(fp)
		}

		//TODO：读取需要上传的报警文件，
		if fp := alarmCsv.ReadUpFilePath(); fp != "" {
			//TODO：
			alarmCsv.DeleteUpFile(fp)
		}
	}
}

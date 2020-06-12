package updata

import (
	modle "eabi_gateway/impl"
	"eabi_gateway/impl/config"
	"eabi_gateway/module"
	myLog "eabi_gateway/module/my_log"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/imroc/req"
)

//临时上传文件使用uuid命名，并且存在内存文件系统中，不长期保存，上传到服务器后直接销毁
//在sd卡中存储所有采集的数据，没有sd卡，不存储数据，按每天一个文件进行存储。
var log module.LogInterfase
var dataCsv *DataCsv
var alarmCsv *DataCsv

//Init 初始化数据和报警数据上传
func Init() {
	log = &myLog.L{}

	//TODO:文件路径靠配置
	dataCsv = NewDataCsv("updata", "/home/immm/sdb/")
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

//TODO：按上传周期主动上传。
func dataUpCycle() {
	for {
		<-time.After(time.Duration(config.SysParamDataUpCycle()) * time.Second)
		//TODO：读取需要上传的文件，
		//if fp := dataCsv.ReadUpFilePath(); fp != "" {
		//file, _ := os.Open(fp)
		file, _ := os.Open("/home/immm/sdb/updata-2020-5.csv")

		url := "http://192.168.0.194:8067/Index/handlePowerData/msgId/" + uuid.New().String() + "/msgGwId/" + config.SysParamGwId()
		fmt.Println(url)
		resp, err := req.Post(url, req.FileUpload{
			File:      file,
			FieldName: "updata",
			FileName:  "updata-2020-5.csv",
		})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(resp.ToString())

		file.Close()

		//	dataCsv.DeleteUpFile(fp)
		//}

		//TODO：读取需要上传的报警文件，
		if fp := alarmCsv.ReadUpFilePath(); fp != "" {
			//TODO：
			alarmCsv.DeleteUpFile(fp)
		}
	}
}

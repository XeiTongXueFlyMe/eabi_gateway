package updata

import (
	modle "eabi_gateway/impl"
	"eabi_gateway/impl/config"
	"eabi_gateway/module"
	myLog "eabi_gateway/module/my_log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/imroc/req"
)

//SDPATH 由外部脚本决定，只会是"/root/sdb/"
var SDPATH string = "/root/sdb/"

//临时上传文件使用uuid命名，并且存在内存文件系统中，不长期保存，上传到服务器后直接销毁
//在sd卡中存储所有采集的数据，没有sd卡，不存储数据，按每天一个文件进行存储。
var log module.LogInterfase
var dataCsv *DataCsv

//Init 初始化数据和报警数据上传
func Init() {
	log = &myLog.L{}

	dataCsv = NewDataCsv("updata", SDPATH)
	go dataUpCycle()

	return
}

//WriteUpdata 写入需要上传的数据
func WriteUpdata(updata modle.UpDataMetaInfo) {
	dataCsv.Write(updata)
}

//按上传周期主动上传。
func dataUpCycle() {
	for {
		<-time.After(time.Duration(config.SysParamDataUpCycle()) * time.Second)

		if fp := dataCsv.ReadUpFilePath(); fp != "" {
			file, _ := os.Open(fp)

		_updata:
			fn := uuid.New().String()
			url := "http://gas.elitesemicon.com.cn/Index/handlePowerData/msgId/" + fn + "/msgGwId/" + config.SysParamGwId()
			_, err := req.Post(url, req.FileUpload{
				File:      file,
				FieldName: "updata",
				FileName:  fn + ".csv",
			})
			if err != nil {
				log.PrintlnErr(err)
				time.Sleep(time.Second * 5)
				goto _updata
			}

			file.Close()
			dataCsv.DeleteUpFile(fp)
		}
	}
}

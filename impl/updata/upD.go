package updata

import (
	modle "eabi_gateway/impl"
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
)

//TEMPPATH 默认的内存文件系统的卡挂载路径
var TEMPPATH = "/dev/shm/"

//DataCsv 传感器数据写入csv文件对象
type DataCsv struct {
	one sync.Once

	name         string
	tempFileName string
	sdPath       string
}

//NewDataCsv 申请一个新的DataCsv
func NewDataCsv(name, sdpath string) *DataCsv {
	return &DataCsv{name: name, tempFileName: "temp", sdPath: sdpath}
}

//Write 写入数据
func (t *DataCsv) Write(d interface{}) {
	s := []string{}
	var info modle.UpDataMetaInfo

	switch d.(type) {
	case modle.UpDataMetaInfo:
		info, _ = d.(modle.UpDataMetaInfo)
		s = append(s, info.SourceID)
		s = append(s, info.GwID)
		s = append(s, info.SensorName)
		s = append(s, info.SensorID)
		s = append(s, fmt.Sprintf("%d", info.TimeStamp))
		s = append(s, fmt.Sprintf("%d", info.Channel))
		s = append(s, info.Unit)
		s = append(s, fmt.Sprintf("%f", info.Value))
		s = append(s, info.AlarmID)
		s = append(s, fmt.Sprintf("%f", info.AlarmParamH))
		s = append(s, fmt.Sprintf("%f", info.AlarmParamL))
		s = append(s, info.Isok)
	}

	//写入临时存储
	t.writeTemp(s)

	timestamp := time.Unix(info.TimeStamp, 0)
	tm := fmt.Sprintf("%d-%d-%dT %d:%d:%d", timestamp.Year(), timestamp.Month(), timestamp.Day(), timestamp.Hour(), timestamp.Minute(), timestamp.Second())
	s = append(s, tm)
	//写永久储存的数据
	t.writeSdb(s)

	return
}

func isFile(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return false
}

//ReadUpFilePath 读取需要上传的文件路径及其名称，如果没有返回空字符串
func (t *DataCsv) ReadUpFilePath() string {
	if isFile(TEMPPATH + t.tempFileName + ".csv") {
		defer func() { t.tempFileName = "temp" }()
		return TEMPPATH + t.tempFileName + ".csv"
	}
	return ""
}

//DeleteUpFile 删除上传文件，以便新的文件生成
func (t *DataCsv) DeleteUpFile(path string) {
	os.Remove(path)
}

func (t *DataCsv) writeSdb(s []string) {
	year, month, _ := time.Now().Date()
	fn := fmt.Sprintf(t.sdPath+t.name+"-%d-%d.csv", year, month)
	f, err := os.OpenFile(fn, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		t.one.Do(func() {
			log.PrintlnErr(err)
		})
		return
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write(s)
}

func (t *DataCsv) writeTemp(s []string) {
	if t.tempFileName == "temp" {
		t.tempFileName = uuid.New().String()
	}

	fn := TEMPPATH + t.tempFileName + ".csv"
	f, err := os.OpenFile(fn, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		t.one.Do(func() {
			log.PrintlnWarring(err)
		})
		return
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write(s)
}

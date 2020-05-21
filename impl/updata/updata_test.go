package updata

import (
	modle "eabi_gateway/impl"
	"fmt"
	"testing"
)

func TestUpdata(t *testing.T) {
	info := modle.UpDataMetaInfo{"a", "b", "c", "d", 1, 1, "f", 1}

	Init()
	df := NewDataCsv("updata", "/home/immm/sdb/")
	df.Write(info)
	fmt.Println(df.ReadUpFilePath())
	df.Write(info)
	df.DeleteUpFile(df.ReadUpFilePath())
}

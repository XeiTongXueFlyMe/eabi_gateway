package myLog

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type L struct {
	f        *os.File
	logger   *log.Logger
	fileName string
	one      sync.Once
}

func (t *L) openFile() {
	var err error
	t.one.Do(func() {
		t.fileName = ""
		t.f = nil
	})

	year, month, _ := time.Now().Date()
	if t.fileName != fmt.Sprintf("%d-%d.log", year, month) {
		if t.f != nil {
			t.f.Close()
		}

		t.fileName = fmt.Sprintf("%d-%d.log", year, month)
		t.f, err = os.OpenFile(t.fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
		if err != nil {
			panic(err)
		}
		t.logger = log.New(t.f, "", log.LstdFlags|log.Lshortfile)
	}

}

func (t *L) PrintfErr(format string, v ...interface{}) {
	t.openFile()
	str := fmt.Sprintf(format, v...)
	t.logger.Printf("[err]" + str)
}
func (t *L) PrintfWarring(format string, v ...interface{}) {
	t.openFile()
	str := fmt.Sprintf(format, v...)
	t.logger.Printf("[warring]" + str)
}
func (t *L) PrintfInfo(format string, v ...interface{}) {
	t.openFile()
	str := fmt.Sprintf(format, v...)
	t.logger.Printf("[info]" + str)
}
func (t *L) PrintlnErr(v ...interface{}) {
	t.openFile()
	str := fmt.Sprintln(v...)
	t.logger.Printf("[err]" + str)
}
func (t *L) PrintlnWarring(v ...interface{}) {
	t.openFile()
	str := fmt.Sprintln(v...)
	t.logger.Printf("[warring]" + str)
}
func (t *L) PrintlnInfo(v ...interface{}) {
	t.openFile()
	str := fmt.Sprintln(v...)
	t.logger.Printf("[info]" + str)
}

func (t *L) Printftml(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
func (t *L) Printlntml(v ...interface{}) {
	fmt.Println(v...)
}

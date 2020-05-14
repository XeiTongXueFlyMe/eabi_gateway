package lora

import (
	"sync"
	"time"

	"github.com/tarm/serial"
)

var common *serial.Port
var one sync.Once

//Open Name: "/dev/ttyUSB0", Baud: 115200,
func Open(name string, baud int, t time.Duration) error {
	var err error

	common, err = serial.OpenPort(&serial.Config{
		Name:        name,
		Baud:        baud,
		ReadTimeout: t,
	})

	return err
}

//Write []byte
func Write(b []byte) (int, error) {
	return common.Write(b)
}

//Read []byte
func Read(b []byte) (int, error) {
	return common.Read(b)
}

//Read []byte
func Close() error {
	return common.Close()
}

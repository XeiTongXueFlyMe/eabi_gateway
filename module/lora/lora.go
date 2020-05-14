package lora

import (
	"errors"
	"io"
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
	if common == nil {
		return 0, errors.New("common is null")
	}

	return common.Write(b)
}

//Read []byte
func Read(b []byte) (int, error) {
	count := 0
	buf := make([]byte, 1024)
	for {
		if common == nil {
			return 0, errors.New("common is null")
		}
		if n, err := common.Read(buf); err == io.EOF {
			return count, nil
		} else {
			if err != nil {
				return 0, err
			}
			copy(b[count:], buf[:n])
			count += n
		}
	}
}

//Read []byte
func Close() error {
	if common == nil {
		return errors.New("common is null")
	}
	return common.Close()
}

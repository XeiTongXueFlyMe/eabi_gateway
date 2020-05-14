package modbus

import (
	"fmt"
	"testing"
)

func TestByte(t *testing.T) {
	fmt.Println(ReadDeviceReg(0x01, 0x0102, 0x0304))
}

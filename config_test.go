package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMsgGwId(t *testing.T) {
	sysParamGwId("ABABABABABAB")
	assert.EqualValues(t, "ABABABABABAB", sysParamGwId())
	//"120.55.191.153:8286", "/"
	assert.EqualValues(t, "192.168.0.168:8286", sysParamHost("192.168.0.168:8286"))
	assert.EqualValues(t, "/", sysParamPath("/"))

}

func TestWriteSysParamToFile(t *testing.T) {
	assert.NoError(t, writeSysParamToFile())
}

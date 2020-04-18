package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSysParam(t *testing.T) {
	sysParamGwId("ABABABABABAB")
	assert.EqualValues(t, "ABABABABABAB", sysParamGwId())
	sysParamGwIp("192.168.0.102")
	assert.EqualValues(t, "192.168.0.102", sysParamGwIp())

	sysParamDataUpCycle(10)
	assert.EqualValues(t, 10, sysParamDataUpCycle())
	sysParamHeartCycle(15)
	assert.EqualValues(t, 15, sysParamHeartCycle())
	sysParamDataReadCycle(60)
	assert.EqualValues(t, 60, sysParamDataReadCycle())

	assert.EqualValues(t, "/", sysParamPath("/"))

	sysParamRf("ACACACACACAC", "1", "BEBEBEBEBEBE")
	ip, channel, netip := sysParamRf()
	assert.EqualValues(t, ip, "ACACACACACAC")
	assert.EqualValues(t, channel, "1")
	assert.EqualValues(t, netip, "BEBEBEBEBEBE")
}

func TestWriteSysParamToFile(t *testing.T) {
	assert.NoError(t, writeSysParamToFile())
}

func TestSysParamServerIPAndPort(t *testing.T) {
	//"120.55.191.153:8286"
	sysParamServerIPAndPort("192.168.0.168", "8286")
	i, p := sysParamServerIPAndPort()
	assert.EqualValues(t, i, "192.168.0.168")
	assert.EqualValues(t, p, "8286")
}

package main

import (
	modle "eabi_gateway/impl"
	"eabi_gateway/impl/config"
	"eabi_gateway/impl/modbus"
	rfNet "eabi_gateway/impl/rf_net"
	"runtime"
	"time"
)

var deviceIDTransmitChan chan uint8
var deviceDataTransmitChan chan modbus.RespInfo
var deviceList []modle.SensorInfo

func findDeviceAndChannel(id uint8) modle.SensorInfo {
	for _, v := range deviceList {
		if uint8(v.SensorAdder) == id {
			return v
		}
	}
	return modle.SensorInfo{}
}

//TODO;rfNet UPDATA
func deviceMarshaler() {
	runtime.Gosched()
	for {
		b := <-deviceDataTransmitChan

		//解析数据
		id, _ := b.ReadDeviceID()
		code, _ := b.ReadDeviceCode()
		_, _, bufsize, _ := b.ReadDeviceData()

		if code == 0x03 {
			v := findDeviceAndChannel(id)
			for _, c := range v.ChannelList {
				if (uint16(c.ValueAdder) + uint16(c.ValueSize)) > bufsize {
					continue
				}
				//TODO

			}
		}

		//TODO：code != 0x03

	}
}

//TODO	//判断是否报警

func waitRfData() {
	buf := make(modbus.RespInfo, 1024)
	for {
		n := rfNet.Read(buf)
		log.Printlntml("read rfdata num = ", n)
		log.Printlntml("read rfdata  = ", buf[:n])

		b := buf[:n]
		if b.IsSupportModbusForm() {
			id, _ := b.ReadDeviceID()

			select {
			case deviceIDTransmitChan <- id:
			default:
			}

			deviceDataTransmitChan <- b[:n]
		} else {
			log.PrintfWarring("rfNet invalid data:%x", b)
		}
	}
}

//采用轮训模式
func modbusDataTransmit() {

	//每次轮训都重新读取传感器配置和报警配置
	for {
		t := time.Now().Unix()
		deviceList = config.ReadSensorConfig()

		//modbusDataTransmit 由于lora网络传输数据较慢，所以通过接受超时来控制发送速率
		for _, v := range deviceList {
			rfNet.Send(modbus.ReadDeviceReg(uint8(v.SensorAdder), uint16(v.DataAdder), uint16(v.DataSize)))
			for {
				//等待数据返回，或超时
				select {
				case <-time.After(time.Second * 5):
					goto _continue
				case id := <-deviceIDTransmitChan:
					if id == uint8(v.SensorAdder) {
						goto _continue
					}
				}
			}
		_continue:
		}

		//按dataReadCycle周期休眠，需要判断读取数据总共用时多少
		if int64(config.SysParamDataReadCycle()) > (time.Now().Unix() - t) {
			time.Sleep(time.Duration(int64(config.SysParamDataReadCycle()) - (time.Now().Unix() - t)))
		}
	}
}

func rfInit() {
	deviceIDTransmitChan = make(chan uint8)
	deviceDataTransmitChan = make(chan modbus.RespInfo)

	go waitRfData()
	go modbusDataTransmit()
	go deviceMarshaler()
}

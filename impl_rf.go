package main

import (
	modle "eabi_gateway/impl"
	busNet "eabi_gateway/impl/Industrial_bus"
	"eabi_gateway/impl/config"
	"eabi_gateway/impl/modbus"
	rfNet "eabi_gateway/impl/rf_net"
	"eabi_gateway/impl/updata"
	"fmt"
	"runtime"
	"time"

	"github.com/google/uuid"
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

type alarmFactry struct {
	AlarmID string
	Isok    string
}

var AlarmMap map[string]alarmFactry

func deviceMarshaler() {
	runtime.Gosched()
	for {
		b := <-deviceDataTransmitChan

		//解析数据modbus
		id, _ := b.ReadDeviceID()
		code, _ := b.ReadDeviceCode()
		_, buf, bufsize, _ := b.ReadDeviceData()

		//FIXME:暂时只处理　0x03
		if code == 0x03 {
			v := findDeviceAndChannel(id)

			for _, c := range v.ChannelList {
				if c.Channel == 0 {
					continue
				}

				adder := uint16((c.Channel - 1) * 4)
				if adder >= bufsize {
					continue
				}

				value := ByteToFloat32(buf[adder:])

				//写rfnetinfo
				rfNet.WriteInfo(v.SensorID, v.SensorName, "v0.0.0", "v0.0.0", fmt.Sprint(c.Channel))

				//判断是否存在报警历史，如果没有，就新建一个报警,报警状态ok
				if _, ok := AlarmMap[v.SensorID+fmt.Sprintf("%d", c.Channel)]; !ok {
					AlarmMap[v.SensorID+fmt.Sprintf("%d", c.Channel)] = alarmFactry{
						AlarmID: uuid.New().String(),
						Isok:    "ok",
					}
				}

				//判断是否报警
				alarm := AlarmMap[v.SensorID+fmt.Sprintf("%d", c.Channel)]

				if config.IsAlarm(v.SensorID, c.Channel, value) {
					if alarm.Isok != "alarm" {
						alarm.AlarmID = uuid.New().String()
					}
					alarm.Isok = "alarm"

				} else {
					if alarm.Isok != "ok" {
						alarm.AlarmID = uuid.New().String()
					}
					alarm.Isok = "ok"
				}

				AlarmMap[v.SensorID+fmt.Sprintf("%d", c.Channel)] = alarm

				//读取报警参数
				l, h, err := config.ReadAlarmParamLH(v.SensorID, c.Channel)
				if err != nil {
					l = 0
					h = 0
				}

				//写上传文件
				d := modle.UpDataMetaInfo{
					SourceID:    uuid.New().String(),
					GwID:        config.SysParamGwId(),
					TimeStamp:   time.Now().Unix(),
					SensorID:    v.SensorID,
					SensorName:  v.SensorName,
					Channel:     uint32(c.Channel),
					Unit:        c.ValueType,
					Value:       value,
					AlarmID:     alarm.AlarmID,
					AlarmParamH: h,
					AlarmParamL: l,
					Isok:        alarm.Isok,
				}
				updata.WriteUpdata(d)
			}
		}
	}
}

func waitRfData() {
	buf := make(modbus.RespInfo, 1024)
	for {
		n := rfNet.Read(buf)
		log.Printlntml("read rfdata num = ", n)
		log.Printlntml("read rfdata  = ", buf[:n])

		//判断是否是485通信
		if config.SysHardware() != "lora" {
			continue
		}

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

func waitBusData() {
	buf := make(modbus.RespInfo, 1024)
	for {
		n := busNet.Read(buf)
		log.Printlntml("read rfdata num = ", n)
		log.Printlntml("read rfdata  = ", buf[:n])

		//判断是否是485通信
		if config.SysHardware() != "485" {
			continue
		}

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

func sendAdapterData(v modle.AdapterInfo) {

	for _, chanV := range v.ChannelSetList {
		buf := make([]byte, 0)
		buf = append(buf, byte((uint16(chanV.UbgAdder)&0xff00)>>8))
		buf = append(buf, byte(chanV.UbgAdder&0xff))
		buf = append(buf, Float32ToByte(float32(chanV.RangeLow))...)
		buf = append(buf, Float32ToByte(float32(chanV.RangeHigh))...)
		buf = append(buf, Float32ToByte(float32(chanV.K))...)
		buf = append(buf, Float32ToByte(float32(chanV.B))...)
		buf = append(buf, byte(uint16(chanV.Period&0xff00)>>8))
		buf = append(buf, byte(chanV.Period&0xff))
		buf = append(buf, byte(uint16(chanV.ChannelEn&0xff00)>>8))
		buf = append(buf, byte(chanV.ChannelEn&0xff))
		buf = append(buf, byte(uint16(chanV.ModbusAdder&0xff00)>>8))
		buf = append(buf, byte(chanV.ModbusAdder&0xff))
		buf = append(buf, byte(uint16(chanV.Bufse&0xff00)>>8))
		buf = append(buf, byte(chanV.Bufse&0xff))

		//判断 lora通信，485通信
		switch config.SysHardware() {
		case "lora":
			rfNet.Send(modbus.WriteDeviceReg(uint8(v.SensorAdder), uint16(chanV.Channel)*20+10000, 13, 26, buf[:]))
		case "485":
			busNet.Send(modbus.WriteDeviceReg(uint8(v.SensorAdder), uint16(chanV.Channel)*20+10000, 13, 26, buf[:]))
		default:
			continue
		}

		//等待数据返回，或超时
		select {
		case <-time.After(time.Second * 5):
			log.PrintlnErr("send adapter data timeout(5s):", modbus.WriteDeviceReg(uint8(v.SensorAdder), uint16(chanV.Channel)*20+10000, 13, 26, buf[:]))
			goto _continue
		case id := <-deviceIDTransmitChan:
			if id == uint8(v.SensorAdder) {
				goto _continue
			}
			log.PrintfErr("modbus return id = %d  != %d", id, v.SensorAdder)
		}

	_continue:
	}
}

//采用轮训模式
func modbusDataTransmit() {

	//每次轮训都重新读取传感器配置和报警配置
	//modbusDataTransmit 由于lora，485网络传输数据较慢，所以通过接受超时来控制发送速率
	for {
		t := time.Now().Unix()
		deviceList = config.ReadSensorConfig()

		//轮训发送数据数据
		for _, v := range deviceList {
			//判断 lora通信，485通信
			switch config.SysHardware() {
			case "lora":
				rfNet.Send(modbus.ReadDeviceReg(uint8(v.SensorAdder), uint16(v.DataAdder), uint16(v.DataSize)))
			case "485":
				busNet.Send(modbus.ReadDeviceReg(uint8(v.SensorAdder), uint16(v.DataAdder), uint16(v.DataSize)))
			default:
				continue
			}

			for {
				//等待数据返回，或超时
				select {
				case <-time.After(time.Second * 5):
					log.PrintfErr("modbus wait sensorId = %d return timeout(5s)", v.SensorAdder)
					goto _continue
				case id := <-deviceIDTransmitChan:
					if id == uint8(v.SensorAdder) {
						goto _continue
					}
					log.PrintfErr("modbus return id = %d  != %d", id, v.SensorAdder)
				}
			}
		_continue:
			select { //适配器配置数据
			case v := <-adapterSendRfDataChannel:
				sendAdapterData(v)
			default:
			}
		}

		//按dataReadCycle周期休眠，需要判断读取数据总共用时多少
		if int64(config.SysParamDataReadCycle()) > (time.Now().Unix() - t) {
			time.Sleep(time.Duration(int64(config.SysParamDataReadCycle())-(time.Now().Unix()-t)) * time.Second)
		}
	}
}

func rfInit() {
	AlarmMap = make(map[string]alarmFactry)
	deviceIDTransmitChan = make(chan uint8)
	deviceDataTransmitChan = make(chan modbus.RespInfo)
	//lora
	go waitRfData()
	//485
	go waitBusData()

	go modbusDataTransmit()
	go deviceMarshaler()
}

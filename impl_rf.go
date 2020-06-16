package main

import (
	modle "eabi_gateway/impl"
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

		//TODO:暂时只处理　0x03
		if code == 0x03 {
			v := findDeviceAndChannel(id)

			for _, c := range v.ChannelList {
				adder := uint16(c.Channel * 4)
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
		var buf [20]byte
		// regNum uint16, byteSize uint8, buf []byte
		fmt.Println(modbus.WriteDeviceReg(uint8(v.SensorAdder), uint16(chanV.Channel)*20+10000, 10, buf[:]))
		rfNet.Send(modbus.WriteDeviceReg(uint8(v.SensorAdder), uint16(chanV.Channel)*20+10000, 10, buf[:]))
		//等待数据返回，或超时
		select {
		case <-time.After(time.Second * 5):
			goto _continue
		case id := <-deviceIDTransmitChan:
			if id == uint8(v.SensorAdder) {
				goto _continue
			}
		}

	_continue:
	}
}

//采用轮训模式
func modbusDataTransmit() {

	//每次轮训都重新读取传感器配置和报警配置
	//modbusDataTransmit 由于lora网络传输数据较慢，所以通过接受超时来控制发送速率
	for {
		t := time.Now().Unix()
		deviceList = config.ReadSensorConfig()

		//轮训发送数据数据
		for _, v := range deviceList {
			fmt.Println(modbus.ReadDeviceReg(uint8(v.SensorAdder), uint16(v.DataAdder), uint16(v.DataSize)))
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

				//TODO:等待适配器配置数据
				select {
				case v := <-adapterSendRfDataChannel:
					sendAdapterData(v)
				default:
				}
			}
		_continue:
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

	go waitRfData()
	go modbusDataTransmit()
	go deviceMarshaler()
}

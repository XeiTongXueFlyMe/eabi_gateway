package modle

type StdReq struct {
	MsgType      string `json:"msgType"`
	MsgID        string `json:"msgId"`
	MsgGwID      string `json:"msgGwId"`
	MsgTimeStamp int64  `json:"msgTimeStamp"`
	MsgParam     string `json:"msgParam"`
}

type StdResp struct {
	MsgType      string `json:"msgType"`
	MsgID        string `json:"msgId"`
	MsgGwID      string `json:"msgGwId"`
	MsgTimeStamp int64  `json:"msgTimeStamp"`
	MsgParam     string `json:"msgParam"`
	MsgResp      string `json:"msgResp"`
}

type GatewayParmResp struct {
	MsgType      string `json:"msgType"`
	MsgID        string `json:"msgId"`
	MsgGwID      string `json:"msgGwId"`
	MsgTimeStamp int64  `json:"msgTimeStamp"`
	MsgParam     string `json:"msgParam"`
	MsgResp      string `json:"msgResp"`
	GwID         string `json:"gwId"`
	GwIP         string `json:"gwIP"`
	ServerIP     string `json:"serverIP"`
	ServerPort   string `json:"serverPort"`
	RfID         string `json:"rfId"`
	RfChannel    string `json:"rfChannel"`
	RfNetID      string `json:"rfNetId"`
	DataUpCycle  int    `json:"dataUpCycle"`
	HeartCycle   int    `json:"heartCycle"`
}

type RfNetInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	SoftwareVer string `json:"softwareVer"`
	HardwareVer string `json:"hardwareVer"`
	Channel1    string `json:"channel_1"`
	Channel2    string `json:"channel_2"`
	Channel3    string `json:"channel_3"`
	Channel4    string `json:"channel_4"`
	Channel5    string `json:"channel_5"`
	Channel6    string `json:"channel_6"`
	Channel7    string `json:"channel_7"`
	Channel8    string `json:"channel_8"`
}

type RfNetInfoResp struct {
	MsgType      string      `json:"msgType"`
	MsgID        string      `json:"msgId"`
	MsgGwID      string      `json:"msgGwId"`
	MsgTimeStamp int64       `json:"msgTimeStamp"`
	MsgParam     string      `json:"msgParam"`
	MsgResp      string      `json:"msgResp"`
	RfNetNum     int         `json:"rfNetNum"`
	RfNetInfo    []RfNetInfo `json:"rfNetInfo"`
}

type ChanneInfo struct {
	Channel   int    `json:"channel"`   //通道号
	ValueType string `json:"valueType"` //通道值类型
}

type SensorInfo struct {
	SensorID    string       `json:"sensorId"`
	SensorName  string       `json:"sensorName"`
	SensorAdder int          `json:"sensorAdder"`
	DataAdder   int          `json:"dataAdder"`
	DataSize    int          `json:"dataSize"`
	ChannelList []ChanneInfo `json:"channelList"`
}

//AdapterChanInfo 适配器通道信息
type AdapterChanInfo struct {
	Channal     string `json:"channal"`
	UbgAdder    string `json:"ubgAdder"`    //UBG地址设定
	RangeLow    string `json:"rangeLow"`    //零量程
	RangeHigh   string `json:"rangeHigh"`   //满量程
	K           string `json:"k"`           //修正系数K
	B           string `json:"b"`           //修正系数B
	Period      string `json:"period"`      //传感器周期
	ChannelEn   string `json:"channelEn"`   //通道使能
	ModbusAdder string `json:"modbusAdder"` //通道对应MODBUS地址
	Bufse       string `json:"bufse"`       //通道对应数据长度
}

//AdapterInfo 适配器信息
type AdapterInfo struct {
	SensorID       string            `json:"sensorId"`
	SensorAdder    int               `json:"sensorAdder"`
	ChannelSetList []AdapterChanInfo `json:"channelSetList"`
}

type AdapterInfoResp struct {
	MsgType      string      `json:"msgType"`
	MsgID        string      `json:"msgId"`
	MsgGwID      string      `json:"msgGwId"`
	MsgTimeStamp int64       `json:"msgTimeStamp"`
	MsgParam     string      `json:"msgParam"`
	MsgResp      string      `json:"msgResp"`
	AdapterInfo  AdapterInfo `json:"adapterInfo"`
}

type AdapterInfoReq struct {
	MsgType      string      `json:"msgType"`
	MsgID        string      `json:"msgId"`
	MsgGwID      string      `json:"msgGwId"`
	MsgTimeStamp int64       `json:"msgTimeStamp"`
	MsgParam     string      `json:"msgParam"`
	AdapterInfo  AdapterInfo `json:"adapterInfo"`
}

type SensorInfoReq struct {
	MsgType       string       `json:"msgType"`
	MsgID         string       `json:"msgId"`
	MsgGwID       string       `json:"msgGwId"`
	MsgTimeStamp  int64        `json:"msgTimeStamp"`
	MsgParam      string       `json:"msgParam"`
	SensorListNum int          `json:"sensorListNum"`
	SensorList    []SensorInfo `json:"sensorList"`
}

type SensorInfoResp struct {
	MsgType       string       `json:"msgType"`
	MsgID         string       `json:"msgId"`
	MsgGwID       string       `json:"msgGwId"`
	MsgTimeStamp  int64        `json:"msgTimeStamp"`
	MsgParam      string       `json:"msgParam"`
	MsgResp       string       `json:"msgResp"`
	SensorListNum int          `json:"sensorListNum"`
	SensorList    []SensorInfo `json:"sensorList"`
}

type AlarmInfo struct {
	SensorID    string  `json:"sensorId"`
	Channel     int     `json:"channel"`
	AlarmValueL float32 `json:"alarmValue_l"`
	AlarmValueH float32 `json:"alarmValue_h"`
}

type AlarmInfoReq struct {
	MsgType      string      `json:"msgType"`
	MsgID        string      `json:"msgId"`
	MsgGwID      string      `json:"msgGwId"`
	MsgTimeStamp int64       `json:"msgTimeStamp"`
	MsgParam     string      `json:"msgParam"`
	AlarmListNum int         `json:"alarmListNum"`
	AlarmList    []AlarmInfo `json:"alarmList"`
}

type AlarmInfoResp struct {
	MsgType      string      `json:"msgType"`
	MsgID        string      `json:"msgId"`
	MsgGwID      string      `json:"msgGwId"`
	MsgTimeStamp int64       `json:"msgTimeStamp"`
	MsgParam     string      `json:"msgParam"`
	MsgResp      string      `json:"msgResp"`
	AlarmListNum int         `json:"alarmListNum"`
	AlarmList    []AlarmInfo `json:"alarmList"`
}

//UpDataMetaInfo 传感器上传的元信息
type UpDataMetaInfo struct {
	SourceID    string
	GwID        string
	SensorName  string
	SensorID    string
	TimeStamp   int64
	Channel     uint32
	Unit        string //单位
	Value       float32
	AlarmID     string
	AlarmParamH float32
	AlarmParamL float32
	Isok        string
}

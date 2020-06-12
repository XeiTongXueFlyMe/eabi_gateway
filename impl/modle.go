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
	SourceID   string
	GwID       string
	SensorName string
	SensorID   string
	TimeStamp  int64
	Channel    uint32
	Unit       string //单位
	Value      float32
}

//AlarmMetaInfo 数据报警的元信息
type AlarmMetaInfo struct {
	AlarmID     string
	GwID        string
	SensorName  string
	SensorID    string
	TimeStamp   int64
	Channel     uint32
	AlarmParamH float32
	AlarmParamL float32
	Param       float32
	Isok        string
}

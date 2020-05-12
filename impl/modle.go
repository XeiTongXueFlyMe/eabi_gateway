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
	Channel    int    `json:"channel"`    //通道号
	AAdder     int    `json:"a_Adder"`    //标定系数ａ地址
	ASize      int    `json:"a_Size"`     //标定系数ａ大小，浮点数
	BAdder     int    `json:"b_Adder"`    //标定系数b地址
	BSize      int    `json:"b_Size"`     //标定系数b大小，浮点数
	WorkAdder  int    `json:"work_Adder"` //工作状态地址
	WorkSize   int    `json:"work_Size"`  //工作状态地址大小
	ValueAdder int    `json:"valueAdder"` //通道值地址
	ValueSize  int    `json:"valueSize"`  //通道值大小
	ValueType  string `json:"valueType"`  //通道值类型
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

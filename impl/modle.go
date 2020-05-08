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
	RfNetNum     string      `json:"rfNetNum"`
	RfNetInfo    []RfNetInfo `json:"rfNetInfo"`
}

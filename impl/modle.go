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
	MsgType       string `json:"msgType"`
	MsgID         string `json:"msgId"`
	MsgGwID       string `json:"msgGwId"`
	MsgTimeStamp  int64  `json:"msgTimeStamp"`
	MsgParam      string `json:"msgParam"`
	MsgResp       string `json:"msgResp"`
	GwID          string `json:"gwId"`
	GwIP          string `json:"gwIP"`
	ServerIP      string `json:"serverIP"`
	ServerPort    string `json:"serverPort"`
	Hardware      string `json:"hardware"`
	RfID          string `json:"rfId"`
	RfChannel     string `json:"rfChannel"`
	RfNetID       string `json:"rfNetId"`
	DataUpCycle   int    `json:"dataUpCycle"`
	HeartCycle    int    `json:"heartCycle"`
	DataReadCycle int    `json:"dataReadCycle"`
}

type RfNetInfo struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	SoftwareVer string   `json:"softwareVer"`
	HardwareVer string   `json:"hardwareVer"`
	Channel     []string `json:"channel"`
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
	CName     string `json:"cName"`     //通道别名
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
	Channel     int     `json:"channel"`
	UbgAdder    int     `json:"ubgAdder"`    //UBG地址设定
	RangeLow    float64 `json:"rangeLow"`    //零量程
	RangeHigh   float64 `json:"rangeHigh"`   //满量程
	K           float64 `json:"k"`           //修正系数K
	B           float64 `json:"b"`           //修正系数B
	Period      int     `json:"period"`      //传感器周期
	ChannelEn   int     `json:"channelEn"`   //通道使能
	ModbusAdder int     `json:"modbusAdder"` //通道对应MODBUS地址
	Bufse       int     `json:"bufse"`       //通道对应数据长度
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
	SensorID     string      `json:"sensorId"`
	AdapterInfo  AdapterInfo `json:"adapterInfo"`
}

type AdapterInfoReq struct {
	MsgType      string      `json:"msgType"`
	MsgID        string      `json:"msgId"`
	MsgGwID      string      `json:"msgGwId"`
	MsgTimeStamp int64       `json:"msgTimeStamp"`
	MsgParam     string      `json:"msgParam"`
	SensorID     string      `json:"sensorId"`
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

//AutoFeeding 自动投料相关数据
type AutoFeeding struct {
	TimerOneFlag       bool  `json:"timerOneFlag"`
	TimerOneHour       int64 `json:"timerOneHour"`
	TimerOneMinutes    int64 `json:"timerOneMinutes"`
	TimerTwoFlag       bool  `json:"timerTwoFlag"`
	TimerTwoHour       int64 `json:"timerTwoHour"`
	TimerTwoMinutes    int64 `json:"timerTwoMinutes"`
	TimeFlag           bool  `json:"timeFlag"`
	TimeStamp          int64 `json:"timeStamp"`
	IncrementFlag      bool  `json:"incrementFlag"`
	Increment          int64 `json:"increment"`
	IntervalTimeFlag   bool  `json:"intervalTimeFlag"`
	ITimerDay          int64 `json:"iTimerDay"`
	ITimerHour         int64 `json:"iTimerHour"`
	ITimerMinutes      int64 `json:"iTimerMinutes"`
	IntervalExitFlag   bool  `json:"intervalExitFlag"`
	TimeModuleFlag     bool  `json:"timeModuleFlag"`
	TimerOneModuleFlag bool  `json:"timerOneModuleFlag"`
	TimerTwoModuleFlag bool  `json:"timerTwoModuleFlag"`
	TimeModuleExitFlag bool  `json:"timeModuleExitFlag"`
	IntervalMoudleFlag bool  `json:"intervalMoudleFlag"`
}

type AutoFeedingReq struct {
	MsgType      string      `json:"msgType"`
	MsgID        string      `json:"msgId"`
	MsgGwID      string      `json:"msgGwId"`
	MsgTimeStamp int64       `json:"msgTimeStamp"`
	MsgParam     string      `json:"msgParam"`
	AutoFeeding  AutoFeeding `json:"autoFeeding"`
}

type AutoFeedingResp struct {
	MsgType      string      `json:"msgType"`
	MsgID        string      `json:"msgId"`
	MsgGwID      string      `json:"msgGwId"`
	MsgTimeStamp int64       `json:"msgTimeStamp"`
	MsgParam     string      `json:"msgParam"`
	MsgResp      string      `json:"msgResp"`
	AutoFeeding  AutoFeeding `json:"autoFeeding"`
}

//MaterialNum 加液量参数
type MaterialNum struct {
	Yg      float32 `json:"yg"`
	Dti     float32 `json:"dti"`
	Dto     float32 `json:"dto"`
	Dci     float32 `json:"dci"`
	Ht      float32 `json:"ht"`
	Hr      float32 `json:"hr"`
	Pco     float32 `json:"pco"`
	Qg      float32 `json:"qg"`
	Qw      float32 `json:"qw"`
	Twh     float32 `json:"twh"`
	Tr      float32 `json:"tr"`
	Hour    float32 `json:"hour"`
	Minutes float32 `json:"minutes"`
	OldPt   float32 `json:"oldPt"`
	OldPc   float32 `json:"oldPc"`
	Ygjpd   float32 `json:"ygjpd"`
}

type MaterialNumReq struct {
	MsgType      string      `json:"msgType"`
	MsgID        string      `json:"msgId"`
	MsgGwID      string      `json:"msgGwId"`
	MsgTimeStamp int64       `json:"msgTimeStamp"`
	MsgParam     string      `json:"msgParam"`
	MaterialNum  MaterialNum `json:"materialNum"`
}

type MaterialNumResp struct {
	MsgType      string      `json:"msgType"`
	MsgID        string      `json:"msgId"`
	MsgGwID      string      `json:"msgGwId"`
	MsgTimeStamp int64       `json:"msgTimeStamp"`
	MsgParam     string      `json:"msgParam"`
	MsgResp      string      `json:"msgResp"`
	MaterialNum  MaterialNum `json:"materialNum"`
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

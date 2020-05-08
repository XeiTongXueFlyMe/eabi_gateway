package main

type rfNetInfo struct {
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

var rfNetInfoMap map[string]rfNetInfo
var rfNetChan chan []byte

func rfNetInfoInit() {
	rfNetInfoMap = make(map[string]rfNetInfo)
	//TODO:

	//createMsgField("")

}

func cleanRfNetInfo() {
	for k := range rfNetInfoMap {
		delete(rfNetInfoMap, k)
	}
}

func readRfNetInfo() {
	//TODO:sendNetData()
}

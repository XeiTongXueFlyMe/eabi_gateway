package updata

//临时上传文件使用uuid命名，并且存在内存文件系统中，不长期保存，上传到服务器后直接销毁
//在sd卡中存储所有采集的数据，没有sd卡，不能启动系统，按每天一个文件进行存储。

//TODO：主动上传，通过上传周期。

//TODO:写入数据接口（传感器数据，报警数据）

//Updata 数据和报警数据上传对象
type Updata struct {
	udata  upD
	updata alarmD
}

//Init 初始化数据和报警数据上传
func Init() {
	//TODO
	return
}

//WriteUpdata 写入需要上传的数据
func (t *Updata) WriteUpdata() {
	//TODO
	return
}

//WriteAlarmdata 写入需要上传的报警数据
func (t *Updata) WriteAlarmdata() {
	//TODO
	return
}

var common Updata

//WriteUpdata 写入需要上传的数据
func WriteUpdata() {
	//TODO
	common.WriteUpdata()
	return
}

//WriteAlarmdata 写入需要上传的报警数据
func WriteAlarmdata() {
	//TODO
	common.WriteAlarmdata()
	return
}

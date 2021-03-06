
## 通信采用websocket与http
    1. 主要使用 websocket 
    2. 同时兼容 http
    3. 采用json作为数据序列化
    4. 后期加密采用 wss https
    5. 服务记录账户操作日志 

## 关于msg结构

|key field|type|说明|
|--|--|--|
|msgId|string|采用google的uuid|
|gwId|string|设备唯一设别id|
|msgTimeStamp|string|Unix time stamp|
|type|string|(GET,POST,PUT,DELETE)。(GET,PUT,DELETE)幂等操作，POST不做幂等操作|
|msgParam|string|消息身份描述 eg:ping, dataUp, ....|
|msgResp|string|参照　msg_resp　取值|
|...|...|...|


## 关于　msgResp　取值
* ok
* errorIdentity (身份认证错误)
* errorTimeout (消息接受端认为消息过时，返回)
* errorSql
* errorSys
* errorUnkonw
* .......................　

# 1. ping
* 接口说明
>  服务器与嵌入式相互验证连接是否成功

**websocket example:**

* req:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"ping",
>}
>```
* resp:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"pong",
>"msgResp":"参照　msgResp　取值"
>}
>```

**http example**
* 接口地址
>GET /common/ping/{msgId}/{msgGwId}/{msgTimeStamp}

* req

* resp
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"pong",
>"msgResp":"参照　msgResp　取值"
>}
>```

# 2.1 gatewayParam
* 接口说明
>  获取网关参数
**websocket example:**
* req:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"gatewayParam",
>}
>```
* resp:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"gatewayParam",
>"msgResp":"ok",
>"gwId":"AFAF73BADCF6", 
>"gwIP":"192.168.0.2",
>"serverIP":"10.56.1.34",
>"serverPort":"8287",
>"rfId":"AFAF73BADCED",
>"rfChannel":"1",
>"rfNetId":"AFAF73BADCAA",
>"dataUpCycle":1, //单位秒
>"heartCycle":1  //单位秒
>}
>```

# 2.2 gatewayParam
* 接口说明
>  设置网关参数

**websocket example:**
* req:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"gatewayParam",
>"gwId":"AFAF73BADCF6", 
>"serverIP":"10.56.1.34",
>"serverPort":"8287",
>"rfId":"AFAF73BADCED",
>"rfChannel":"1", 
>"rfNetId":"AFAF73BADCAA",
>"dataUpCycle":1, //单位秒
>"heartCycle":1  //单位秒
>}
>```
* resp:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"gatewayParam",
>"msgResp":"ok"
>}
>```
# 2.3 gatewayParam
* 接口说明
>  删除网关参数，网关参数会恢复出厂默认

**websocket example:**
* req:
>```json
>{
>"msgType":"DELETE",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"gatewayParam",
>}
>```
* resp:
>```json
>{
>"msgType":"DELETE",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"gatewayParam",
>"msgResp":"ok"
>}
>```

# 3.1 rfNetInfo
* 接口说明
>  射频层已入网设备信息

**websocket example:**
* req:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"rfNetInfo",
>}
>```
* resp:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"rfNetInfo",
>"msgResp":"ok",
>"rfNetNum":10,
>"rfNetInfo":[
>   {
>   "id":"AFAF73BADCF6",
>   "name":"东一层２号气表",
>   "softwareVer":"1.0.1",
>   "hardwareVer":"1.0.2",
>    "channel": [
>        "channel_0",
>        "channel_1",
>        "channel_2",
>        "channel_3"
>    ]  
>   },
>   {
>   "id":"AFAF73BADCF7",
>   "name":"东二层２号气表",
>   "softwareVer":"1.0.1",
>   "hardwareVer":"1.0.2",
>    "channel": [
>        "channel_0",
>        "channel_1",
>        "channel_2",
>        "channel_3"
>    ]
>   }
> ]
>}
>```

# 3.2 rfNetInfo
* 接口说明
>  清除射频层已入网设备信息，等待重新收集。

**websocket example:**
* req:
>```json
>{
>"msgType":"DELETE",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"rfNetInfo",
>}
>```
* resp:
>```json
>{
>"msgType":"DELETE",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"rfNetInfo",
>"msgResp":"ok",
>}
>```

# 4 abnormalAlarm
* 接口说明
>  业务报警,嵌入式按每天对报警信息归类到一个文件，当产生报警消息时，及时上传当天的报警文件到服务器

>POST /common/abnormalAlarm/{msgId}/{msgGwId}/
* req
//xxx.csv

* resp
>```json
>{
>"msgType":"POST",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"abnormalAlarm",
>"msgResp":"参照　msgResp　取值"
>}
>```

# 5 sensorData
* 接口说明
>  上传传感器数据，通过POST上传文件，采用csv格式的文件,遵循上传周期配置，默认1分钟
>  嵌入式按每天对报警信息归类到一个文件，系统不做实时读取，因为网关获取数据需要时间过长。

>POST /common/sensorData/{msgId}/{msgGwId}/
* req
//xxx.csv

* resp
>```json
>{
>"msgType":"POST",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"sensorData",
>"msgResp":"参照　msgResp　取值"
>}
>```

# 6.1 获取传感器配置
* 接口说明
>  读取网关传感器配置

**websocket example:**
* req:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"sensorInfo",
>}
>```
* resp:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"sensorInfo",
>"msgResp":"ok",
>"sensorListNum": 1,
>"sensorList":
>   [
>      {
>            "sensorId":"AFAF73BADCF6",
>            "sensorName":"东１层２号气表",
>            "sensorAdder":0,
>            "dataAdder":0,
>            "dataSize":８,
>            "channelList":
>            [
>                {
>                 "channel":1,
>                 "cName":"xxx",//通道名称
>                 "valueType":"压力(Psi)",
>                }, 
>                {
>                 "channel":2,
>                 "cName":"xxx",//通道名称
>                 "valueType":"距离(mm)",
>                }, 
>            ]
>      },
>   ] 
>}
>```
# 6.２ 设置传感器配置
* 接口说明
>  设置传感器配置，为了遵循modbus协议

**websocket example:**
* req:
>```json
>{
>"msgType":"PUT",　
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"sensorInfo",
>"sensorListNum": 1,
>"sensorList":
>   [
>      {
>            "sensorId":"AFAF73BADCF6",
>            "sensorName":"东１层２号气表",
>            "sensorAdder":1,　//modbus地址
>            "dataAdder":0,
>            "dataSize":8,
>            "channelList":
>            [
>                {
>                 "channel":1,
>                 "cName":"xxx",//通道名称
>                 "valueType":"压力(Psi)",
>                }, 
>                {
>                 "channel":2,
>                 "cName":"xxx",//通道名称
>                 "valueType":"距离(mm)",
>                }, 
>            ]
>      },
>    ]    
>}
>```
* resp:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"sensorInfo",
>"msgResp":"ok",
>}
>```
# 7.１ 读取报警配置
**websocket example:**
* req:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"alarmConfig",
>}
>```

* resp:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"alarmConfig",
>"msgResp":"ok",
>"alarmListNum": 1,
>"alarmList":
>[
>   {
>      "sensorId":"AFAF73BADCF6",
>      "channel":1,
>      "alarmValue_l":10,
>      "alarmValue_h":100,
>   },
>]    
>}
>```

# 7.2 设置报警配置
**websocket example:**
* req:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"alarmConfig",
>"alarmListNum": 1,
>"alarmList":
>[
>   {
>      "sensorId":"AFAF73BADCF6",
>      "channel":1,
>      "alarmValue_l":10,
>      "alarmValue_h":100,
>   },
>]    
>}
>```
* resp:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"alarmConfig",
>"msgResp":"ok",
>}
>```


# 9 异常处理及数据对账
* 接口说明
> 网管自身异常日志，

>POST /common/getawayLog/{msgId}/{msgGwId}/
* req
//xxx.log

* resp
>```json
>{
>"msgType":"POST",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"gatewayLog",
>"msgResp":"参照　msgResp　取值"
>}
>```

＃　10　适配器配置
* 接口说明
>  设置适配器配置，使其可以转发三方仪器仪表数据

**websocket example:**
* req:
>```json
>{
>"msgType":"PUT",　
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"adapter",
>"sensorId":"AFAF73BADCF6",
>"adapterInfo":
>{
>"sensorId":"AFAF73BADCF6",
>"sensorAdder":1,　//modbus地址
>"channelSetList":
>[
>   {
>       "chennal":1,
>       "ubgAdder":0,//UBG地址设定
>       "rangeLow":0.0, //零量程
>       "rangeHigh":100.0,//满量程
>       "k":1.0,//修正系数K
>       "b":1.0,//修正系数B
>       "period":10,//传感器周期
>       "channelEn":0,//通道使能
>       "modbusAdder":25,//通道对应MODBUS地址
>       "bufse":2, //通道对应数据长度
>   },
>]
>}
>}
>```
* resp:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"adapter",
>"msgResp":"ok",
>}
>```

＃　10.1　读取适配器配置
* req:
>```json
>{
>"msgType":"GET",　
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"adapter",
>"sensorId":"AFAF73BADCF6",
>```
* resp:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"adapter",
>"msgResp":"ok",
>"sensorId":"AFAF73BADCF6",
>"adapterInfo":
>{
>"sensorId":"AFAF73BADCF6",
>"sensorAdder":1,　//modbus地址
>"channelSetList":
>[
>   {
>       "channel":1,
>       "ubgAdder":0,//UBG地址设定
>       "rangeLow":0.0, //零量程
>       "rangeHigh":100.0,//满量程
>       "k":1.0,//修正系数K
>       "b":1.0,//修正系数B
>       "period":10,//传感器周期
>       "channelEn":0,//通道使能
>       "modbusAdder":25,//通道对应MODBUS地址
>       "bufse":2, //通道对应数据长度
>   },
>]
>}
>}
>}
>```

# 11 传感器别名设置
**websocket example:**
* req:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"sensorName",
>"sensorId":"AFAF73BADCF6",
>"name":"foobar",
>}
>```
* resp:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"sensorName",
>"msgResp":"ok",
>}
>```

# 12.1 自动投料相关的数据
* req:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"autoFeeding",
>"autoFeeding":
>{
>"timerOneFlag":true,
>"timerOneHour":1,
>"timerOneMinutes":30,
>"timerTwoFlag":true,
>"timerTwoHour":1,
>"timerTwoMinutes":30,
>"timeFlag":true,
>"timeStamp":1586162503,
>"incrementFlag":true,
>"increment":100,
>"intervalTimeFlag":true,
>"iTimerDay":1,
>"iTimerHour":1,
>"iTimerMinutes":30,
>"intervalExitFlag":true,
>"timeModuleFlag":true,
>"timerOneModuleFlag":true,
>"timerTwoModuleFlag":true,
>"timeModuleExitFlag":true,
>"intervalMoudleFlag":true,
>}
>}
>```
* resp:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"autoFeeding",
>"msgResp":"ok",
>}
>```

# 12.1 自动投料相关的数据
* req:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"autoFeeding",
>}
>```
* resp:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"autoFeeding",
>"msgResp":"ok",
>"autoFeeding":
>{
>"timerOneFlag":true,
>"timerOneHour":1,
>"timerOneMinutes":30,
>"timerTwoFlag":true,
>"timerTwoHour":1,
>"timerTwoMinutes":30,
>"timeFlag":true,
>"timeStamp":1586162503,
>"incrementFlag":true,
>"increment":100,
>"intervalTimeFlag":true,
>"iTimerDay":1,
>"iTimerHour":1,
>"iTimerMinutes":30,
>"intervalExitFlag":true,
>"timeModuleFlag":true,
>"timerOneModuleFlag":true,
>"timerTwoModuleFlag":true,
>"timeModuleExitFlag":true,
>"intervalMoudleFlag":true,
>}
>}
>```


# 13.1 加液量参数
* req:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"materialNum",
>"materialNum"
>{//界面默认以下参数
>"yg":0.56,
>"dti":64.0,
>"dto":73.0,
>"dci":112.0,
>"ht":2000.0,
>"hr":2100.0,
>"pco":15,
>"qg":1.0,
>"qw":1.0,
>"twh":22.0,
>"tr":120.0,
>"hour":0,
>"minutes":0,
>"oldPt":0,
>"oldPc":0,
>"ygjpd":2.5,
>}
>}
>```
* resp:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"materialNum",
>"msgResp":"ok",
>}
>```
# 13.2 加液量参数
* req:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"materialNum",
>}
>```
* resp:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"materialNum",
>"msgResp":"ok",
>"materialNum"
>{//界面默认以下参数
>"yg":0.56,
>"dti":64.0,
>"dto":73.0,
>"dci":112.0,
>"ht":2000.0,
>"hr":2100.0,
>"pco":15,
>"qg":1.0,
>"qw":1.0,
>"twh":22.0,
>"tr":120.0,
>"hour":0,
>"minutes":0,
>"oldPt":0,
>"oldPc":0,
>"ygjpd":2.5,
>}
>}
>```

# 14.１ 读取集控器系统日志
**websocket example:**
* req:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"eabiLog",
>"year":"2020",
>"month":"7",
>}
>```

* resp:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"eabiLog",
>"msgResp":"ok",
>"logData":"xxxxxxxxxxx",
>}
>```

# 14.2 删除集控器系统日志
**websocket example:**
* req:
>```json
>{
>"msgType":"DELETE",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"eabiLog",
>"year":"2020",
>"month":"7",
>}
>```

* resp:
>```json
>{
>"msgType":"DELETE",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"eabiLog",
>"msgResp":"ok",
>}
>```


# 15.１ 读取集控器算法日志
**websocket example:**
* req:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"feedingAlgo",
>}
>```

* resp:
>```json
>{
>"msgType":"GET",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"feedingAlgo",
>"msgResp":"ok",
>"logData":"xxxxxxxxxxx",
>}
>```

# 16 系统重启
**websocket example:**
* req:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"reboot",
>"cause":"xxxxx"
>}
>```
* resp:
>```json
>{
>"msgType":"PUT",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"reboot",
>"msgResp":"ok",
>}
>```
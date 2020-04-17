
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
>"msgParam":"getGatewayParam",
>"msgResp":"ok",
>"gwId":"AFAF73BADCF6", 
>"gwIP":"192.168.0.2",
>"serverIP":"10.56.1.34",
>"serverPort":"8287",
>"rfId":"AFAF73BADCED",
>"rfChannel":1,//TODO:范围？
>"rfNetId":"AFAF73BADCAA",
>"dataUpCycle":1, //单位秒
>"heartCycle":1,  //单位秒
>"dataReadCycle":1 //单位秒,gw读取所有的终端设备，一次轮训的时间
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
>"rfChannel":1,   //TODO:范围？
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

# 3 rfNetInfo
* 接口说明
>  射频层入网设备信息

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
>[{
>   "Id":"AFAF73BADCF6",
>   "softwareVer":"1.0.1",
>   "hardwareVer":"1.0.2",
>   "Channel_1":"yes",//yes,no  
>   "Channel_2":"yes",//yes,no
>   "Channel_3":"yes",//yes,no
>   "Channel_4":"yes",//yes,no
>   "Channel_5":"yes",//yes,no
>   "Channel_6":"yes",//yes,no
>   "Channel_7":"yes",//yes,no
>   "Channel_8":"yes",//yes,no
>   },
>   {
>   "Id":"AFAF73BADCF7",
>   "softwareVer":"1.0.1",
>   "hardwareVer":"1.0.2",
>   "Channel_1":"yes",//yes,no  
>   "Channel_2":"yes",//yes,no
>   "Channel_3":"yes",//yes,no
>   "Channel_4":"yes",//yes,no
>   "Channel_5":"yes",//yes,no
>   "Channel_6":"yes",//yes,no
>   "Channel_7":"yes",//yes,no
>   "Channel_8":"yes",//yes,no
>   }
> ]
>}
>```

# 4 abnormalAlarm
* 接口说明
>  业务报警（支持n个终端，每个终端n个通道 n >= 8）

**websocket example:**
* req:
>```json
>{
>"msgType":"POST",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162503,
>"msgParam":"abnormalAlarm",
>[
>    {
>        "alarmId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>        "timeStamp"：1586162503，
>        "sensorId"："AFAF73BADCF7",
>        "channelNum"：1，
>        "alarmMsg":"电压过高，超过阀值0.8V"
>    },
>    {
>        "alarmId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>        "timeStamp"：1586162503，
>        "sensorId"："AFAF73BADCF7",
>        "channelNum"：1，
>        "alarmMsg":"电压过高，超过阀值0.8V"
>    }
>]
>}
>```
* resp:
>```json
>{
>"msgType":"POST",
>"msgId":"a7356eac-71ae-4862-b66c-a212cd292baf",
>"msgGwId":"AFAF73BADCF6",
>"msgTimeStamp":1586162656,
>"msgParam":"abnormalAlarm",
>"msgResp":"ok"
>}
>```

# 5 sensorData
* 接口说明
>  上传传感器数据，通过POST上传文件，采用csv格式的文件

* TODO:文件格式，上传那些数据。包括文件命名
* 网管id	设备id(仪表号)	时间戳(网管生产)	通道	通道类型	值


# 6 



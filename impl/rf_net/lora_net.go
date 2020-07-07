package rfNet

import (
	"eabi_gateway/module"
	"eabi_gateway/module/lora"
	myLog "eabi_gateway/module/my_log"
	"sync"
	"time"
)

var log module.LogInterfase
var mu sync.RWMutex

//var defttyName = "/dev/ttyUSB1"

var defttyName = "/dev/ttyS7"
var defBaud = 9600
var defReadTimeOut = time.Millisecond * 1
var receiveChan chan []byte

func Send(b []byte) {
	mu.Lock()
	mu.Unlock()

	for {
		if _, err := lora.Write(b); err != nil {
			log.PrintlnErr(err)
			rebootOpenTTY()
			continue
		}
		break
	}

}

func Read(buf []byte) int {
	b := <-receiveChan
	copy(buf, b)
	return len(b)
}

func waitReceive() {
	for {
		buf := make([]byte, 1024)
		mu.Lock()
		mu.Unlock()
		if n, err := lora.Read(buf); err != nil {
			log.PrintlnErr(err)
			time.Sleep(time.Second * 10)
		} else {
			if n == 0 {
				continue
			}
			receiveChan <- buf[0:n]
		}
	}
}

func rebootOpenTTY() {
	var one sync.Once
	lora.Close()

	mu.Lock()
	defer mu.Unlock()
	t := time.Now().Unix()

	for {
		if err := lora.Open(defttyName, defBaud, defReadTimeOut); err != nil {
			one.Do(func() {
				log.PrintlnErr(err)
			})
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}

	log.PrintfInfo("Reconnect open lora module %d Second after ", time.Now().Unix()-t)
}

//LoraInit 初始化lora网络模块，维护
func LoraInit() {
	log = &myLog.L{}
	receiveChan = make(chan []byte)

	var one sync.Once
	t := time.Now().Unix()
	for {
		if err := lora.Open(defttyName, defBaud, defReadTimeOut); err != nil {
			one.Do(func() {
				log.PrintlnErr(err)
			})
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}
	log.PrintfInfo("Reconnect open lora module %d Second after ", time.Now().Unix()-t)
	go waitReceive()
}

package busNet

//工业485总线

import (
	"eabi_gateway/module"
	myLog "eabi_gateway/module/my_log"
	rs "eabi_gateway/module/rs_485"
	"sync"
	"time"
)

var log module.LogInterfase
var mu sync.RWMutex

var defttyName = "/dev/ttyS4"
var defBaud = 9600
var defReadTimeOut = time.Millisecond * 1
var receiveChan chan []byte

func Send(b []byte) {
	mu.Lock()
	mu.Unlock()

	for {
		if _, err := rs.Write(b); err != nil {
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
		if n, err := rs.Read(buf); err != nil {
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
	rs.Close()

	mu.Lock()
	defer mu.Unlock()
	t := time.Now().Unix()

	for {
		if err := rs.Open(defttyName, defBaud, defReadTimeOut); err != nil {
			one.Do(func() {
				log.PrintlnErr(err)
			})
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}

	log.PrintfInfo("Reconnect open bus module %d Second after ", time.Now().Unix()-t)
}

func Init() {
	log = &myLog.L{}
	receiveChan = make(chan []byte)

	var one sync.Once
	t := time.Now().Unix()
	for {
		if err := rs.Open(defttyName, defBaud, defReadTimeOut); err != nil {
			one.Do(func() {
				log.PrintlnErr(err)
			})
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}
	log.PrintfInfo("Reconnect open bus module %d Second after ", time.Now().Unix()-t)
	go waitReceive()
}

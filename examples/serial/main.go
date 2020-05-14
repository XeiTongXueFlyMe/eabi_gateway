package main

import (
	"log"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{
		Name:        "/dev/ttyUSB0",
		Baud:        115200,
		ReadTimeout: 0,
	}

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	n, err := s.Write([]byte("012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"))
	if err != nil {
		log.Fatal(err)
	}

	for {
		buf := make([]byte, 1024)
		n, err = s.Read(buf)
		if n > 0 {
			log.Printf("%q", buf[:n])
		}
		log.Printf("a")
	}

}

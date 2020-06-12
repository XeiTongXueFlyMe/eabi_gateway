package main

import (
	"fmt"
	"os"
	"time"

	"github.com/imroc/req"
)

type Temp struct {
	Msg string `json:"msg"`
}

func main() {
	file, _ := os.Open("202-5.log")
	resp, err := req.Post("http://192.168.0.194:8085/Test/addFileUpload", req.FileUpload{
		File:      file,
		FieldName: "file",
		FileName:  "日志2015.log",
	})
	if err != nil {
		fmt.Println(err)
	}
	temp := &Temp{}
	resp.ToJSON(temp)

	fmt.Println(temp.Msg)

	for {
		time.Sleep(1 * time.Hour)
	}
}

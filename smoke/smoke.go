package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"time"
)

func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Printf("GPIO 针脚读取错误！")
	}

	pin := rpio.Pin(21)

	defer rpio.Close()

	pin.Input()
	//pin.PullDown()

	fmt.Println("开启烟雾检测...")
	for {
		state := pin.Read()
		switch state {
		case rpio.High:
			fmt.Println("检测到烟雾！" + string(state))
		case rpio.Low:
			fmt.Println("没有烟雾！" + string(state))
		}
		time.Sleep(time.Second * 1)
	}

}

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
	p18 := rpio.Pin(18)

	defer rpio.Close()

	pin.Input()
	p18.Output()
	fmt.Println("开启烟雾检测...")
	for {
		state := pin.Read()
		switch state {
		case rpio.High:
			p18.Low()
			fmt.Println("没有烟雾！" + string(state))
		case rpio.Low:
			p18.High()
			fmt.Println("检测到烟雾！" + string(state))
		}
		time.Sleep(time.Second * 1)
	}

}

package main

import (
    "fmt"
    "github.com/stianeikeland/go-rpio"
)

func main() {
    err := rpio.Open()
    if err != nil {
        fmt.Printf("GPIO 针脚读取错误！")
    }

    p6 := rpio.Pin(6)
    p19 := rpio.Pin(19)
    p13 := rpio.Pin(13)

    defer rpio.Close()

    p6.Output()
    p13.Output()
    p19.Output()

    p6.Low();
    p13.Low();
    p19.Low();
}
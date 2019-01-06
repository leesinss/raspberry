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

    ds := rpio.Pin(6)
    shcp := rpio.Pin(19)
    stcp := rpio.Pin(13)

    defer rpio.Close()

    ds.Output()
    shcp.Output()
    stcp.Output()

    ds.Low();
    shcp.Low();
    stcp.Low();
    for {
        for j := 0; j < 8; j++ {
            for i := 0; i < 8; i++ {
                if j == i {
                    ds.High()
                } else {
                    ds.Low()
                }
                shcp.Low()
                shcp.High()

            }
            stcp.Low()
            stcp.High()
            if (j == 7) {
                j = 0
            }
        }
    }
}

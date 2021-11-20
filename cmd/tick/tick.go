package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("tick prototype")
	for i := 0; i < 3; i++ {
		fmt.Println("sleep 100ms")
		time.Sleep(100)
	}
}

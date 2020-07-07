package main

import (
	"fmt"
	"time"
)

func Run1() {
	fmt.Println("\n	1.")
	for i := 0; i < 3; i++ {
		fmt.Println("time now: ", time.Now().UnixNano()/1e6)
		if i == 3 {
			fmt.Println("kết thúc")
			break
		}
		time.Sleep(3 * time.Second)
	}
}

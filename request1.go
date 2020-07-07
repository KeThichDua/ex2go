package main

import (
	"fmt"
	"time"
)

func Run1() {
	// 1. Tạo 1 chương trình cứ 3s in ra dòng chữ: `time now: {milliseconds}`:
	// Sau khi in được 3 lần thì in ra dòng chữ `kết thúc`
	fmt.Println("\n	1.")
	for i := 0; i < 3; i++ {
		fmt.Println("time now: ", time.Now().UnixNano()/1e6)
		if i == 2 {
			fmt.Println("kết thúc")
			break
		}
		time.Sleep(3 * time.Second)
	}
}

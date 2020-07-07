package main

import (
	"fmt"
	"time"
)

func Run8() {
	// 8. Tạo 1 interval time sao cho cứ 100ms in ra dùng chữ `${time.Now().Unix()} done`
	fmt.Println("\n	8.")
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(time.Now().Unix(), " done")
	}
}

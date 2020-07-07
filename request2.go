package main

import (
	"fmt"
	"time"
)

func Run2() {
	// 2. Viết 1 đoạn chương trình tính ra ngày hiện tại theo timestamp.
	// Điểm mốc từ mức 0  tại 1/1/1970
	fmt.Println("\n	2.")
	fmt.Println("Số ngày hiện tại: ", time.Now().Unix()/3600/24)
}

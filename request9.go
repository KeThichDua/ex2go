package main

import (
	"fmt"
	"time"
)

func Run9() {
	//9. Tạo 1 đoạn code sử dụng `time.AfterFunc()`, sau 100ms thì in ra dòng chữ `i'm study`
	fmt.Println("\n	9.")
	// tạo thời hạn là 100ms
	DurationOfTime := time.Duration(100) * time.Millisecond

	//định nghĩa giá trị của AfterFunc()
	f := func() {
		// in ra `i'm study`khi dc gọi bởi AfterFunc() trong thời gian bắt đầu
		fmt.Println("i'm study")
	}

	// gọi AFterFunc() cùng f
	Timer := time.AfterFunc(DurationOfTime, f)
	defer Timer.Stop()
	time.Sleep(1 * time.Second)
}

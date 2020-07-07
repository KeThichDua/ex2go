package main

import (
	"context"
	"fmt"
	"time"
)

func Run9() {
	//9. Tạo 1 đoạn code sử dụng `time.AfterFunc()`, sau 100ms thì in ra dòng chữ `i'm study`
	fmt.Println("\n	9.")

	ctx := context.Background()
	err := TimeOut4(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
	}
}

func TimeOut4(ctx context.Context) error {
	// tạo thời hạn là 100ms
	DurationOfTime := time.Duration(100) * time.Millisecond

	//định nghĩa giá trị của AfterFunc()
	f := func() {
		// in ra `i'm study`khi dc gọi bởi AfterFunc() trong thời gian bắt đầu
		fmt.Println("i'm study")
	}
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Time out")
			cancel()
			return ctx.Err()
		default:
			// gọi AFterFunc() cùng f
			Timer1 := time.AfterFunc(DurationOfTime, f)
			defer Timer1.Stop()
			time.Sleep(100 * time.Millisecond)
		}
	}
}

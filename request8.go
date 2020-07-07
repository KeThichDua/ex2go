package main

import (
	"context"
	"fmt"
	"time"
)

func Run8() {
	// 8. Tạo 1 interval time sao cho cứ 100ms in ra dùng chữ `${time.Now().Unix()} done`
	fmt.Println("\n	8.")
	ctx := context.Background()
	err := TimeOut3(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
	}
}

func TimeOut3(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("TIME OUT")
			cancel()
			return ctx.Err()
		default:
			time.Sleep(100 * time.Millisecond)
			fmt.Println(time.Now().Unix(), " done")
		}
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func Run3() {
	fmt.Println("\n	3.")
	ctx := context.Background()
	err := TimeOut(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
	}
}

func TimeOut(ctx context.Context) error {
	// tạo context để tạo time out và cancel để hủy khi time out
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	// gọi cancel khi kết thúc hàm
	defer cancel()
	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done(): // kiểm tra xem time out chưa
			fmt.Println("TIME OUT")
			cancel()         // gọi cancel để hủy hàm
			return ctx.Err() // trả ra lỗi khi chạy hàm
		default:
			time.Sleep(3 * time.Second)
			log.Println("Sleep: ", i+1)
		}
	}

	fmt.Println("ALL DONE")
	return nil
}

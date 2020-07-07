package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func Run7() {
	fmt.Println("\n	7.")
	ns := time.Now().UnixNano()
	ctx := context.Background()
	err := TimeOut2(ctx, ns)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
	}
}

func TimeOut2(ctx context.Context, ns int64) error {
	// tạo context để tạo time out và cancel để hủy khi time out
	ctx, cancel := context.WithTimeout(ctx, 7*time.Second)
	// gọi cancel khi kết thúc hàm
	defer cancel()
	for {
		select {
		case <-ctx.Done(): // kiểm tra xem time out chưa
			fmt.Println("TIME OUT")
			cancel()         // gọi cancel để hủy hàm
			return ctx.Err() // trả ra lỗi khi chạy hàm
		default:
			time.Sleep(3 * time.Second)
			t := time.Now().UnixNano()
			log.Println("Hiệu: ", t-ns)
		}
	}
}

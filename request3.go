package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func Run3() {
	// 3. Thực hiện 1 chương trình với 1 vòng lặp for và 3 lần sleep mỗi lần sleep 3sec
	// Nhưng sau 3s thì kết thúc hàm đấy
	// Sử dụng và tìm hiểu context. Nêu được tác dụng của context trong chương trình.
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

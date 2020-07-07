package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func Run7() {
	// 	7. Tạo 1 đối tượng context với 1 value là timestamp hiện tại tính theo `ns`
	// chạy qua 1 hàm như sau. Sau 3s in ra hiệu của thời gian của thời gian hiện
	// tại với biến dữ liệu trong context. in ra màn hình kết quả.
	// ```go
	// func x(ctx context.Context) {

	// }
	// ```
	// gợi ý sử dụng một method của time để tính tạo ra 1 tính năng như setTimeout trong js
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

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

	f := func(ctx context.Context, k int64) {
		if v := ctx.Value(k); v != nil {
			log.Println("found value:", v)
			time.Sleep(3 * time.Second)
			log.Println("hiệu:", time.Now().UnixNano()-k)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := time.Now().UnixNano()
	ctx := context.WithValue(context.Background(), k, "k")

	f(ctx, k)
}

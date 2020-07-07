package main

import (
	"fmt"
	"time"
)

func Run6() {
	//6. Trong golang mặc định thì thời gian dạng số được sử dụng với các loại mốc đơn vị nào?
	fmt.Println("\n	6.")
	t := time.Now()
	fmt.Println("Năm, tháng, ngày, giờ, phút, giây, múi giờ")
	fmt.Println(t)
}

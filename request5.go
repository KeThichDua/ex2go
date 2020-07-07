package main

import (
	"fmt"
	"time"
)

func CheckDay(i int) {
	switch i {
	case 1:
		fmt.Println("Thứ hai")
	case 2:
		fmt.Println("Thứ ba")
	case 3:
		fmt.Println("Thứ tư")
	case 4:
		fmt.Println("Thứ năm")
	case 5:
		fmt.Println("Thứ sáu")
	case 6:
		fmt.Println("Thứ bảy")
	case 7:
		fmt.Println("Chủ nhật")
	default:
		fmt.Println("Lỗi")
	}
}

func Run5() {
	//5. Tính ra số ngày trong tuần(dạng string và number) của mốc thời gian sau `1592190385`
	fmt.Println("\n	5.")
	fmt.Println("Số ngày: ", 1592190385/3600/24)
	fmt.Println("Ngày này là: ", time.Unix(1592190385, 0).Weekday())
}

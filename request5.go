package main

import (
	"fmt"
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
	fmt.Println("\n	5.")
	i := 1592190385 + 25200 // 25200 để chuyển múi giờ VN
	fmt.Println(i)

	index := 4
	for {
		i = i - 86400
		index += 1
		if index > 7 {
			index = 1
		}
		if i < 86400 {
			break
		}
	}
	fmt.Print("Ngay thu ", index, " trong tuần là: ")
	CheckDay(index)
}

package main

import (
	"fmt"
	"time"
)

func Run5() {
	//5. Tính ra số ngày trong tuần(dạng string và number) của mốc thời gian sau `1592190385`
	fmt.Println("\n	5.")
	fmt.Println("Số ngày: ", 1592190385/3600/24)
	fmt.Println("Ngày này là: ", time.Unix(1592190385, 0).Weekday())
}

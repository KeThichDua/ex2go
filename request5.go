package main

import (
	"fmt"
	"time"
)

func Run5() {
	//5. Tính ra số ngày trong tuần(dạng string và number) của mốc thời gian sau `1592190385`
	fmt.Println("\n	5.")
	weekday := time.Unix(1592190385, 0).Weekday()
	fmt.Println("Số ngày trong tuần là: ", int(weekday))
	fmt.Println("Ngày này là: ", weekday)
}

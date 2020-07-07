package main

import (
	"fmt"
	"time"
)

func Run2() {
	fmt.Println("\n	2.")
	now := time.Now()
	i := now.Unix()
	fmt.Println(now, "	", i)

	y := 1970
	for {
		if y%100 == 0 {
			if y%4 == 0 {
				i = i - 31622400 // số seconds năm nhuận
			} else {
				i = i - 31536000 // số seconds năm ko nhuận
			}
		} else {
			if y%4 == 0 {
				i = i - 31622400
			} else {
				i = i - 31536000
			}
		}
		y += 1
		if i < 31536000 {
			break
		}
	}
	fmt.Println("year: ", y)

	m := 0
	for {
		if m == 0 || m == 2 || m == 4 || m == 6 || m == 7 || m == 9 || m == 11 || m == 12 {
			i = i - 2678400 // số seconds 31 ngày
		} else if m == 3 || m == 5 || m == 8 || m == 10 {
			i = i - 2592000 // số seconds 30 ngày
		} else {
			if y%100 == 0 {
				if y%4 == 0 {
					i = i - 2505600 // số seconds 29 ngày
				} else {
					i = i - 2419200 // số seconds 28 ngày
				}
			} else {
				if y%4 == 0 {
					i = i - 2505600
				} else {
					i = i - 2419200
				}
			}
		}
		m += 1
		if i < 2419200 {
			m += 1
			break
		}
	}
	fmt.Println("month: ", m)

	d := 0
	for {
		i = i - 86400 // số seconds 1 ngày
		d += 1
		if i < 86400 {
			d += 1
			break
		}
	}
	fmt.Println("day: ", d)

	h := 0
	for {
		i = i - 3600
		h += 1
		if i < 3600 {
			// đổi h vùng VN
			h += 7
			break
		}
	}
	fmt.Println("hour: ", h)

	mi := 0
	for {
		i = i - 60
		mi += 1
		if i < 60 {
			break
		}
	}
	fmt.Println("minute: ", mi)
	fmt.Println("second: ", i)

}

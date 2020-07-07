package main

import (
	"fmt"
)

func Run4() {
	fmt.Println("\n	4.")
	i := 1592190294764144364
	fmt.Println(i)

	y := 1970
	for {
		if y%100 == 0 {
			if y%4 == 0 {
				i = i - 31622400*1e9
			} else {
				i = i - 31536000*1e9
			}
		} else {
			if y%4 == 0 {
				i = i - 31622400*1e9
			} else {
				i = i - 31536000*1e9
			}
		}
		y += 1
		if i < 31536000*1e9 {
			break
		}
	}
	fmt.Println("year: ", y)

	m := 0
	for {
		if m == 0 || m == 2 || m == 4 || m == 6 || m == 7 || m == 9 || m == 11 || m == 12 {
			i = i - 2678400*1e9
		} else if m == 3 || m == 5 || m == 8 || m == 10 {
			i = i - 2592000*1e9
		} else {
			if y%100 == 0 {
				if y%4 == 0 {
					i = i - 2505600*1e9
				} else {
					i = i - 2419200*1e9
				}
			} else {
				if y%4 == 0 {
					i = i - 2505600*1e9
				} else {
					i = i - 2419200*1e9
				}
			}
		}
		m += 1
		if i < 2419200*1e9 {
			m += 1
			break
		}
	}
	fmt.Println("month: ", m)

	d := 0
	for {
		i = i - 86400*1e9
		d += 1
		if i < 86400*1e9 {
			d += 1
			break
		}
	}
	fmt.Println("day: ", d)

	h := 0
	for {
		i = i - 3600*1e9
		h += 1
		if i < 3600*1e9 {
			// đổi h vùng VN
			h += 7
			break
		}
	}
	fmt.Println("hour: ", h)

	mi := 0
	for {
		i = i - 60*1e9
		mi += 1
		if i < 60*1e9 {
			break
		}
	}
	fmt.Println("minute: ", mi)
}

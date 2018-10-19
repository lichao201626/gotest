package myatoi

import (
	"log"
	"strconv"
	"strings"
)

func Myatoi(intd string) int {
	var res int
	var flag int = 1
	/* 	var intmax = int(^uint(0) >> 33)
	   	var intmin = ^intmax
	*/
	var intmax = 1<<31 - 1
	var intmin = -1 << 31
	intd = strings.TrimLeft(intd, " ")
	if len(intd) < 1 {
		return 0
	}
	if string(intd[0]) == "-" {
		flag = -1
	}
	for k, s := range intd {
		log.Println(string(s))
		if string(s) >= "0" && string(s) <= "9" {
			temp, _ := strconv.Atoi(string(s))
			res = res*10 + temp
			//到2147483648
			if res > intmax+1 {
				break
			}
		} else {
			// flag = -1
			if k != 0 {
				break
			}
			if k == 0 && string(s) != "-" && string(s) != "+" {
				break
			}
		}
	}
	if flag != 1 {
		res = res * -1
	}
	if res > intmax {
		return intmax
	}
	if res < intmin {
		return intmin
	}
	return res
}

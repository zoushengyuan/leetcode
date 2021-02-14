package main

import (
	"fmt"
)

func myAtoi(s string) int {
	var pos int
	var d byte
	bs := []byte(s)
	for pos, d = range bs {
		if d != ' ' {
			break
		}
	}
	if pos >= len(s) {
		return 0
	}
	sign := 1
	if s[pos] == '-' {
		sign = -1
		pos++
	} else if s[pos] == '+' {
		pos++
	}

	num := 0
	max := 1<<31 - 1
	for _, d := range bs[pos:] {
		if !('0' <= d && d <= '9') {
			break
		}
		num = num*10 + int(d-'0')
		if num > max+(1-sign)/2 { //
			num = max + (1-sign)/2
			break
		}
	}

	return num * sign
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("-42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("91283472332"))
}

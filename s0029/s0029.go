package main

import (
	"fmt"
	"math"
)

func divide(a int, b int) int {
	if a == math.MinInt32 {
		if b == -1 {
			return math.MaxInt32
		}
		if b == 1 {
			return a
		}
	}
	sign := 1
	if a > 0 {
		a = -a
		sign = -sign
	}
	if b > 0 {
		b = -b
		sign = -sign
	}

	return sign * div(a, b)
}

func div(a int, b int) int {
	if a > b {
		return 0
	}
	q := 1
	tb := b
	for (tb + tb) > a {
		q += q
		tb += tb
	}
	return q + div(a-tb, b)
}

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
	fmt.Println(divide(0, 3))
	fmt.Println(divide(0, -3))
}

package main

import (
	"fmt"
)

var lens [10]byte = [10]byte{0, 0, 3, 3, 3, 3, 3, 4, 3, 4}
var chars [][]byte = [][]byte{
	{},                   //0
	{},                   //1
	{'a', 'b', 'c'},      //2
	{'d', 'e', 'f'},      //3
	{'g', 'h', 'i'},      //4
	{'j', 'k', 'l'},      //5
	{'m', 'n', 'o'},      //6
	{'p', 'q', 'r', 's'}, //7
	{'t', 'u', 'v'},      //8
	{'w', 'x', 'y', 'z'}, //9
}

func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return nil
	}
	ds := make([]byte, n, n)
	units := make([]byte, n, n)
	total := 1
	for i, d := range []byte(digits) {
		ds[i] = d - '0'
		units[i] = lens[ds[i]]
		total *= int(units[i])
	}
	x := make([]byte, n, n)
	//fmt.Println(x)
	r := []string{}
	//fmt.Println(toString(ds, x))
	for i := 0; i < total; i++ {
		r = append(r, toString(ds, x))
		x = addOne(units, x)
	}
	return r
}

func addOne(units []byte, x []byte) []byte {
	n := len(units)
	var carry byte = 0
	x[0]++
	for i := 0; i < n; i++ {
		x[i] += carry
		carry = 0
		if x[i] == units[i] {
			x[i] = 0
			carry = 1
		} else {
			break
		}
	}
	return x
}

func toString(ds []byte, x []byte) string {
	n := len(x)
	y := make([]byte, n, n)
	for i := 0; i < n; i++ {
		y[i] = chars[ds[i]][x[i]]
	}
	return string(y)
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("29"))
}

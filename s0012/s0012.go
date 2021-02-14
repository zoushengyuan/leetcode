package main

import (
	"fmt"
)

var thousands []string = []string{"M", "MM", "MMM"}
var hundreds []string = []string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var tens []string = []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var digits []string = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

func intToRoman(num int) string {
	th := num / 1000
	num -= th * 1000
	h := num / 100
	num -= h * 100
	t := num / 10
	d := num - t*10
	s := ""
	if th != 0 {
		s += thousands[th-1]
	}
	if h != 0 {
		s += hundreds[h-1]
	}
	if t != 0 {
		s += tens[t-1]
	}
	if d != 0 {
		s += digits[d-1]
	}
	return s

}

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}

package main

import (
	"fmt"
)

/*
func calc(bs []byte) (left int, right int) {
	for _, b := range bs {
		if b == '(' {
			left++
		} else {
			right++
		}
	}
	return left, right
}

func toString(rs [][]byte) []string {
	s := []string{}
	for _, r := range rs {
		s = append(s, string(r))
	}
	return s
}

func generateParenthesis(n int) []string {
	rs := [][]byte{{'('}}
	for i := 0; i < n*2-1; i++ {
		size := len(rs)
		r := [][]byte{}
		for j := 0; j < size; j++ {
			left, right := calc(rs[j])
			if left-right < n && n-left > 0 {
				app := make([]byte, len(rs[j]))
				copy(app, rs[j])
				r = append(r, append(app, '('))
			}
			fmt.Println(i, j, "llll", rs[j], left, right, toString(r))
			if left-right > 0 && n-right > 0 {
				app := make([]byte, len(rs[j]))
				copy(app, rs[j])
				r = append(r, append(app, ')'))
			}
			fmt.Println(i, j, "rrrr", rs[j], left, right, toString(r))
		}
		rs = r
		//r = [][]byte{}
	}
	return toString(rs)

}

func correct(i int, length int) (string, bool) {
	left := 0
	right := 0
	bs := make([]byte, length)
	half := length / 2
	for j := length - 1; j >= 0; j-- {
		if i&1 == 0 {
			left++
			bs[j] = '('
		} else {
			right++
			bs[j] = ')'
		}
		if left > right || right > half {
			return "", false
		}
		i = i >> 1
	}
	return string(bs), true
}

func generateParenthesis(n int) []string {
	size := 1 << (2*n - 1)
	ps := []string{}
	length := 2 * n
	for i := 0; i < size; i++ {
		if p, ok := correct(i, length); ok {
			ps = append(ps, p)
		}
	}
	return ps
}
*/
func toString(rs [][]byte) []string {
	s := []string{}
	for _, r := range rs {
		s = append(s, string(r))
	}
	return s
}

func gp(pss [][]byte) [][]byte {

	return pss
}

func generateParenthesis(n int) []string {
	pss := [][]byte{}
	return toString(pss)
}

func main() {
	//fmt.Println(correct(5, 4))
	fmt.Println(generateParenthesis(3))
	//fmt.Println(generateParenthesis(2))
	//fmt.Println(calc([]byte("(()")))
	//fmt.Println(calc([]byte("((())")))
}

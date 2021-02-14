package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	//fmt.Println(s, p)
	pn := len(p)
	sn := len(s)
	if pn == 0 {
		return sn == 0
	}
	if p[0] == '.' {
		if pn > 2 && p[1] == '*' {
			return isMatch(s, p[2:]) || sn > 0 && isMatch(s[1:], p)
		}
		return sn > 0 && isMatch(s[1:], p[1:])
	}
	if pn > 1 && p[1] == '*' {
		if sn == 0 {
			return isMatch(s, p[2:])
		}
		if s[0] == p[0] {
			return isMatch(s[1:], p) || isMatch(s, p[2:])
		}
		return isMatch(s, p[2:])
	}
	if sn == 0 {
		return false
	}
	return p[0] == s[0] && sn > 0 && isMatch(s[1:], p[1:])
}

func main() {
	var s, p string

	s = "aa"
	p = "a"
	fmt.Println(isMatch(s, p))
	s = "aa"
	p = "a*"
	fmt.Println(isMatch(s, p))

	s = "ab"
	p = ".*"
	fmt.Println(isMatch(s, p))
	s = "b"
	p = "a*b"
	fmt.Println(isMatch(s, p))
	s = "aab"
	p = "c*a*b"
	fmt.Println(isMatch(s, p))
	s = "mississippi"
	p = "mis*is*p*."
	fmt.Println(isMatch(s, p))
	s = "a"
	p = "ab*a"
	fmt.Println(isMatch(s, p))
	s = "ab"
	p = ".*c"
	fmt.Println(isMatch(s, p))
	s = "aa"
	p = "a*"
	fmt.Println(isMatch(s, p))
	s = "aa"
	p = "a*c"
	fmt.Println(isMatch(s, p))
	s = "a"
	p = ".*..a*"
	fmt.Println(isMatch(s, p))
	//*/
	s = ""
	p = "c*c*"
	fmt.Println(isMatch(s, p))
}

package main

import "fmt"

type stack struct {
	m []int
	i int
}

func create(n int) *stack {
	st := &stack{}
	st.m = make([]int, n)
	return st
}

func (st *stack) push(x int) {
	st.m[st.i] = x
	st.i++
}

func (st *stack) pop() (int, bool) {
	if st.i > 0 {
		st.i--
		return st.m[st.i], true
	}
	return -1, false
}

func (st *stack) all() []int {
	return st.m[:st.i]
}

func match(s []byte) {
	n := len(s)
	st := create(n)

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			st.push(i)
			continue
		}
		if j, ok := st.pop(); ok {
			s[i] = 0
			s[j] = 0
		}
	}
}

func longestMatch(s []byte) int {
	longest := 0
	sum := 0
	for _, c := range s {
		if c == 0 {
			sum++
		} else {
			if sum > longest {
				longest = sum
			}
			sum = 0
		}
	}
	if sum > longest {
		longest = sum
	}
	return longest
}

func longestValidParentheses(s string) int {
	bs := []byte(s)
	match(bs)
	return longestMatch(bs)
}

func main() {
	var s string

	s = "(()"
	fmt.Println(longestValidParentheses(s))

	s = ")()())"
	fmt.Println(longestValidParentheses(s))

	s = ""
	fmt.Println(longestValidParentheses(s))

	s = "()(()"
	fmt.Println(longestValidParentheses(s))
}

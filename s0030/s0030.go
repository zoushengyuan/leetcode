package main

import (
	"fmt"
)

func matchOne(word string, s []byte) int {
	n := len(word)
	if n > len(s) {
		return 0
	}
	for i := 0; i < n; i++ {
		if word[i] != s[i] {
			return 0
		}
	}
	return n
}

func matchAny(words []string, s []byte) int {
	wn := len(words)
	for i := 0; i < wn; i++ {
		if r := matchOne(words[i], s); r != 0 {
			if i != 0 {
				words[i], words[0] = words[0], words[i]
			}
			return r
		}
	}
	return 0
}

func matchAll(words []string, s []byte) bool {
	wn := len(words)
	i := 0
	j := 0

	for j < wn {
		if r := matchAny(words[j:], s[i:]); r != 0 {
			j++
			i += r
		} else {
			return false
		}
	}
	return true
}

func findSubstring(s string, words []string) []int {
	bs := []byte(s)
	n := len(bs)
	ps := []int{}
	for i := 0; i < n; i++ {
		if matchAll(words, bs[i:]) {
			ps = append(ps, i)
		}
	}
	return ps
}

func main() {
	words := []string{"foo", "bar"}
	s := "barfoothefoobarman"
	r := findSubstring(s, words)
	fmt.Println(r)

	words = []string{"word", "good", "best", "word"}
	s = "wordgoodgoodgoodbestword"
	r = findSubstring(s, words)
	fmt.Println(r)
}

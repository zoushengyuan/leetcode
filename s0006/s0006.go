package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	numStr := len(s)
	cycle := numRows + numRows - 2
	numCols := numStr / cycle * (numRows - 1)
	r := numStr % cycle
	if r > 0 {
		numCols++
		r -= numRows
	}
	if r > 0 {
		numCols += r
	}
	m := make([][]byte, numRows, numRows)
	for i := 0; i < numRows; i++ {
		m[i] = make([]byte, numCols, numCols)
	}
	for y, yc, pos := 0, 0, 0; y < numCols && pos < numStr; y, yc = y+1, yc+1 {
		if yc == numRows-1 {
			yc = 0
		}
		for x := 0; x < numRows && pos < numStr; x++ {
			if yc != 0 && yc+x != numRows-1 {
				continue
			}
			m[x][y] = s[pos]
			pos++
		}
	}
	bs := make([]byte, numStr, numStr)
	for x, pos := 0, 0; x < numRows && pos < numStr; x++ {
		for y := 0; y < numCols && pos < numStr; y++ {
			if m[x][y] == 0 {
				continue
			}
			bs[pos] = m[x][y]
			pos++
		}
	}

	for x := 0; x < numRows; x++ {
		for y := 0; y < numCols; y++ {
			if m[x][y] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c", m[x][y])
			}
		}
		fmt.Println()
	}

	return string(bs)
}

func main() {
	var s, r string
	//s = "Apalindromeisaword,phrase,number,orothersequenceofunitsthatcanbereadthesamewayineitherdirection,withgeneralallowancesforadjustmentstopunctuationandworddividers."
	s = ""
	for i := 0; i < 16; i++ {
		s += "0123456789"
	}
	r = convert(s, 10)
	fmt.Printf("%q\n", s)
	fmt.Printf("%q\n", r)
	fmt.Println(len(s), len(r))
	//"A,tsaclmapdpohttsmetaltennarhreiheerilosnodlorornahwioawutiw.iwa,suttadnrajstosnasrefcdyr,endtarrdseeqoaaiewncaouderi,buenenhieerptddoenmecbrettgsouciimuneihfnv"
	//"A,tsaclmapdpohttsmetaltennarhreiheerilosnodlorornahwioawutiwiwa,suttadnrajstonasrefcdyr,endtardseeqoaaiewncaoudri,buenenhieerptdoenmecbrettgsoucimuneihfnv\u0000\u0000\u0000\u0000\u0000\u0000"	s = "PAYPALISHIRING"

	/*
		r = convert(s, 3)
		fmt.Printf("%q\n", s)
		fmt.Printf("%q\n", r)
		fmt.Println(len(s), len(r))
	*/
}

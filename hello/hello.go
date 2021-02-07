package main

import (
	"fmt"
)

func main() {
	a := make([][]byte, 10, 10)
	a[0] = make([]byte, 3, 3)
	a[0][0] = 'a'
	a[0][1] = 'b'
	a[0][2] = 'c'
	a[1] = make([]byte, 3, 3)
	a[1][0] = 'a'
	a[1][1] = 'b'
	a[1][2] = 'e'
	b := make([]string, 10, 10)
	for i := 0; i < 10; i++ {
		b[i] = string(a[i])
	}
	fmt.Println(b)
}
//https://vscode.cdn.azure.cn/stable/f06011ac164ae4dc8e753a3fe7f9549844d15e35/code_1.37.1-1565886362_amd64.deb
//https:///vscode.cdn.azure.cn/stable/8490d3dde47c57ba65ec40dd192d014fd2113496/VSCode-darwin.zip

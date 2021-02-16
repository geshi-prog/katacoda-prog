package main

import "fmt"

func main() {
	a := 256 // 基数10
	b := 0b100000000 // 基数2
	c := 0o400 // 基数8
	d := 0x100 // 基数16
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
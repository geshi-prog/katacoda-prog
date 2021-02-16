package main

import "fmt"

func main() {
	a := 'a' // 単一のUnicode文字
	b := '\141' // 8ビットの8進数
	c := '\x61' // 8ビットの16進数
	d := '\u0061' // 16ビットの16進数
	e := '\U00000061' // 32ビットのUnicode番号
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
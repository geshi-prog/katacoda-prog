package main

import "fmt"

func main() {
	a := `message\n` // raw文字列リテラル
	b := "message\n" // interpreted文字列リテラル
	c := `\u3042\u3044\u3046` // raw文字列リテラル
	d := "\u3042\u3044\u3046" // interpreted文字列リテラル
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
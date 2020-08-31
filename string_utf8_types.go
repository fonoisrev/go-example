package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "abc世界"
	// 输出 61 62 63 e4 b8 96 e7 95 8c
	fmt.Printf("% x\n", str)
	// 字符串底层是byte数组，所以这里len(str)是9
	fmt.Println(len(str))

	// rune 的作用
	r := []rune(str)
	// 输出 [61 62 63 4e16 754c]
	fmt.Printf("%x\n", r)
	// 因为这里只有5个unicode字符，所以下面两个输出是等价的
	fmt.Println(len(r))
	fmt.Println(utf8.RuneCountInString(str))

	str2 := "\xe4\xb8\x96\xe7\x95\x8c"
	fmt.Println(str2)
	str2 = "\u4e16\u754c"
	fmt.Println(str2)

}

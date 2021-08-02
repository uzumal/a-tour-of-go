package main

import "fmt"

// 戻り値に名前をつける（短いコードの関数の際に有利）
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

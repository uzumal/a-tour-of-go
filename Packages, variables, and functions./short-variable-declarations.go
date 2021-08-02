package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// 関数の中ではvarの代わりに:=の代入文を用いて、暗黙的な型宣言ができる
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

package main

import "fmt"

var chk [100]bool

func main() {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	runes := []rune(str)
	for _, r := range runes {
		fmt.Println(r, string(r))
		chk[r] = true
	}
	fmt.Println(chk)

	a := rune("A")
}

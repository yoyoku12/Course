package main

import "fmt"

func main() {
	var a string = "ama"
	var b string = Reverse(a)
	if a == b {
		fmt.Println("палиндром")
	} else {
		fmt.Println("не палиндром")
	}
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

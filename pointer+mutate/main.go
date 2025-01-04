package main

import "fmt"

func main() {
	a := 5
	fmt.Println(a)
	double(&a)
	fmt.Println(a)
}

func double(num *int) {
	*num = *num * 2
}

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println("-----------До внесения изменений-------------")
	fmt.Printf("Slice: %v\n", a)
	fmt.Printf("Array: %v\n", c)
	changeSlice(a)
	changeArray(c)
	fmt.Println("-----------После внесения изменений-------------")
	fmt.Printf("Slice: %v\n", a)
	fmt.Printf("Array: %v\n", c)
}

func changeSlice(a []int) {
	a[0] = 0
}

func changeArray(b [5]int) {
	b[0] = 0
}

package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reverse(&a)
	fmt.Println(a)
}

func reverse(array *[10]int) {
	a := len(array) - 1
	reverseS := [len(array)]int{}
	for i := 0; i <= len(array)-1; i++ {
		reverseS[i] = array[a]
		a = a - 1
	}
	*array = reverseS
}

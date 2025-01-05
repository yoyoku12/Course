package main

import "fmt"

func main() {

	var p *int
	c := &p
	d := &c
	fmt.Println(p, c, d)

}

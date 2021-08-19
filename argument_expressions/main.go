package main

import "fmt"

func main() {

	var num = 454445.5465465
	another_var := 2515155
	fmt.Printf("%T\n", num)
	fmt.Printf("%T\n", another_var)

	var un_initlized_num uint64
	var empty_string string
	var bl bool

	fmt.Println(un_initlized_num, empty_string, bl)
}

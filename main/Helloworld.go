package main

import "fmt"
import "../day1"

func main() {
	//fmt.Println("Hello, Porco")
	//day1.Point()
	a, b := 1, 2
	day1.Trans(&a, &b)
	fmt.Println(a, b)
}

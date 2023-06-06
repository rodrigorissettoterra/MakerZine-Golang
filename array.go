package main

import "fmt"

func main() {
	var x [10]int
	x[4] = 80
	var y [5]int
	y[4] = 13
	fmt.Println("x -> ", x)
	fmt.Println("y -> ", y)
}

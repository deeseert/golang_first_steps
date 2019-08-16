package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := adder()
	for i := 0; i < 100; i++ {
		fmt.Println(sum(i))
	}
	fmt.Println("Hello World")
}

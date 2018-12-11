package main

import (
	"go_dev/day1/package_example/calc"
	"fmt"
)

func main() {
	sub := calc.Sub(2, 3)
	sum := calc.ClacSum(2, 3)
	fmt.Println("sub=", sub)
	fmt.Println("sum=", sum)
}

package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func add(num ...int) int {
	sum := 0
	for _, value := range num {
		sum = sum + value
	}
	return sum
}
func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
	res = add(10, 20, 30, 40)
	fmt.Println("sum of 10, 20, 30, 40 =", res)
}

package main

import "fmt"

//leetcode begin

func myPow(x float64, n int) float64 {
	if x == 0.0 {
		return 0.0
	}
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1.0 / x
	}
	if n%2 == 0 {
		return myPow(x*x, n/2)
	} else {
		return x * myPow(x*x, n/2)
	}
}

//leetcode end

func main() {
	fmt.Printf("Pow %v ^ %v ==> %v\n", 2.000, 10, myPow(2.000, 10))
	fmt.Printf("Pow %v ^ %v ==> %v\n", 2.100, 3, myPow(2.100, 3))
	fmt.Printf("Pow %v ^ %v ==> %v\n", 0, 0, myPow(0, 0))
	fmt.Printf("Pow %v ^ %v ==> %v\n", 100, 0, myPow(100, 0))
	fmt.Printf("Pow %v ^ %v ==> %v\n", 10, -2, myPow(10, -2))
}

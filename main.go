package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	a, b := 0, 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return b
}

func Greet(s string) string {
	return fmt.Sprintf("Hello, %s!", s)
}

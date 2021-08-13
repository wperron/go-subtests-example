package main

import "testing"

func TestFib(t *testing.T) {
	t.Run("Fibonacci valid input", valid)
	t.Run("Fibonacci invalid input", invalid)
}

func TestGreet(t *testing.T) {
	t.Run("Hello World", world)
}

func valid(t *testing.T) {
	if res := Fibonacci(8); res != 34 {
		t.Errorf("expect Fib(8) to be 21, got %d", res)
	}
}

func invalid(t *testing.T) {
	if res := Fibonacci(-8); res != 0 {
		t.Errorf("expect Fib(-8) to be 0, got %d", res)
	}
}

func world(t *testing.T) {
	if res := Greet("World"); res != "Hello World!" {
		t.Errorf("wrong output `%s`", res)
	}
}

package main

import "fmt"

// Fibonacci Number
// In mathematics, the Fibonacci numbers, commonly denoted Fn, 
// rm a sequence, called the Fibonacci sequence, such that each 
// number is the sum of the two preceding ones, starting from 0 and 1.

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func main() {
	getFibonacci := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(getFibonacci())
	}
}

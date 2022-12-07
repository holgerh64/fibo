// Anzahl der Aufrufe der Fibonacci-Rekursion
package main

import "fmt"

var g int64

// Fibonacci-Rekursion
func fib(a int64) int64 {
	if a == 1 || a == 0 {
		g++
		return 1
	} else {
		g = g + 2
		return (fib(a-2) + fib(a-1))
	}

}

func main() {
	var x int64

	x = 0
	g = 0
	fmt.Println("x\tg\tfib")
	for x = 0; x < 50; x++ {
		g = 0
		fmt.Println(x, "\t", g, "\t", fib(x))
	}
}

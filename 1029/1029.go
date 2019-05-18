package main

import (
	"fmt"
)

var call int = 0

func main() {
	var n, x int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		res := fib(x)
		fmt.Printf("fib(%d) = %d calls = %d\n", x, call-1, res)
		call = 0
	}
}

func fib(n int) int {
	call++
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

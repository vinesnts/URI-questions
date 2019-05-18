package main

import (
  "fmt"
)

func main() {
  var n, x, y, soma int
  fmt.Scan(&n)
  for i := 0; i < n; i++ {
    fmt.Scanf("%d %d", &x, &y)
    if x > y {
      x, y = y, x
    }
    for x++ ; x < y; x++ {
      if x % 2 != 0 {
        soma += x
      }
    }
    fmt.Printf("%d\n", soma)
    soma = 0
  }
}

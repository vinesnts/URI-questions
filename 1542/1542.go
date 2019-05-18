package main

import (
  "fmt"
)

func main() {
  var q, p, d, x int
  for {
    fmt.Scanf("%d %d %d", &q, &d, &p)
    if q == 0 && d != 0 && p != 0 {
      break
    }
    x = (q * d * p) / (p - q)
    if x == 1 {
      fmt.Printf("%d pagina\n", x)
    } else {
      fmt.Printf("%d paginas\n", x)
    }
  }
}

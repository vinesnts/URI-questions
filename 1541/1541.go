package main

import (
  "fmt"
  "math"
)

func main() {
    var a, b, p int
    var res float64
    for {
      fmt.Scanf("%d %d %d", &a, &b, &p)
      if a == 0 {
        break
      }
      area := a * b
      res = math.Sqrt((float64)((area * 100) / p))
      fmt.Printf("%d\n", int(res))
    }
}

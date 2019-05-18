package main

import (
  "fmt"
)

func main() {
  i := 0.0
  for i <= 2 {
    j := 1.0 + i
    for k := 0; k < 3; k++ {
      if i / 2 == 0 || i / 1 == 1 || i > 1.9 {
        fmt.Printf("I=%0.0f ", i)
        fmt.Printf("J=%0.0f\n", j)
      } else {
        fmt.Printf("I=%0.1f ", i)
        fmt.Printf("J=%0.1f\n", j)
      }
      j += 1
    }
    i += 0.2
  }
}

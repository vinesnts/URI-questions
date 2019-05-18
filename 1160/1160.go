package main

import (
  "fmt"
)

func main() {
  var n, PA, PB int
  var GA, GB float64
  fmt.Scan(&n)

  for i := 0; i < n; i++ {
    fmt.Scanf("%d %d %f %f", &PA, &PB, &GA, &GB)
    res := calcularCrescimento(PA, PB, GA, GB)
    if res > 100 {
      fmt.Printf("Mais de 1 seculo.\n")
    } else {
      fmt.Printf("%d anos.\n", res)
    }
  }
}

func calcularCrescimento(PA, PB int, GA, GB float64) int {
  res := 0
  for PA <= PB && res <= 100 {
    PA += int(GA * float64(PA) / 100.0)
    PB += int(GB * float64(PB) / 100.0)
    res++
  }
  return res
}

package main

import (
  "fmt"
)

func main() {
  var x int
  fmt.Scan(&x)

  fmt.Println(x)
  fmt.Printf("%d nota(s) de R$ 100,00\n", x / 100)

  v50 := x % 100
  fmt.Printf("%d nota(s) de R$ 50,00\n", v50 / 50)

  v20 := v50 % 50
  fmt.Printf("%d nota(s) de R$ 20,00\n", v20 / 20)

  v10 := v20 % 20
  fmt.Printf("%d nota(s) de R$ 10,00\n", v10 / 10)

  v5 := v10 % 10
  fmt.Printf("%d nota(s) de R$ 5,00\n", v5 / 5)

  v2 := v5 % 5
  fmt.Printf("%d nota(s) de R$ 2,00\n", v2 / 2)

  v1 := v2 % 2
  fmt.Printf("%d nota(s) de R$ 1,00\n", v1 / 1)

}

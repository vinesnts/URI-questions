package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scan(&n)
  pokedex := make([]string, n)

  for i := 0; i < n; i++ {
    fmt.Scan(&pokedex[i])
  }
  fmt.Printf("Falta(m) %d pomekon(s).\n", contar(pokedex, n))
}

func contar(pokedex []string, n int) (int) {
  falta := 151
  repetidos := 0
  flag := true
  for i := 0; i < n; i++ {
    flag = true
    for j := 0; j < i; j++ {
      if pokedex[i] == pokedex[j] {
        flag = false
        break
      }
    }
    if(flag) {
      repetidos++
    }
  }
  return falta - repetidos

}

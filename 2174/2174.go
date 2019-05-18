package main

import (
  "fmt"
)

const total = 151
var cont int = 0

func main() {
  var n int
  var p string
  pokedex := make([]string, total)
  fmt.Scan(&n)
  for i := 0; i < n; i++ {
    fmt.Scan(&p)
    if !buscarPokemon(pokedex, p) {
      pokedex[cont] = p
      cont++
    }

  }
  fmt.Printf("Falta(m) %d pokemon(s).\n", total - quant)
}

func buscarPokemon(p []string, s string) bool {
  for i := 0; i < len(p); i++ {
    if p[i] == s {
      return true
    }
  }
  return false
}

package main

import (
  "fmt"
)

const nletras = 5

func main() {
  n := 1
  for n != 0 {
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
      v := make([]int, nletras)
      fmt.Scanf("%d %d %d %d %d", &v[0], &v[1], &v[2], &v[3], &v[4])
      fmt.Printf("%s\n", string(analisar(v)))
    }
  }
}

func analisar(v []int) byte {
  marcadas := 0
  var letra int
  bool := make([]bool, nletras)
  for i := 0; i < len(v); i++ {
    if v[i] <= 127 {
      bool[i] = true
      letra = i
      marcadas++
    } else {
      bool[i] = false
    }
  }
  return processar(letra, marcadas)
}

func processar(letra, marcadas int) byte {
  if marcadas != 1 {
    return '*'
  } else {
    if letra == 0 {
      return 'A'
    } else if letra == 1 {
      return 'B'
    } else if letra == 2 {
      return 'C'
    } else if letra == 3 {
      return 'D'
    } else if letra == 4 {
      return 'E'
    } else {
      return '*'
    }
  }
}

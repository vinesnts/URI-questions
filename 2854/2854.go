package main

import (
  "fmt"
)

func new2DSlice(n, m int) [][]string {
  slice := make([][]string, n)
    for i := 0; i < len(slice); i++ {
        slice[i] = make([]string, 0, m)
    }
    return slice
}

var cont int = 0

func main() {
  var m, n int
  var nome1, relacao, nome2 string

  fmt.Scanf("%d %d", &m, &n)
  slice := new2DSlice(n, m)

  for i := 0; i < n; i++ {
    fmt.Scanf("%s %s %s", &nome1, &relacao, &nome2)
    pessoa1, _ := buscar(slice, nome1)
    pessoa2, _ := buscar(slice, nome2)
    //fmt.Printf("%s %s\n", nome1, nome2)
    if pessoa1 == -1 && pessoa2 == -1 {
      slice = add(slice, nome1, cont)
      slice = add(slice, nome2, cont)
      cont++
    } else if pessoa1 != -1 && pessoa2 == -1 {
      slice = add(slice, nome2, pessoa1)
    } else if pessoa2 != -1 && pessoa1 == -1 {
      slice = add(slice, nome1, pessoa2)
    } else if pessoa1 != -1 && pessoa2 != -1 && pessoa1 != pessoa2 {
      j := len(slice[pessoa2])
      for i := 0; i < j; i++ {
        slice = add(slice, slice[pessoa2][0], pessoa1)
        slice = removeIndex(slice, pessoa2, 0)
      }
    }
  }

  fmt.Println(contarRelacoes(slice))
}

func contarRelacoes(slice [][]string) int {
  cont := 0
  for i := 0; i < len(slice); i++ {
    if len(slice[i]) != 0 {
      cont++
    }
  }
  return cont
}

func add(slice [][]string, s string, pos int) [][]string {
  slice[pos] = append(slice[pos], s)
  return slice
}

func removeIndex(slice [][]string, i, j int) [][]string {
  pedaco := slice[i]
  pedaco = append(pedaco[:j], pedaco[j+1:]...)
  slice[i] = pedaco
  return slice
}

func buscar(slice [][]string, s string) (int, int) {
  for i := 0; i < len(slice); i++ {
    for j := 0; j < len(slice[i]); j++ {
      if slice[i][j] == s {
        return i, j
      }
    }
  }
  return -1, -1
}

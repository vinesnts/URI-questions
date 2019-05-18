package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scan(&n)

  for n != 0 {
    slice := make([]int, 0)
    for i := 0; i < n; i++ {
      var m int
      fmt.Scan(&m)
      slice = append(slice, m)
    }

    res := solitario(slice)
    if res == -1 {
      fmt.Println(0)
    /*  if slice[n - 1] != slice[n - 2] && slice[n - 2] == slice[n - 3] {
        res = n - 1
      }
      if slice[n - 1] == slice[n - 2] && slice[n - 2] == slice[n - 3] {
        res = n - 1
      }*/
    } else {
      fmt.Println(slice[res])
    }
    fmt.Scan(&n)
  }


}

func solitario(slice []int) (int) {
  solitario := -1
  for i := 0; i < len(slice); i++ {
    if buscar(slice, slice[i]) % 2 != 0 {
      return i
    }
  }
  return solitario
}

func buscar(slice []int, m int) int {
  cont := 0
  for i := 0; i < len(slice); i++ {
    if slice[i] == m {
      cont++
    }
  }
  return cont
}

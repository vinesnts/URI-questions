package main

import (
  "fmt"

)

func main() {
  var n1 string
  var n2 string
  fmt.Scan(&n1, &n2)

  for n1 != "0" && n2 != "0" {
    fmt.Printf("%s\n", verificar(n1, n2))
    fmt.Scan(&n1, &n2)
  }
}

func verificar(n1 string, n2 string) []byte {
  var n3 []byte
  i := 0
  for i < len(n2) {
    if string(n2[i]) != n1 {
      n3 = append(n3, n2[i])
      i++
    } else {
      for i != len(n2) && (string(n2[i]) == "0" || string(n2[i]) == n1) {
        i++
      }
    }
  }

  if n3 == nil {
    n3 = append(n3, '0')
  }
  return n3
}

package main

import (
  "fmt"
  "unicode"
)

var count int = 0

func main() {

  fmt.Scan(&count)
  for i := 0; i < count; i++ {
    var mensagem string
    fmt.Scan(&mensagem)
    fmt.Println(decoder(mensagem))
  }
}

func decoder(s string) string {
  res := ""
  for i := len(s)-1; i >= 0; i-- {
    if unicode.IsLower(rune(s[i])) {
      res = res + string(s[i])
    }
  }
  return res
}

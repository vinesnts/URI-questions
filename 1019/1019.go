package main

import (
  "fmt"
)

func main() {
  var segundos int
  fmt.Scan(&segundos)

  horas := segundos / 60 / 60
  minutos := (segundos / 60) % 60
  segundos = (segundos % 60) % 60

  fmt.Printf("%d:%d:%d\n", horas, minutos, segundos)

}

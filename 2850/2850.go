package main

import (
  "fmt"
  "os"
  "bufio"

)

func main() {

  var input string
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    input = scanner.Text()

    if input == "esquerda" {
      fmt.Println("ingles")
    } else if input == "direita" {
      fmt.Println("frances")
    } else if input == "nenhuma" {
      fmt.Println("portugues")
    } else if input == "as duas" {
      fmt.Println("caiu")
    } else {
      continue
    }
  }
}

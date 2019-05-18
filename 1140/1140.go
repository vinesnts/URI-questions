package main

import(
  "fmt"
  "strings"
  "bufio"
  "os"
)

func main() {
  input := ""
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    input = scanner.Text()
    if input == "*" {
      return
    }
    if(eTautograma(input)) == true {
      fmt.Println("Y")
    } else {
      fmt.Println("N")
    }
  }
}

func eTautograma(s string) bool {
  s = strings.ToLower(s)
  var vetor []string = strings.SplitAfter(s, " ")
  prefix := vetor[0][0]
  for i := 0; i < len(vetor); i++ {
    if vetor[i][0] != prefix {
      return false
    }
  }
  return true
}

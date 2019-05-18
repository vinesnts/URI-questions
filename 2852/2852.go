package main

import (
  "fmt"
  "strings"
)

var cifra = make([][]byte, 26, 26)

func inicializarCifra() {
  for i := 0; i < len(cifra[0]); i++ {
    
  }
}


func main() {

}

func criptografar(mensagem, chave string) string {
  palavras := strings.Split(mensagem, " ")
  for i := 0; i < len(palavras); i++ {
    if !vogalInicial(palavra[i]) {

    }
  }
}

func vogalInicial(palavra string) bool {
  if palavra[0] == 'a'
  || palavra[0] == 'e'
  || palavra[0] == 'i'
  || palavra[0] == 'o'
  || palavra[0] == 'u' {
    return true
  }
  return false
}

package main

import (
  "fmt"
)

type Aluno struct {
  nome, assinatura string
}

func main() {
  var n int
  fmt.Scan(&n)

  for n != 0 {
    var alunos []Aluno

    for i := 0; i < n; i++ {
      var aluno Aluno
      fmt.Scan(&aluno.nome, &aluno.assinatura)
      alunos = append(alunos, aluno)
    }

    var m int
    fmt.Scan(&m)
    nfalsas := 0

    for i := 0; i < m; i++ {
      var aluno Aluno
      fmt.Scan(&aluno.nome, &aluno.assinatura)
      if eFalsa(alunos, aluno) {
        nfalsas++
      }
    }

    fmt.Println(nfalsas)
    fmt.Scan(&n)
  }
}

func eFalsa(alunos []Aluno, aluno Aluno) bool {
  dif := 0
  for i := 0; i < len(alunos); i++ {
    if alunos[i].nome == aluno.nome {
      for j := 0; j < len(aluno.assinatura); j++ {
        if alunos[i].assinatura[j] != aluno.assinatura[j] {
          dif++
        }
      }
      if dif > 1 {
        return true
      }
    }
  }
  return false
}

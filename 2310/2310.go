package main

import (
  "fmt"
)

type Jogador struct {
  nome string
  s, b, a, ss, bs, as int
}


func calcPercentual(partida []Jogador) (float64, float64, float64) {
  var ts, tb, ta, tss, tbs, tas int
  for i := 0; i < len(partida); i++ {
    ts += partida[i].s
    tb += partida[i].b
    ta += partida[i].a

    tss += partida[i].ss
    tbs += partida[i].bs
    tas += partida[i].as
  }

  return float64(tss * 100) / float64(ts),
          float64(tbs * 100) / float64(tb),
          float64(tas * 100) / float64(ta)

}

func main() {
  var n int
  var partida []Jogador

  fmt.Scan(&n)
  for i := 0; i < n; i++ {
    var jogador Jogador
    fmt.Scan(&jogador.nome)
    fmt.Scanf("%d %d %d", &jogador.s, &jogador.b, &jogador.a)
    fmt.Scanf("%d %d %d", &jogador.ss, &jogador.bs, &jogador.as)
    partida = append(partida, jogador)
  }

  ps, pb, pa := calcPercentual(partida)

  fmt.Printf("Pontos de Saque: %0.2f %s.\n", ps, "%")
  fmt.Printf("Pontos de Bloqueio: %0.2f %s.\n",  pb, "%")
  fmt.Printf("Pontos de Ataque: %0.2f %s.\n", pa, "%")
}

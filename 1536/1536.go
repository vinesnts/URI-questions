package main

import (
  "fmt"
)

type Partida struct {
  time1, time2 int
}


func main() {
  var n int
  var partida1, partida2 Partida
  var s string

  fmt.Scan(&n)
  for i := 0; i < n; i++ {
    fmt.Scanf("%d %s %d", &partida1.time1, &s, &partida1.time2)
    fmt.Scanf("%d %s %d", &partida2.time2, &s, &partida2.time1)
    analisarResultados(partida1, partida2)
  }
}

func analisarResultados(partida1, partida2 Partida) {
  partidas := [2]Partida{partida1, partida2}
  var pontos1, pontos2, gols1, gols2, i int
  for k := 0; k < len(partidas); k++ {
    gols1 += partidas[i].time1
    gols2 += partidas[i].time2
    if partidas[i].time1 > partidas[i].time2 {
      pontos1 += 3
    } else if partidas[i].time1 < partidas[i].time2 {
      pontos2 += 3
    } else {
      pontos1 += 1
      pontos2 += 1
    }
    i++
  }
  if pontos1 > pontos2 {
    fmt.Println("Time 1")
  } else if pontos1 < pontos2 {
    fmt.Println("Time 2")
  } else {
    if gols1 - gols2 > gols2 - gols1 {
      fmt.Println("Time 1")
    } else if gols2 - gols1 > gols1 - gols2 {
      fmt.Println("Time 2")
    } else {
      if partida1.time2 > partida2.time1 {
        fmt.Println("Time 2")
      } else if partida2.time1 > partida1.time2 {
        fmt.Println("Time 1")
      } else {
        fmt.Println("Penaltis")
      }
    }
  }
}

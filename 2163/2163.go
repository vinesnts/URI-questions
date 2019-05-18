package main

import (
  "fmt"
)

func new2DSlice(n, m int) [][]int {
  slice := make([][]int, n)
    for i := 0; i < len(slice); i++ {
        slice[i] = make([]int, m)
    }
    return slice
}

func main() {
  var n, m int

  fmt.Scan(&n, &m)
  terreno := new2DSlice(n, m)

  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      fmt.Scanf("%d", &terreno[i][j])
    }
  }

  fmt.Println(encontrarSabre(terreno, n, m))
}

func encontrarSabre(terreno [][]int, n, m int) (int, int) {
  var x, y int

  for i := 1; i < (n - 1); i++ {
    for j := 1; j < (m - 1); j++ {
      if terreno[i][j] == 42 {
        x, y = i, j
        if verificarArredores(terreno, x, y) {
          return x + 1, y + 1
        }
      }
    }
  }
  return 0, 0
}

func verificarArredores(terreno [][]int, x, y int) bool {
  if (terreno[x - 1][y] == 7 &&
    terreno[x - 1][y - 1] == 7 &&
    terreno[x - 1][y + 1] == 7 &&
    terreno[x + 1][y] == 7 &&
    terreno[x + 1][y - 1] == 7 &&
    terreno[x + 1][y + 1] == 7) &&
    terreno[x][y - 1] == 7 &&
    terreno[x][y + 1] == 7 {
    return true
  }
  return false
}

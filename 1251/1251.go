package main

import (
  "fmt"
  "bufio"
  "os"
)

type bytes struct {
  valor byte
  quant int
}


func main() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  var valores []bytes
  n := []byte(scanner.Text())
  for i := 0; i < len(n); i++ {
    if !contains(valores, n[i]) {
      valores = append(valores, containMany(n, n[i]))
    }
  }
  printBytes(selectionSort(valores))
  for scanner.Scan() {
    fmt.Println("")
    var valores []bytes
    n = []byte(scanner.Text())
    for i := 0; i < len(n); i++ {
      if !contains(valores, n[i]) {
        valores = append(valores, containMany(n, n[i]))
      }
    }
    printBytes(selectionSort(valores))
  }
}

func printBytes(slice []bytes) {
  for i := 0; i < len(slice); i++ {
    fmt.Printf("%d %d\n", slice[i].valor, slice[i].quant)
  }
}

func selectionSort(slice []bytes) []bytes {
  for i := 0; i < len(slice); i++ {
    maior := &slice[i]
    for j := 0; j < len(slice); j++ {
      if slice[j].quant > maior.quant {
        slice = swap(slice, i, j)
      }
    }
  }
  sortByASCII(slice)
  return slice
}

func sortByASCII(slice []bytes) []bytes {
  for i := 0; i < len(slice); i++ {
    menor := &slice[i]
    for j := 0; j < len(slice); j++ {
      if slice[j].quant != menor.quant {
        continue
      }
      if slice[j].valor < menor.valor {
        slice = swap(slice, i, j)
      }
    }
  }
  return slice
}

func swap(slice []bytes, i , j int) []bytes {
  slice[i], slice[j] = slice[j], slice[i]
  return slice
}

func contains(slice []bytes, v byte) bool {
  for i := 0; i < len(slice); i++ {
    if slice[i].valor == v && slice[i].quant >= 1 {
      return true
    }
  }
  return false
}

func containMany(slice []byte, v byte) bytes {
  var res bytes
  res.valor = v
  for i := 0; i < len(slice); i++ {
    if slice[i] == v {
      res.quant++
    }
  }
  return res
}

func removeIndex(slice []byte, index int) []byte {
  return append(slice[:index], slice[index+1:]...)
}

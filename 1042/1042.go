package main

import (
  "fmt"
)

func main() {
  var a, b, c int
  fmt.Scanf("%d %d %d", &a, &b, &c)

  if a < b && b < c {
    fmt.Printf("%d\n%d\n%d\n\n", a, b, c)
  } else if a < c && c < b {
    fmt.Printf("%d\n%d\n%d\n\n", a, c, b)
  } else if b < c && c < a {
    fmt.Printf("%d\n%d\n%d\n\n", b, c, a)
  } else if b < a && a < c {
    fmt.Printf("%d\n%d\n%d\n\n", b, a, c)
  } else if c < b && b < a {
    fmt.Printf("%d\n%d\n%d\n\n", c, b, a)
  } else if c < a && a < b {
    fmt.Printf("%d\n%d\n%d\n\n", c, a, b)
  }
  fmt.Printf("%d\n%d\n%d\n", a, b, c)
}

package main

import (
  "./minhash"
  "fmt"
  )

func main() {
  fmt.Println(minhash.Minhash(
    []string{"1", "歳jk", "ビル", "む", "アヒ", "リルート", "携",},
    []string{"21", "歳", "ビール", "飲む", "アサヒ", "リクルート", "連携",}))
}

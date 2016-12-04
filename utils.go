package main

import (
  "time"
  "math/rand"
)

func randId(a int) int {
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)
  return r1.Intn(a)
}

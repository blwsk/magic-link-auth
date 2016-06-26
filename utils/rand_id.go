package utils

import (
  "time"
  "math/rand"
)

func RandId(a int) int {
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)
  return r1.Intn(a)
}

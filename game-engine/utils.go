package main

import (
  "crypto/rand"
  "fmt"
)

func randToken() string {
  b := make([]byte, 8)
  rand.Read(b)
  return fmt.Sprintf("%x", b)
}

func truncateInt(number, min, max int) int {
  if number > max {
    return max
  }
  if number < min {
    return min
  }
  return number
}

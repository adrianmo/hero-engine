package main

import (
  "crypto/rand"
  "fmt"
  "math"
  "time"
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

func getTTLForLevel(level int) time.Duration {
  return time.Duration(levelUpSeconds * (math.Pow(levelUpBase, float64(level))))
}

func ttlToDatetime(ttl time.Duration) time.Time {
  return time.Now().Add(ttl * time.Second)
}

package main

import (
  "time"
  "math/rand"
)

func randomString(strlen int) string {

  rand.Seed(time.Now().UTC().UnixNano())
  const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
  result := make([]byte, strlen)
  for i := 0; i < strlen; i++ {
    result[i] = chars[rand.Intn(len(chars))]
  }
  return string(result)
}


func randomNumber(min int, max int) int {

  rand.Seed(time.Now().UTC().UnixNano())
  return int(rand.Float64()*(float64(max) - float64(min)) + float64(min))
}
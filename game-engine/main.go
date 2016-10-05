package main

import (
  "fmt"
  "runtime"
)

func main() {
  configRuntime()
  StartGame()
}

func configRuntime() {
  nuCPU := runtime.NumCPU()
  runtime.GOMAXPROCS(nuCPU)
  fmt.Printf("Running with %d CPUs\n", nuCPU)
}

package main

import (
  "fmt"
  "runtime"
)

func main() {
  configRuntime()
  StartEngine()
  StartAPI()
}

func configRuntime() {
  nuCPU := runtime.NumCPU()
  runtime.GOMAXPROCS(nuCPU)
  fmt.Printf("Running with %d CPUs\n", nuCPU)
}

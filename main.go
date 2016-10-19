package main

import (
  "fmt"
  "os"
  "runtime"

  log "github.com/Sirupsen/logrus"
)

func init() {
  log.SetLevel(log.DebugLevel)
}

func main() {

  adminToken := os.Getenv("ADMIN_TOKEN")
  if len(adminToken) <= 0 {
    adminToken = randToken()
    log.Printf("Admin token not defined. Using '%s'", adminToken)
  }

  configRuntime()
  StartGame(adminToken)
}

func configRuntime() {
  nuCPU := runtime.NumCPU()
  runtime.GOMAXPROCS(nuCPU)
  fmt.Printf("Running with %d CPUs\n", nuCPU)
}

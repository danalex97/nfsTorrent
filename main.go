package main

import (
  "github.com/danalex97/nfsTorrent/simulation"
  "math/rand"
  "time"
  "fmt"
  "os"
)

func main() {
  rand.Seed(time.Now().UTC().UnixNano())

  s := simulation.SmallTorrentSimulation(
    new(simulation.SimulatedCachedNode),
    nil,
  )
  s.Run()

  time.Sleep(time.Duration(float64(time.Second) * 10))
  fmt.Println("Done")
  s.Stop()

  os.Exit(0)
}

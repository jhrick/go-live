package main

import (
	"github.com/jhrick/go-live/internal/reader"
	"github.com/jhrick/go-live/internal/server"
)

func main() {
  server := server.NewServer(reader.Read())
  go server.Start()
  
  change := make(chan bool, 1)

  for {
    go reader.WatchChanges(server.LastReload, change)

    <-change
    server.Content = reader.Read()
  }
}

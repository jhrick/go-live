package main

import (
	"github.com/jhrick/go-live/internal/reader"
	"github.com/jhrick/go-live/internal/server"
)

func main() {
  fileContent := reader.Read()

  server.NewServer(fileContent)
}

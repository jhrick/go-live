package main

import (
	"fmt"
	"os"
)

func main() {
  dir, _ := os.Getwd()

  content, err := os.ReadDir(dir)

  if err != nil {
    panic(err)
  }

  for _, file := range content {
    stat, err := os.Stat(file.Name())
    if err != nil {
      panic(err)
    }

    if !stat.IsDir() {
      if file.Name() == "index.html" {
        data, err := os.ReadFile(file.Name())
        if err != nil {
          panic(err)
        }

        fmt.Println(string(data))

        os.Exit(0)
      } else {
        fmt.Println(file.Name())
      }
    }
  }
}

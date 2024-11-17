package reader

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var currFile string

func Read() []byte {
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
        currFile = filepath.Join(dir, file.Name())

        data, err := os.ReadFile(file.Name())
        if err != nil {
          panic(err)
        }

        return data 
      } else {
        fmt.Println(file.Name())
      }
    }
  }

  return nil
}

func WatchChanges(lastReload time.Time, change chan bool) {
  for {
    info, err := os.Stat(currFile)
    if err != nil {
      panic(err)
    }

    if info.ModTime().Sub(lastReload) > 0 {
      change <- true
    }

    time.Sleep(500 * time.Millisecond)
  }
}

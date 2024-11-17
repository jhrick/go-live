package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type server struct {
  Server     *http.Server
  Handler    *http.ServeMux
  LastReload time.Time
  Content    []byte
}

func NewServer(content []byte) server {
  handler := http.NewServeMux()
  
  s := &http.Server{
    Addr: ":8080",
    Handler: handler,
  }

  return server{
    Server: s,
    Handler: handler,
    LastReload: time.Now(),
    Content: content,
  }
}

func (s *server) Start() {
  s.Handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
    w.Header().Set("Content-Type", "text/html")
    w.Write(s.Content)

    s.LastReload = time.Now()
  })

  fmt.Println("Server Running!")
  log.Fatal(s.Server.ListenAndServe())
}

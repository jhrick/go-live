package server

import (
	"fmt"
	"log"
	"net/http"
)

func NewServer(mainFile []byte) {
  mux := http.NewServeMux()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
    w.Header().Set("Content-Type", "text/html")
    w.Write(mainFile)
  })

  fmt.Println("Server Running!")
  log.Fatal(http.ListenAndServe(":8080", mux))
}

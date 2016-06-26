package main

import (
  "log"
  "net/http"

  "github.com/gorilla/mux"

  "github.com/blwsk/ginger/handlers"
)

type Server struct {
  router  *mux.Router
}

func (server *Server) ServeHTTP(resWriter http.ResponseWriter, req *http.Request) {
  // http://stackoverflow.com/a/248186381

  origin := req.Header.Get("Origin")

  if origin != "" {
    resWriter.Header().Set("Access-Control-Allow-Origin", origin)
    resWriter.Header().Set("Access-Control-Allow-Methods",
      "POST, GET, OPTIONS, PUT, DELETE")
    resWriter.Header().Set("Access-Control-Allow-Headers",
      "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
  }

  // Stop here if its Preflighted OPTIONS request
  if req.Method == "OPTIONS" {
    return
  }

  // Lets Gorilla work
  server.router.ServeHTTP(resWriter, req)
}

func main() {
  router := mux.NewRouter().StrictSlash(true)

  router.HandleFunc("/", handlers.IndexHandler)
  router.HandleFunc("/posts", handlers.PostsIndexHandler)
  router.HandleFunc("/posts/{postId}", handlers.PostHandler)

  http.Handle("/", &Server{ router })

  log.Fatal(http.ListenAndServe(":8080", nil))
}

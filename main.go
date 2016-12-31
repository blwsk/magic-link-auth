package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  // "github.com/blwsk/ginger/db"
)

type Server struct {
  router    *mux.Router
}

func (server *Server) ServeHTTP(resWriter http.ResponseWriter, req *http.Request) {
  origin := req.Header.Get("Origin")

  if origin != "" {
    resWriter.Header().Set("Access-Control-Allow-Origin",
      origin)
    resWriter.Header().Set("Access-Control-Allow-Methods",
      "POST, GET, OPTIONS, PUT, DELETE")
    resWriter.Header().Set("Access-Control-Allow-Headers",
      "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
  }

  if req.Method == "OPTIONS" {
    return
  }

  server.router.ServeHTTP(resWriter, req)
}

func buildRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)

  router.HandleFunc("/", IndexHandler).Methods("GET")
  router.HandleFunc("/posts", PostsIndexHandler).Methods("GET")
  router.HandleFunc("/posts/{id}", PostHandler).Methods("GET")
  router.HandleFunc("/auth/{hash}", AuthHandler).Methods("GET")
  router.HandleFunc("/protected", IsAuthenticated(ProtectedHandler)).Methods("GET")

  return router
}

func main() {
  router := buildRouter()

  http.Handle("/", &Server{ router })

  log.Fatal(http.ListenAndServe(":8080", nil))
}

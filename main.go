package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "database/sql"

  "github.com/blwsk/ginger/data"
)

type Server struct {
  router    *mux.Router
}

var (
  DbConn *sql.DB
)

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

  router.HandleFunc("/auth/{hash}", AuthHandler).Methods("GET")
  router.HandleFunc("/protected", IsAuthenticated(ProtectedHandler)).Methods("GET")
  router.HandleFunc("/magic-link", MagicLinkHandler).Methods("POST")

  return router
}

func main() {
  var err error
  DbConn, err = data.ConnectToDb()

  if err != nil {
    fmt.Print(err)
    log.Fatal(err)
  }

  err = DbConn.Ping()

  if err != nil {
    fmt.Print(err)
    log.Fatal(err)
  }

  router := buildRouter()

  http.Handle("/", &Server{ router })

  log.Fatal(http.ListenAndServe(":8080", nil))
}

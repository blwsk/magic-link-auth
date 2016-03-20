package main

import (
  "fmt"
  "encoding/json"
  "time"
  "math/rand"
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

func randId(a int) int {
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)
  return r1.Intn(a)
}

type Post struct {
  Title   string  `json:"title"`
  Id      int     `json:"id"`
}

type Posts struct {
  PostArray   []Post  `json:"posts"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Index")
}

func PostsIndexHandler(w http.ResponseWriter, r *http.Request) {
  res1 := Post{"Tour de France 2016", randId(100000)}
  res2 := Post{"Paris Roubaix 2016", randId(100000)}

  m := Posts{
    []Post{
      res1,
      res2,
    },
  }

  blob, _ := json.Marshal(m)

  fmt.Fprintln(w, string(blob))
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  postId := vars["postId"]

  fmt.Fprintln(w, "Post")
  fmt.Fprintln(w, "id:", postId)
}

type MyServer struct {
  r *mux.Router
}

// http://stackoverflow.com/a/24818638
func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
  if origin := req.Header.Get("Origin"); origin != "" {
      rw.Header().Set("Access-Control-Allow-Origin", origin)
      rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
      rw.Header().Set("Access-Control-Allow-Headers",
          "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
  }

  // Stop here if its Preflighted OPTIONS request
  if req.Method == "OPTIONS" {
      return
  }

  // Lets Gorilla work
  s.r.ServeHTTP(rw, req)
}

func main() {
  router := mux.NewRouter().StrictSlash(true)

  router.HandleFunc("/", IndexHandler)
  router.HandleFunc("/posts", PostsIndexHandler)
  router.HandleFunc("/posts/{postId}", PostHandler)

  http.Handle("/", &MyServer{router})

  log.Fatal(http.ListenAndServe(":8080", nil))
}

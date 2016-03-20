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
  res1 := Post{"Tour de France 2015", randId(100000)}
  res2 := Post{"Paris Roubaix 2015", randId(100000)}

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

func main() {
  router := mux.NewRouter().StrictSlash(true)

  router.Headers("Access-Control-Allow-Origin", "*")

  router.HandleFunc("/", IndexHandler)
  router.HandleFunc("/posts", PostsIndexHandler)
  router.HandleFunc("/posts/{postId}", PostHandler)

  log.Fatal(http.ListenAndServe(":8080", router))
}

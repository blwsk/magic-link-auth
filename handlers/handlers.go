package handlers

import (
  "net/http"
  "fmt"
  "encoding/json"

  "github.com/gorilla/mux"

  "github.com/blwsk/ginger/utils"
  "github.com/blwsk/ginger/data"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Index")
}

func PostsIndexHandler(w http.ResponseWriter, r *http.Request) {
  res1 := data.Post{"Tour de France 2016", utils.RandId(100000)}
  res2 := data.Post{"Paris Roubaix 2016", utils.RandId(100000)}

  m := data.Posts{
    []data.Post{
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

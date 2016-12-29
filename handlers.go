package main

import (
  "os"
  "net/http"
  "fmt"
  "strings"
  "encoding/json"
  "github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Index")
}

func PostsIndexHandler(w http.ResponseWriter, r *http.Request) {
  res1 := Post{
    "Tour de France 2016",
    randId(100000),
  }

  res2 := Post{
    "Paris Roubaix 2016",
    randId(100000),
  }

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

func VarHandler(w http.ResponseWriter, r *http.Request) {
  vars := os.Environ()

  v := make([]Var, len(vars))

  for i, e := range vars {
    c := strings.Split(e, "=")

    v[i] = Var{
      c[0],
      c[1],
    }
  }

  x := Vars{
    v,
  }

  blob, _ := json.Marshal(x)

  fmt.Fprintln(w, string(blob))
}

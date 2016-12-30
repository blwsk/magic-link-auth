package main

import (
  // "os"
  "net/http"
  "fmt"
  // "strings"
  "time"
  "encoding/json"
  "github.com/gorilla/mux"
)

const cookieName string = "_krb_cookiee"

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

// func VarHandler(w http.ResponseWriter, r *http.Request) {
//   vars := os.Environ()

//   v := make([]Var, len(vars))

//   for i, e := range vars {
//     c := strings.Split(e, "=")

//     v[i] = Var{
//       c[0],
//       c[1],
//     }
//   }

//   x := Vars{
//     v,
//   }

//   blob, _ := json.Marshal(x)

//   fmt.Fprintln(w, string(blob))
// }

func AuthHandler(w http.ResponseWriter, r *http.Request) {
  c := http.Cookie{
    Name: cookieName,
    Value: "hello",
    Expires: time.Now().Add(time.Hour),
    HttpOnly: false,
    MaxAge: 50000,
    Path: "/",
  }

  http.SetCookie(w, &c)

  fmt.Fprintln(w, true)
}

func UnauthenticatedHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "failure")
}

func Authenticated(f http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if _, err := r.Cookie(cookieName); err == nil {
      f(w, r)
    } else {
      UnauthenticatedHandler(w, r)
    }
  }
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "success")
}

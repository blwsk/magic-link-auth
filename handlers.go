package main

import (
  "fmt"
  "time"
  "net/http"
  "encoding/json"

  "github.com/gorilla/mux"
)

const authCookieName string = "_krb_cookie"

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

func AuthHandler(w http.ResponseWriter, r *http.Request) {
  token, err := CreateAuthToken("test@test.com")

  if err != nil {
    UnauthenticatedHandler(w, r)
  }

  c := http.Cookie{
    Name: authCookieName,
    Value: token,
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

func isValidCookie(c *http.Cookie) bool {
  if c.Name != authCookieName {
    return false
  }

  v := HasValidAuthToken(c.Value)

  fmt.Println(v)

  return v
}

func IsAuthenticated(f http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    c, err := r.Cookie(authCookieName)

    if err == nil && c != nil && isValidCookie(c) {
      f(w, r)
    } else {
      // cookie not present or invalid
      UnauthenticatedHandler(w, r)
    }
  }
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "success")
}

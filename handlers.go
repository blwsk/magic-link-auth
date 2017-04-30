package main

import (
  "fmt"
  // "log"
  "time"
  "net/http"
  "encoding/json"
  // "database/sql"
  _ "github.com/lib/pq"
  "github.com/gorilla/mux"

  "github.com/blwsk/ginger/auth"
  "github.com/blwsk/ginger/data"
)

const authCookieName string = "_krb_cookie"

func AuthHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  hash := vars["hash"]

  if auth.IsValidHash(hash) == false {
    fmt.Fprintln(w, "failed to set cookie, must fix")
    return
  }

  token, err := auth.CreateAuthToken("test@test.com")

  if err != nil {
    fmt.Fprintln(w, "failed to create auth token")
    return
  }

  c := http.Cookie{
    Name: authCookieName,
    Value: token,
    Expires: time.Now().Add(time.Hour),
    HttpOnly: false,
    MaxAge: 50000,
    Path: "/",
    Domain: ".blwsk.com",
  }

  http.SetCookie(w, &c)

  fmt.Fprintln(w, "cookie set")
}

func UnauthenticatedHandler(w http.ResponseWriter, r *http.Request) {
  rec := "k@blwsk.com"
  err := auth.SendAuthEmail(rec)

  blob := data.Action{
    Type: "SENT_AUTH_EMAIL",
    Payload: rec,
  }

  m, err := json.Marshal(blob)

  if err != nil {
    fmt.Fprintln(w, "Try auth-ing again, maybe?")
  }

  fmt.Fprintln(w, string(m))
}

func isValidCookie(c *http.Cookie) bool {
  if c.Name != authCookieName {
    return false
  }

  v := auth.HasValidAuthToken(c.Value)

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

func SaveMagicString(email string, hash string) (string, error) {
  // save in postgres

  stmt, err := DbConn.Prepare(
    `INSERT INTO "magic_string" (magic_string, expires_at, email) VALUES ($1, $2, $3);`)

  _, err = stmt.Exec(hash, time.Now().Add(time.Minute * 15), email)

  if err != nil {
    return "Problem", err
  }

  return "Success", err
}

func SendMagicStringEmail(email string, hash string) error {
  return nil
}

func MagicLinkHandler(w http.ResponseWriter, r *http.Request) {
  params := r.URL.Query()

  emailVals := params["email"]

  if len(emailVals) == 0 {
    fmt.Fprintln(w, "Error: must provide an email query parameter")
    return
  }

  email := emailVals[0]

  hash, err := auth.GenerateHashString()

  blah, err := SaveMagicString(email, hash)

  if err != nil {
    fmt.Fprintln(w, err)
    return
  }

  fmt.Fprintln(w, blah)
}

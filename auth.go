package main

import (
  "fmt"
  "time"
  "github.com/dgrijalva/jwt-go"
)

type AuthClaims struct {
  User string `json:"user"`
  jwt.StandardClaims
}

var secretKey interface{} = []byte("AuthKey?")

func getSecretKey(token *jwt.Token) (interface{}, error) {
  return secretKey, nil
}

func CreateAuthToken(user string) (string, error) {
  standardClaims := jwt.StandardClaims{
    ExpiresAt: time.Now().Add(time.Hour).Unix(),
    Issuer:    "test",
  }

  claims := AuthClaims{
    user,
    standardClaims,
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  return token.SignedString(secretKey)
}

func HasValidAuthToken(v string) bool {
  token, err := jwt.ParseWithClaims(v, &AuthClaims{}, getSecretKey)

  if err != nil {
    fmt.Println(err)
    return false
  }

  if _, ok := token.Claims.(*AuthClaims); ok && token.Valid {
    return true
  } else {
    return false
  }
}

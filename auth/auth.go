package auth

import (
  "os"
  "fmt"
  "time"
  "errors"
  "github.com/dgrijalva/jwt-go"
  "github.com/nu7hatch/gouuid"

  "github.com/blwsk/ginger/email"
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

func GenerateHash() (*uuid.UUID, error) {
  return uuid.NewV4()
}

func GenerateHashString() (string, error) {
  hash, err := uuid.NewV4()

  if err != nil {
    return "", err;
  }

  return hash.String(), nil
}

func SendAuthEmail(recipient string, authHash string) error {
  pass := os.Getenv("EMAIL_PASS")

  if pass != "" {
    return actuallySendAuthEmail(recipient, authHash)
  }

  return errors.New("Failed to send email")
}

func actuallySendAuthEmail(recipient string, authHash string) error {
  sender := email.NewSender(os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_PASS"))

  recipients := []string{
    recipient,
  }

  subject := "Click on the magic link to login"
  body := "<html><body>Click <a href=\"https://kbielawski.com/auth?hash=" +
    authHash + "\">here</a></body></html>"

  err := sender.SendMail(
    recipients,
    subject,
    body)

  return err
}

package data

import (
  "os"
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)

var (
  host     = os.Getenv("POSTGRES_HOST")
  port     = os.Getenv("POSTGRES_PORT")
  user     = os.Getenv("POSTGRES_USER")
  password = os.Getenv("POSTGRES_PASSWORD")
  dbname   = os.Getenv("POSTGRES_DBNAME")
)

func ConnectToDb() (*sql.DB, error) {
  psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
    host, port, user, password, dbname)

  return sql.Open("postgres", psqlInfo)
}

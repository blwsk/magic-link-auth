package db

import (
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)

const (
  DB_USER     = "test"
  DB_PASSWORD = "test"
  DB_NAME     = "test"
)

func ConnectToDb() (*sql.DB, error) {
  dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
  db, err := sql.Open("postgres", dbinfo)

  return db, err
}

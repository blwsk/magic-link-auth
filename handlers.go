package main

import (
	"fmt"
	// "log"
	"net/http"
	"time"
	// "encoding/json"
	// "database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/blwsk/ginger/auth"
	// "github.com/blwsk/ginger/data"
)

const authCookieName string = "_krb_cookie"

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]

	email, err := getEmailFromHash(hash)

	if err != nil || email == "" {
		w.WriteHeader(404)
		fmt.Fprintln(w, "Invalid hash")
		return
	}

	token, err := auth.CreateAuthToken(email)

	if err != nil {
		fmt.Fprintln(w, "failed to create auth token")
		return
	}

	c := http.Cookie{
		Name:     authCookieName,
		Value:    token,
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: false,
		MaxAge:   50000,
		Path:     "/",
		Domain:   "kbielawski.com",
	}

	http.SetCookie(w, &c)

	fmt.Fprintln(w, "cookie set")
}

func UnauthenticatedHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://kbielawski.com/login", 300)
}

func isValidCookie(c *http.Cookie) bool {
	if c.Name != authCookieName {
		return false
	}

	v := auth.HasValidAuthToken(c.Value)

	fmt.Println(v)

	return v
}

func getEmailFromHash(hash string) (string, error) {
	rows, err := DbConn.Query(
		`SELECT email FROM magic_string WHERE magic_string = $1 LIMIT 1;`, hash)

	var email string

	for rows.Next() {
		err = rows.Scan(&email)
	}

	if err != nil {
		return "Error", err
	} else {
		return email, err
	}
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

func SaveMagicString(email string, hash string) error {
	stmt, err := DbConn.Prepare(
		`INSERT INTO "magic_string" (magic_string, expires_at, email) VALUES ($1, $2, $3);`)

	_, err = stmt.Exec(
		hash,
		time.Now().Add(time.Minute*15),
		email)

	return err
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

	err = SaveMagicString(email, hash)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	err = auth.SendAuthEmail(email, hash)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintln(w, "Success")
}

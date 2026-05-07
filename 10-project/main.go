package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

var db *sql.DB
var sessionMap = make(map[string]int)

func init() {
	var err error
	connectionString := "user=poastgres password=password dbname=users sslmode=disable"
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connesso a Postgres")
}

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			t, _ := template.ParseFiles("login.html")
			t.Execute(w, nil)
		}
		fmt.Fprintf(w, "WIP")

		name := r.FormValue("name")
		password := r.FormValue("password")

		var user User
		err := db.QueryRow("select id, name, email from users where name = $1 and password = $2", name, password).Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			http.Error(w, "Credenziali non valide", 400)
		}

		sessionID := uuid.New().String()
		cookie := &http.Cookie{
			Name:  "session_id",
			Value: sessionID,
		}

		http.SetCookie(w, cookie)
		sessionMap[sessionID] = user.ID
		http.Redirect(w, r, "/user", http.StatusSeeOther)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "WIP")
	})
	http.ListenAndServe(":8080", nil)
}

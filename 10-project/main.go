package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

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

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			t, _ := template.ParseFiles("login.html")
			t.Execute(w, nil)
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		var userID int
		err := db.QueryRow("insert into users(name, email, password) values ($1, $2, $3) returning id", name, email, password).Scan(&userID)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Errore durante registrazione", 500)
			return
		}

		fmt.Fprintf(w, "Utente registrato con successo")
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "Non autenticato", http.StatusUnauthorized)
			return
		}

		userID, ok := sessionMap[cookie.Value]
		if !ok {
			http.Error(w, "Sessione non valida", http.StatusUnauthorized)
		}

		var user User
		switch r.Method {
		case http.MethodGet:
			err := db.QueryRow("select id, name, email from user where id = $1", userID).Scan(&user.ID, &user.Name, &user.Email)
			if err != nil {
				http.Error(w, "Utente non trovato", http.StatusNotFound)
				return
			}
			t, err := template.ParseFiles("user.html")
			t.Execute(w, user)

		case http.MethodPost:
			newName := r.FormValue("name")
			newPassword := r.FormValue("password")

			err := db.QueryRow("update users set name = $1, password = $2 from user where id = $3", newName, newPassword, userID)
			if err != nil {
				http.Error(w, "Errore aggiornamento utente", http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "Utente aggiornato")

		default:
			http.Error(w, "Metodo non supportato", http.StatusNotImplemented)
		}
	})

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "Non autenticato", http.StatusUnauthorized)
			return
		}

		delete(sessionMap, cookie.Value)

		http.SetCookie(w, &http.Cookie{
			Name:    "session_id",
			Value:   "",
			Expires: time.Unix(0, 0),
		})

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})

	http.ListenAndServe(":8080", nil)
}

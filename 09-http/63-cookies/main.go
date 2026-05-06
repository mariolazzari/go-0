package main

import (
	"fmt"
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour)

	cookie := http.Cookie{
		Name:    "FirstVisit",
		Value:   "1",
		Expires: expiration,
	}

	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "Cookie impostato. Per favore vai alla pagina /check-cookie per verificare il cookie.")
}

func checkCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("FirstVisit")

	if cookie != nil {
		fmt.Fprintf(w, "Bentornato! Hai visitato questo sito %s volta\n", cookie.Value)
	} else {
		fmt.Fprintf(w, "Benvenuto! Non hai mai visitato questo sito prima o il tuo cookie è scaduto.")
	}
}

func main() {
	http.HandleFunc("/set-cookie", setCookie)
	http.HandleFunc("/check-cookie", checkCookie)

	fmt.Println("Server avviato su localhost:8080")

	http.ListenAndServe(":8080", nil)
}

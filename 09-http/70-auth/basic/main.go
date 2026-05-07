package main

import (
	"fmt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok || username != "admin" || password != "password" {
		w.Header().Set("WWW-Authenticate", `Basic realm="Inserire utente e password"`)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Non sei autorizzato")
	}
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.ListenAndServe(":8080", nil)
}

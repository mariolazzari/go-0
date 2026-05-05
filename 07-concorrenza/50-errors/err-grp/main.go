package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	var urls = []string{
		"http://www.google.com",
		"http://www.yahoo.com",
		"http://www.bing.com",
	}
	for _, url := range urls {
		// Copia l'URL corrente nella variabile locale per evitare problemi di sincronizzazione dei dati.
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Recuperati con successo tutti gli URL.")
	} else {
		fmt.Println("Errore riscontrato:", err)
	}
}

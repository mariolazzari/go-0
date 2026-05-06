package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		length, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Ricevuto messaggio dal client: %s\n", string(buffer[:length]))

		message := []byte("Messaggio ricevuto")
		_, err = conn.WriteToUDP(message, clientAddr)
		if err != nil {
			log.Fatal(err)
		}
	}
}

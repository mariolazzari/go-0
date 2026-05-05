package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func gestConn(conn net.Conn) {
	lettore := bufio.NewReader(conn)
	messaggio, err := lettore.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Messaggio dal client:", messaggio)

	risposta := "Messaggio ricevuto"
	conn.Write([]byte(risposta))

	fmt.Println("Messaggio inviato al client:", risposta)
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go gestConn(conn)

	}
}

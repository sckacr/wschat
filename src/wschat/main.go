package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var connections = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func main() {
	go handleMessages()
	http.Handle("/", websocket.Handler(handleConnections))
	log.Fatal(http.ListenAndServe(":12345", nil))
}

func handleConnections(conn *websocket.Conn) {
	connections[conn] = true
	for {
		var msg Message
		if err := websocket.JSON.Receive(conn, &msg); err != nil {
			log.Println("[ERROR] websocket.JSON.Receive: " + err.Error())
			delete(connections, conn)
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for conn := range connections {
			if _, err := conn.Write([]byte(fmt.Sprintf("%v: %v\n", msg.Name, msg.Message))); err != nil {
				log.Println("[ERROR] ws.Conn.Write: " + err.Error())
				delete(connections, conn)
				break
			}
		}
	}
}

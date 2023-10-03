package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		// Read in a message
		massageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println("Yes :" + string(p))

		if err := conn.WriteMessage(massageType, p); err != nil {
			log.Println(err)
			return
		}
	}

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, My Name is Smith!")
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error in handler:", err)
		return
	}
	log.Println("Client connected.")

	reader(conn)
}

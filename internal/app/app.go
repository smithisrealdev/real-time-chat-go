package app

import (
	"fmt"
	"log"
	"net/http"
	"real-time-chat-go/internal/routes"
)

func Start() {
	fmt.Println("Starting ...")
	routes.Routes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

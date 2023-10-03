package routes

import "net/http"

func Routes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", websocketHandler)
}

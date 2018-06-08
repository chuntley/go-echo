package main

import (
	"io"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

func echoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func main() {
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	http.Handle("/", websocket.Handler(echoHandler))
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		panic("Error starting: " + err.Error())
	}
}

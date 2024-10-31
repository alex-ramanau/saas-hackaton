package main

import (
	// "flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ServerVariables struct{
	ListenPort string
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	//load the env vars
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	ListenPort := os.Getenv("DARK_NET_LISTEN_PORT")

	fmt.Println("The env var: ", ListenPort)

	// var addr = flag.String("addr", ":8080", "The application address")
	// flag.Parse()

	// Create new Chat room
	room := newRoom()

	// Init the chat room as a coroutine
	go room.run()
	
	//TOOD

	// Main endpoint for the client to connect to the room
	http.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		serveWs(room, w, r) 
	})
	
	// and init the web-server
	// TODO: Maybe this needs to be started upfront
	log.Println("Starting web server on", ListenPort)
	conn := http.ListenAndServe(ListenPort, nil)
	if conn != nil {
		log.Fatalf("Failed to start server: %v", conn)
	}
}
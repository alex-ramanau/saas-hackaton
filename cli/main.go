package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func registerUser() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your name: ")
	// name, err := reader.ReadString('\n')
	// if err != nil {
	// 	log.Fatal("Error reading name:", err)
	// }

	// // send request to server to register user
	// fmt.Print("User registered successfully: ", name)
}

func startChat() {
	// Connect to WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer conn.Close()

	// Start a goroutine to receive messages
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}
			fmt.Println("User:", string(message))
		}
	}()

	// Send messages from CLI
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("")
		message, _ := reader.ReadString('\n')

		fmt.Print("\033[1A\033[K") // clear the input of this message

		conn.WriteMessage(websocket.TextMessage, []byte(message))

	}
}

func main() {
	registerUser()
	startChat()
}

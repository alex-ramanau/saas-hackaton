package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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
			_, recv, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}

			// convert emojis syntax to actual emojis
			message := replaceEmojis(string(recv))

			fmt.Println("User:", message)
		}
	}()

	// Send messages from CLI
	reader := bufio.NewReader(os.Stdin)
	for {
		message, _ := reader.ReadString('\n')

		// clear the input of this message so it doesn't show up in the chat duplicated with the same server message
		fmt.Print("\033[1A\033[K")

		conn.WriteMessage(websocket.TextMessage, []byte(message))

	}
}

// list of allowed emojis using :emoji: syntax
var emojiMap = map[string]string{
	":smile:":  "😊",
	":laugh:":  "😂",
	":scared:": "😱",
	":sad:":    "😢",
	":angry:":  "😠",
}

func replaceEmojis(message string) string {
	for k, v := range emojiMap {
		message = strings.ReplaceAll(message, k, v)
	}
	return message
}

func main() {
	registerUser()
	startChat()
}

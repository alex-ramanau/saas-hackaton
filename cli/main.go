package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

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

			// instead of lack of features, we like to call it a "privacy feature"! everyone is anonymous :)
			fmt.Println("Anon:", message)
		}
	}()

	// Send messages from CLI
	reader := bufio.NewReader(os.Stdin)
	for {
		message, _ := reader.ReadString('\n')

		// check if message is empty
		if message == "\n" {
			fmt.Print("\033[1A\033[K")
			continue
		}

		// Check if the message is a command
		if strings.HasPrefix(message, "/") {
			changeColor(message)
			continue // Skip sending the command as a message
		}

		// clear the input of this message so it doesn't show up in the chat duplicated with the same server message
		fmt.Print("\033[1A\033[K")

		conn.WriteMessage(websocket.TextMessage, []byte(message))

	}
}

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

var colorMap = map[string]string{
	"/red":     Red,
	"/green":   Green,
	"/yellow":  Yellow,
	"/blue":    Blue,
	"/magenta": Magenta,
	"/cyan":    Cyan,
	"/white":   White,
}

func changeColor(command string) {
	color := strings.TrimSpace(strings.ToLower(command))
	if colorCode, exists := colorMap[color]; exists {
		fmt.Print(colorCode)
		return
	}
	fmt.Print(Reset)
}

func replaceEmojis(message string) string {
	for k, v := range emojiMap {
		message = strings.ReplaceAll(message, k, v)
	}
	return message
}

func main() {
	// Handle SIGINT and SIGTERM
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// goroutine to handle signals and reset color on exit
	go func() {
		<-signalChan // Wait for an interrupt signal
		fmt.Println()
		fmt.Print("\033[1A\033[K") // so ^C doesnt show up
		fmt.Print(Reset)           // Reset color on exit
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()

	registerUser()
	startChat()
}

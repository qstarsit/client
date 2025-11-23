package client

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	ServerAddress string
}

func (c *Client) Connect() *websocket.Conn {
	conn, response, err := websocket.DefaultDialer.Dial(c.serverAddress, nil)
	if err != nil {
		fmt.Printf("Failed to connect to WebSocket: %v", err)
	}

	/*
		You can only defer a function.
		conn.Close returns an error
		so we wrap it in a function
	*/
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("Error closing connection: %v", err)
		}
	}(conn)

	// Check if the response status is 101 (Switching Protocols)
	if response.StatusCode != 101 {
		fmt.Printf("Unexpected response status: %d\n", response.StatusCode)
	}

	return conn
}

func (c *Client) SendMessage(conn *websocket.Conn, message string) {
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		fmt.Printf("Failed to send message: %v", err)
	}
}

func (c *Client) ReadMessage(conn *websocket.Conn) (string, error) {
	messageType, message, err := conn.ReadMessage()
	if err != nil {
		return "", fmt.Errorf("failed to read message: %w", err)
	}

	// Check if it's a text message
	if messageType != websocket.TextMessage {
		return "", fmt.Errorf("unexpected message type: %d", messageType)
	}

	return string(message), nil
}

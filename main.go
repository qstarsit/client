package client

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	ServerAddress string
}

// Connect Establishes a WebSocket connection to the server and returns the connection object.
func (c *Client) Connect() *websocket.Conn {
	conn, response, err := websocket.DefaultDialer.Dial(c.ServerAddress, nil)
	if err != nil {
		fmt.Printf("Failed to connect to WebSocket: %v", err)
	}

	// Check if the response status is 101 (Switching Protocols)
	if response.StatusCode != 101 {
		fmt.Printf("Unexpected response status: %d\n", response.StatusCode)
	}

	return conn
}

// SendMessage Sends a text message over the WebSocket connection.
func (c *Client) SendMessage(conn *websocket.Conn, message string) {
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		fmt.Printf("Failed to send message: %v", err)
	}
}

// ReadMessage Reads a text message from the WebSocket connection.
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

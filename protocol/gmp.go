package protocol

import (
	"net"
	"log"
	"bufio"
	"strings"
	"fmt"
)

type EventType int

const (
	UNKNOWN EventType = iota
	PUBLISH
	CONSUME
)

func (e EventType) String() string {
	switch e {
	case PUBLISH :
		return "PUBLISH"

	case CONSUME :
		return "CONSUME"

	default:
		return "UNKNOWN"
	}
}

type Publish struct {
	event EventType
	key string
	payload string // a serialized json string
}

func connect() {
	listener, err := net.Listen("tcp", "8090")

	if err != nil {
		log.Fatal("Error listening: ", err)
	}

	defer listener.Close() //close connection at the end

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println("Failed accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()

	reader := bufio.NewReader(conn)

	message, err := reader.ReadString('\n')

	if err != nil {
        log.Printf("Read error: %v", err)
        return
    }
// Create and send a response back to the client, demonstrating two-way communication.
    ackMsg := strings.ToUpper(strings.TrimSpace(message))
    response := fmt.Sprintf("ACK: %s\n", ackMsg)
    _, err = conn.Write([]byte(response))
    if err != nil {
        log.Printf("Server write error: %v", err)
    }


}
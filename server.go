package main

import (
	"encoding/json"
	"io"
	"log"
	"net"

	"github.com/Microsoft/go-winio"
)

func Callme() string {
	return "Hello from server"
}
func Callme2() string {
	return "Hello from server 2"
}

type Msg struct {
	ID         uint64 `json:"id"`
	Msg        string `json:"msg"`
	Name       string `json:"name"`
	ActionFunc string `json:"action"`
}

func handleClient(c net.Conn) {
	defer c.Close()
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())

	// buf := make([]byte, 512)
	// for {
	// 	n, err := c.Read(buf)
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			log.Printf("read error: %v\n", err)
	// 		}
	// 		break
	// 	}
	// 	str := string(buf[:n])
	// 	log.Printf("read %d bytes: %q\n", n, str)
	// }
	// log.Println("Client disconnected")
	// read json bytes from client
	buf := make([]byte, 512)
	for {
		n, err := c.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("read error: %v\n", err)
			}
			break
		}
		var msg Msg
		err = json.Unmarshal(buf[:n], &msg)
		if err != nil {
			log.Printf("json unmarshal error: %v\n", err)
			continue
		}
		//  call function based on action without

		log.Printf("Received message: %+v\n", msg)
	}
	log.Println("Client disconnected")
}

func main() {
	pipePath := `\\.\pipe\mypipename`

	l, err := winio.ListenPipe(pipePath, nil)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()
	log.Printf("Server listening op pipe %v\n", pipePath)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go handleClient(conn)
	}
}

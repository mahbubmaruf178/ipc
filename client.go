package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Microsoft/go-winio"
)

type Msg struct {
	ID         uint64 `json:"id"`
	Msg        string `json:"msg"`
	Name       string `json:"name"`
	ActionFunc string `json:"action"`
}

func main() {
	pipePath := `\\.\pipe\mypipename`
	f, err := winio.DialPipe(pipePath, nil)
	if err != nil {
		log.Fatalf("error opening pipe: %v", err)
	}
	defer f.Close()
	// n, err := f.Write([]byte("message from client!"))
	// if err != nil {
	// 	log.Fatalf("write error: %v", err)
	// }
	// log.Println("wrote:", n)
	// get input from user
	for {
		var input string
		log.Println("Enter message to send to server: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatalf("error reading input: %v", err)
		}
		newmsg := Msg{2, input, input, "Callme2"}
		jsondata, err := json.Marshal(newmsg)
		if err != nil {
			log.Fatalf("error marshalling json: %v", err)

		}

		n, err := f.Write(jsondata)
		if err != nil {
			log.Fatalf("write error: %v", err)
		}
		log.Println("wrote:", n)
	}
}

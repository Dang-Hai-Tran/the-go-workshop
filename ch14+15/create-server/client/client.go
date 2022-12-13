package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type message struct {
	Message string `json: "message"`
}

func getDataAndReturnResponse() message {
	url := "http://localhost:8080"
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	// Read data
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	mes := message{}
	// Parse data into a message struct
	err = json.Unmarshal(data, &mes)
	if err != nil {
		log.Fatal(err)
	}
	return mes
}

func sendData() {
	newMessage := message{Message: "Hi server"}
	url := "http://localhost:8080"
	jsonByte, _ := json.Marshal(newMessage)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonByte))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Send data success.")
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func main() {
	sendData()
}

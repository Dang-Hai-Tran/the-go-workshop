package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type server struct{}
type message struct {
	Message string `json : "message"`
}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	messageData := message{}
	jsonBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonBytes, &messageData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messageData.Message)
	w.Write([]byte("Hi client"))
}

func main() {
	port := 8080
	fmt.Printf("Server listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server{}))
}

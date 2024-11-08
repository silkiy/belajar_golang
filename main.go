package main

import (
	"fmt"
	"go-rest-api/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/messages", handlers.GetAllMessages)
	http.HandleFunc("/message", handlers.GetMessage)
	http.HandleFunc("/message/create", handlers.CreateMessage)
	http.HandleFunc("/message/update", handlers.UpdateMessage)
	http.HandleFunc("/message/delete", handlers.DeleteMessage)

	fmt.Println("Server berjalan di port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

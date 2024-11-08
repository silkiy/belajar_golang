package main

import (
	"fmt"
	"log"
	"net/http"
	"go-rest-api/handlers"
	"go-rest-api/utils"
)

func main() {
	// Inisialisasi koneksi ke database
	utils.InitDB()
	defer utils.CloseDB()

	// Menentukan route dan handler
	http.HandleFunc("/messages", handlers.GetAllMessages)        // Menampilkan semua message
	http.HandleFunc("/message", handlers.GetMessage)            // Menampilkan message berdasarkan ID
	http.HandleFunc("/message/create", handlers.CreateMessage)  // Menambah message
	http.HandleFunc("/message/update", handlers.UpdateMessage)  // Mengupdate message
	http.HandleFunc("/message/delete", handlers.DeleteMessage)  // Menghapus message

	// Menjalankan server di port 8080
	fmt.Println("Server berjalan di port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
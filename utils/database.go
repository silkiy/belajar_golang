package utils

import (
	"database/sql"
	"fmt"
	"go-rest-api/models"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Inisialisasi koneksi database
func InitDB() {
	var err error
	db, err = sql.Open("mysql", "root@tcp(localhost:3306)/messages")
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}
	fmt.Println("Koneksi ke database berhasil!")
}

// Tutup koneksi database
func CloseDB() {
	if err := db.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}

// Fungsi untuk mendapatkan semua messages
func GetAllMessages() ([]models.Message, error) {
	rows, err := db.Query("SELECT id, status, message FROM messages")
	if err != nil {
		return nil, fmt.Errorf("Error fetching messages: %v", err)
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.ID, &message.Status, &message.Message); err != nil {
			return nil, fmt.Errorf("Error scanning row: %v", err)
		}
		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %v", err)
	}

	return messages, nil
}

// Fungsi untuk mendapatkan message berdasarkan ID
func GetMessageByID(id int) (models.Message, error) {
	var message models.Message
	err := db.QueryRow("SELECT id, status, message FROM messages WHERE id = ?", id).Scan(&message.ID, &message.Status, &message.Message)
	if err != nil {
		if err == sql.ErrNoRows {
			return message, fmt.Errorf("Message not found with ID: %d", id)
		}
		return message, fmt.Errorf("Error fetching message by ID: %v", err)
	}
	return message, nil
}

// Fungsi untuk menambahkan message
func AddMessage(message models.Message) error {
	_, err := db.Exec("INSERT INTO messages (status, message) VALUES (?, ?)", message.Status, message.Message)
	if err != nil {
		return fmt.Errorf("Error adding message: %v", err)
	}
	return nil
}

// Fungsi untuk memperbarui message
func UpdateMessage(id int, updatedMessage models.Message) (bool, error) {
	result, err := db.Exec("UPDATE messages SET status = ?, message = ? WHERE id = ?", updatedMessage.Status, updatedMessage.Message, id)
	if err != nil {
		return false, fmt.Errorf("Error updating message: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Error checking rows affected: %v", err)
	}
	return rowsAffected > 0, nil
}

// Fungsi untuk menghapus message
func DeleteMessage(id int) (bool, error) {
	result, err := db.Exec("DELETE FROM messages WHERE id = ?", id)
	if err != nil {
		return false, fmt.Errorf("Error deleting message: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Error checking rows affected: %v", err)
	}
	return rowsAffected > 0, nil
}

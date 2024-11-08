package handlers

import (
	"encoding/json"
	"net/http"
	"go-rest-api/models"
	"go-rest-api/utils"
	// "fmt"
	"strconv"
)

// Handler untuk mendapatkan semua message
func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	messages, err := utils.GetAllMessages()
	if err != nil {
		http.Error(w, "Error fetching messages", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
}

// Handler untuk mendapatkan message berdasarkan ID
func GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	message, err := utils.GetMessageByID(id)
	if err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(message)
}

// Handler untuk membuat message baru
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestBody struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if requestBody.Status == "" || requestBody.Message == "" {
		http.Error(w, "Both status and message are required", http.StatusBadRequest)
		return
	}

	newMessage := models.Message{
		Status:  requestBody.Status,
		Message: requestBody.Message,
	}

	if err := utils.AddMessage(newMessage); err != nil {
		http.Error(w, "Error adding message", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"status":  "success",
		"message": "Message created successfully",
	}

	json.NewEncoder(w).Encode(response)
}

// Handler untuk memperbarui message
func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedMessage models.Message
	if err := json.NewDecoder(r.Body).Decode(&updatedMessage); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	updatedMessage.ID = id

	success, err := utils.UpdateMessage(id, updatedMessage)
	if err != nil || !success {
		http.Error(w, "Message not found or update failed", http.StatusNotFound)
		return
	}

	response := map[string]string{
		"status":  "success",
		"message": "Message updated successfully",
	}

	json.NewEncoder(w).Encode(response)
}

// Handler untuk menghapus message
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	success, err := utils.DeleteMessage(id)
	if err != nil || !success {
		http.Error(w, "Message not found or deletion failed", http.StatusNotFound)
		return
	}

	response := map[string]string{
		"status":  "success",
		"message": "Message deleted successfully",
	}

	json.NewEncoder(w).Encode(response)
}

package handlers

import (
	"encoding/json"
	"go-rest-api/models"
	"go-rest-api/utils"
	"net/http"
	"strconv"
)

type RequestBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	messages := utils.GetAllMessages()
	json.NewEncoder(w).Encode(messages)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	message, found := utils.GetMessageByID(id)
	if !found {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(message)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if requestBody.Status == "" || requestBody.Message == "" {
		http.Error(w, "Status and Message cannot be empty", http.StatusBadRequest)
		return
	}

	newMessage := models.Message{
		Status:  requestBody.Status,
		Message: requestBody.Message,
	}

	newMessage.ID = len(utils.GetAllMessages()) + 1
	utils.AddMessage(newMessage)

	response := map[string]string{
		"status":  "success",
		"message": "Message created successfully",
	}

	json.NewEncoder(w).Encode(response)
}

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

	existingMessage, found := utils.GetMessageByID(id)
	if !found {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	if updatedMessage.Status != "" {
		existingMessage.Status = updatedMessage.Status
	}

	if updatedMessage.Message != "" {
		existingMessage.Message = updatedMessage.Message
	}

	success := utils.UpdateMessage(id, existingMessage)
	if !success {
		http.Error(w, "Failed to update message", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"status":  "success",
		"message": "Message updated successfully",
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	success := utils.DeleteMessage(id)
	if !success {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	response := map[string]string{"status": "success", "message": "Message deleted successfully"}
	json.NewEncoder(w).Encode(response)
}

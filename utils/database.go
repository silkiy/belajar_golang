package utils

import "go-rest-api/models"

var database = []models.Message{
	{ID: 1, Status: "success", Message: "Welcome to Go REST API!"},
}

func GetAllMessages() []models.Message {
	return database
}

func GetMessageByID(id int) (models.Message, bool) {
	for _, message := range database {
		if message.ID == id {
			return message, true
		}
	}
	return models.Message{}, false
}

func AddMessage(message models.Message) {
	database = append(database, message)
}

func UpdateMessage(id int, updatedMessage models.Message) bool {
	for i, message := range database {
		if message.ID == id {
			database[i] = updatedMessage
			return true
		}
	}
	return false
}

func DeleteMessage(id int) bool {
	for i, message := range database {
		if message.ID == id {
			database = append(database[:i], database[i+1:]...)
			return true
		}
	}
	return false
}

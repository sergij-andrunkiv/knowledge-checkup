package services

import (
	"encoding/json"
	"knowledge_checkup/backend/dataStorage"
	"knowledge_checkup/backend/model"
	"net/http"
)

type MessageManager struct{}

// Додати повідомлення
func (m MessageManager) Push(r *http.Request, w http.ResponseWriter, status model.MessageType, title string, messageText string) {
	var messages []model.Message
	var message model.Message = model.Message{Title: title, Text: messageText, Status: status}
	currentMessages := "{}"
	session, _ := dataStorage.GetStore().Get(r, "user-data-session")

	if _, ok := session.Values["currentMessages"]; ok {
		currentMessages = session.Values["currentMessages"].(string)
	}

	json.Unmarshal([]byte(currentMessages), &messages)
	messages = append(messages, message)
	jsonString, _ := json.Marshal(messages)

	session.Values["currentMessages"] = string(jsonString)
	session.Save(r, w)
}

// Отримати всі повідомлення
func (m MessageManager) Flush(r *http.Request, w http.ResponseWriter) string {
	session, _ := dataStorage.GetStore().Get(r, "user-data-session")

	currentMessages := "{}"

	if _, ok := session.Values["currentMessages"]; ok {
		currentMessages = session.Values["currentMessages"].(string)
	}

	session.Values["currentMessages"] = "{}"
	session.Save(r, w)

	return currentMessages
}

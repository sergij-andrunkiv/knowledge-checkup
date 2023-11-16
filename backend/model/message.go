package model

const (
	MESSAGE_STATUS_SUCCESS = "success"
	MESSAGE_STATUS_FAILURE = "failure"
)

type MessageType string

type Message struct {
	Title  string
	Text   string
	Status MessageType
}

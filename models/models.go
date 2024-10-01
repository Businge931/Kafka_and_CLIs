package models

type Message struct {
	Channel  string            // The channel where the message is sent/received
	Server   string            // The server to communicate with
	Group    string            // The group identifier for the message
	Payload  string            // The actual message content being sent
	Metadata map[string]string // Optional: for additional context (e.g., headers)
}

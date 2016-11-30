package common

import "strings"

type Message struct {
	Identifier string
	Parameter  string
}

func ReadMessage(buf [256]byte) *Message {
	var message Message

	// Trim everything after \x00 (assuming both identifer and parameter are null-terminated strings)
	message.Identifier = strings.TrimRight(string(buf[:10]), "\x00")
	message.Parameter = strings.TrimRight(string(buf[11:]), "\x00")

	return &message
}

func WriteMessage(message *Message) []byte {
	var buf [256]byte

	copy(buf[:10], message.Identifier)
	copy(buf[11:], message.Parameter)

	return buf[:]
}

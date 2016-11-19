package common

type Message struct {
	Identifier string
	Parameter  string
}

func ReadMessage(buf [256]byte) *Message {
	return &Message{Identifier: string(buf[:10]), Parameter: string(buf[11:])}
}

func WriteMessage(message *Message) []byte {
	var buf [256]byte
	copy(buf[:10], message.Identifier)
	copy(buf[11:], message.Parameter)
	return buf[:]
}

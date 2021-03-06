package protocol

/*
Message represents a message that is
either transmitted or stored locally.
*/
type Message struct {
	SentAt  int64  `json:"sent"`
	Content string `json:"body"`
	User    *User  `json:"user"`
}

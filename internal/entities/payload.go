package entities

type Payload struct {
	UserID          string    `json:"user_id"`
	Username        string    `json:"username"`
	Action          string    `json:"action"`
	SendMessage     string    `json:"send_message"`
	ReceiveMessages []Message `json:"receive_messages"`
}

type Message struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	UserName string `json:"username"`
	Content  string `json:"content"`
}

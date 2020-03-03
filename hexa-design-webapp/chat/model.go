package chat

type Chat struct {
	Sender    string `json:"sender"`
	Recipent  string `json:"recipient"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp,omitIfBlank"`
}

type Messages struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp,omitIfBlank"`
}

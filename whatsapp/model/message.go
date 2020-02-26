package model

type Message struct {
	Text string `json:"text`
	ID   int    `json:"Id"`
}

type Chat struct {
	Sender    string `json:"sender"`
	Recipent  string `json:"recipient"`
	MessageId int    `json:"messageId"`
}

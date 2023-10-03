package port

type Message interface {
	createMessage()
}

type MessageService interface {
	PostMessage()
}

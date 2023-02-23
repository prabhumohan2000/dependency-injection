package depcontainer

type MessageService interface {
	SendEmail(messgae string) error
}

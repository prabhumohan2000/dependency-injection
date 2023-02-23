package depcontainer

import "log"

type EmailService struct{}

func (e EmailService) SendEmail(message string) error {
	log.Print("printing the message", message)
	return nil
}

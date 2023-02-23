package user

import (
	depcontainer "github.com/prabhumohan2000/dep/dep_container"
)

type User struct {
	MessageService depcontainer.MessageService
}

func (u User) sendEmailMessage(message string) error {
	return u.MessageService.SendEmail(message)
}

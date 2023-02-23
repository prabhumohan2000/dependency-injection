package main

import (
	depcontainer "github.com/prabhumohan2000/dep/dep_container"
	"github.com/prabhumohan2000/dep/user"
)

// create Instance of newUser
func NewUser(MessageService depcontainer.MessageService) user.User {
	return user.User{
		MessageService: MessageService,
	}
}

func main() {
	emailService := depcontainer.EmailService{}

	user_msg := NewUser(emailService)

	user_msg.MessageService.SendEmail("hello bro")

}

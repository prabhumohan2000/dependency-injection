package users

import "githu.com/prabhumohan2000/test_8/graph/model"

type UserService interface {
	GetUser() ([]*model.User, error)
	CreateUser(Id string, name string) (*model.User, error)
	UpdateUser(Id string, name string) (*model.User, error)
}

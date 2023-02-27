package graph

import (
	"githu.com/prabhumohan2000/test_8/users"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService users.UserService
}

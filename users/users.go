package users

import (
	"githu.com/prabhumohan2000/test_8/graph/model"
	"gorm.io/gorm"
)

//	type userRepository interface {
//		GetUser() ([]*model.User, error)
//	}
type UserServicefile struct {
	DB *gorm.DB
}

func (u UserServicefile) GetUser() ([]*model.User, error) {
	var user []*model.User
	// database.DbInstance.First(&user)
	u.DB.Last(&user)
	return user, nil
}

func (u UserServicefile) CreateUser(Id string, name string) (*model.User, error) {
	// u.DB.Create(&model.User{
	// 	ID:   Id,
	// 	Name: name,
	// })
	// return &model.User{ID: Id, Name: name}, nil
	// using gorm transcation package
	tx := u.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, err
	}
	user := &model.User{
		ID:   Id,
		Name: name,
	}
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	err := tx.Commit().Error
	if err != nil {
		return nil, err
	}
	return user, nil

}
func (u UserServicefile) UpdateUser(Id string, name string) (*model.User, error) {
	user := &model.User{
		ID: Id,
	}
	u.DB.Model(&user).UpdateColumn("name", name)
	return user, nil

}

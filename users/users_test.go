package users

import (
	"reflect"
	"testing"

	"githu.com/prabhumohan2000/test_8/graph/model"
	"gorm.io/gorm"
)

func TestUserServicefile_GetUser(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserServicefile{
				DB: tt.fields.DB,
			}
			got, err := u.GetUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserServicefile.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserServicefile.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserServicefile_CreateUser(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		Id   string
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserServicefile{
				DB: tt.fields.DB,
			}
			got, err := u.CreateUser(tt.args.Id, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserServicefile.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserServicefile.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserServicefile_UpdateUser(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		Id   string
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserServicefile{
				DB: tt.fields.DB,
			}
			got, err := u.UpdateUser(tt.args.Id, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserServicefile.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserServicefile.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

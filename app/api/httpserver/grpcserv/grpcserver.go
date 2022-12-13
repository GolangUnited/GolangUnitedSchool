package grpcserv

import (
	"context"
	"github.com/lozovoya/GolangUnitedSchool/app/api/httpserver/grpcserv/api"
)

type StudentServer struct {
	//api.UnimplementedStudentServiceServer
}

func (s *StudentServer) AddNewPerson(ctx context.Context, res *__.AddUserRequest) (*__.AddUserResponse, error) {

	return &__.AddUserResponse{
		Error:  "",
		UserId: 0,
	}, nil
}

func (s *StudentServer) GetPersons(context.Context, *__.Empty) (*__.PersonsList, error) {

	var kek []*__.User

	kk := NewUser()

	kek = append(kek, kk)

	return &__.PersonsList{
		PersonList: kek,
	}, nil
}

func NewUser() *__.User {
	return &__.User{
		FirstName:  "l",
		LastName:   "l",
		Login:      "l",
		Patronymic: "l",
		Email:      "l",
		Birthday:   "l",
		RoleId:     "ll",
		Passwd:     "l",
	}
}

package grpcserv

import (
	"context"
	"fmt"
	"github.com/lozovoya/GolangUnitedSchool/app/api/httpserver/grpcserv/api"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

type StudentServer struct{}

func (s *StudentServer) AddUser(ctx context.Context, res *api.AddUserRequest) (*api.AddUserResponse, error) {

	var AddingUser model.Person

	AddingUser.FirstName = res.FirtsName
	AddingUser.LastName = res.LastName
	AddingUser.Login = res.Login

	err := AddingUser.ValidatePerson()
	if err != nil {
		return nil, fmt.Errorf("bad validation: %w ", err)
	}

	return &api.AddUserResponse{
		UserAdded: res.LastName + res.LastName + " added to db",
		UserId:    0,
	}, nil
}

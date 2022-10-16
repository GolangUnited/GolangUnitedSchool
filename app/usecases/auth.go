package usecases

import "github.com/lozovoya/GolangUnitedSchool/app/repository"

type AuthSt struct {
	r repository.Repository
}

func NewAuth(r repository.Repository) Auth {
	return &AuthSt{
		r: r,
	}
}

func (a *AuthSt) Login() {

}

func (a *AuthSt) Logout() {

}

package usecases

import "github.com/lozovoya/GolangUnitedSchool/internal/repository"

type CourseSt struct {
	r repository.Repository
}

func NewCourseCases(r repository.Repository) Course {
	return &CourseSt{r: r}
}

package postgres

import (
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"golang.org/x/net/context"
)

func (r *PostgresRepository) GetCourseLectures(ctx context.Context) ([]model.CourseLecture, error) {
	panic("empty")
}
func (r *PostgresRepository) GetCourseLectureById(ctx context.Context, id int64) (*model.CourseLecture, error) {
	panic("empty")
}
func (r *PostgresRepository) AddCourseLecture(ctx context.Context, data *model.CourseLecture) error {
	panic("empty")
}
func (r *PostgresRepository) UpdateCourseLectureById(ctx context.Context, id int64, data *model.CourseLecture) error {
	panic("empty")
}
func (r *PostgresRepository) DeleteCourseLectureById(ctx context.Context, id int64) error {
	panic("empty")
}

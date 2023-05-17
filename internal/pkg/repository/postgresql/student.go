package postgresql

import (
	"context"
	"errors"
	"homework-5/internal/pkg/db"
	"homework-5/internal/pkg/repository"
)

type StudentRepo struct {
	db db.DBops
}

func NewStudent(db db.DBops) *StudentRepo {
	return &StudentRepo{db: db}
}

func (r *StudentRepo) Add(ctx context.Context, student *repository.Student) (uint64, error) {
	var id uint64
	err := r.db.ExecQueryRow(ctx, "INSERT INTO student(name, grades, univ_apply_id) VALUES ($1, $2, $3) RETURNING id", student.Name, student.Grades, student.UnivID).Scan(&id)
	return id, err
}
func (r *StudentRepo) GetById(ctx context.Context, id uint64) (*repository.Student, error) {
	students := make([]*repository.Student, 10)
	err := r.db.Select(ctx, &students, "select id, name, grades, univ_apply_id from student where id = $1", id)
	if len(students) != 0 {
		return students[0], err
	}
	return nil, errors.New("No data found by this id.")
}
func (r *StudentRepo) Update(ctx context.Context, data *repository.Student) (bool, error) {
	result, err := r.db.Exec(ctx, "update student set name = $1 where id = $2", data.Name, data.ID)
	return result.RowsAffected() > 0, err
}

func (r *StudentRepo) Delete(ctx context.Context, id uint64) (bool, error) {
	result, err := r.db.Exec(ctx, "DELETE FROM student where id=$1", id)
	if err != nil {
		return result.RowsAffected() > 0, err
	}
	return result.RowsAffected() > 0, err
}

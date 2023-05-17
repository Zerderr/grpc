package postgresql

import (
	"context"
	"errors"
	"homework-5/internal/pkg/db"
	"homework-5/internal/pkg/repository"
)

type UniversityRepo struct {
	db db.DBops
}

func NewUniversity(db db.DBops) *UniversityRepo {
	return &UniversityRepo{db: db}
}

func (r *UniversityRepo) Add(ctx context.Context, data *repository.University) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, "INSERT INTO university(univ_name, facility) VALUES ($1, $2) RETURNING id", data.Name, data.Facility).Scan(&id)
	return id, err
}
func (r *UniversityRepo) GetById(ctx context.Context, id uint64) (*repository.University, error) {
	universities := make([]*repository.University, 10)
	err := r.db.Select(ctx, &universities, "select id, univ_name, facility from university where id = $1", id)
	if len(universities) != 0 {
		return universities[0], err
	}
	return nil, errors.New("No data found by this id.")
}
func (r *UniversityRepo) Update(ctx context.Context, data *repository.University) (bool, error) {
	result, err := r.db.Exec(ctx, "update university set facility = $1 where id = $2", data.Facility, data.ID)
	return result.RowsAffected() > 0, err
}

func (r *UniversityRepo) Delete(ctx context.Context, id uint64) (bool, error) {
	result, err := r.db.Exec(ctx, "DELETE FROM university where id=$1", id)
	if err != nil {
		return result.RowsAffected() > 0, err
	}
	return result.RowsAffected() > 0, err
}

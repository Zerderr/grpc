//go:generate mockgen -source C:\Users\Artyom\gitlabozon\homework-5\internal\pkg\repository\repository.go -destination=C:\Users\Artyom\gitlabozon\homework-5\internal\pkg\repository\mocks\repository.go -package=mock_repository

package repository

import (
	"context"
)

type StudentRepo interface {
	Add(ctx context.Context, student *Student) (uint64, error)
	GetById(ctx context.Context, id uint64) (*Student, error)
	Update(ctx context.Context, student *Student) (bool, error)
	Delete(ctx context.Context, id uint64) (bool, error)
}

type UniversityRepo interface {
	Add(ctx context.Context, university *University) (int64, error)
	GetById(ctx context.Context, id uint64) (*University, error)
	Update(ctx context.Context, university *University) (bool, error)
	Delete(ctx context.Context, id uint64) (bool, error)
}

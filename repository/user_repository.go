package repository

import (
	"context"
	"github.com/trewanek/RepositoryPatternWithGolang/entity"
)

type UserRepository interface {
	List(ctx context.Context) ([]*entity.User, error)
	Get(ctx context.Context, id int64) (*entity.User, error)
	Insert(ctx context.Context, u *entity.User) (*entity.User, error)
	Update(ctx context.Context, u *entity.User) (*entity.User, error)
	Delete(ctx context.Context, id int64) error
}

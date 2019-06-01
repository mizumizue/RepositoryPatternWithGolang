package repository

import (
	"context"
	"errors"
	"github.com/trewanek/RepositoryPatternWithGolang/entity"
)

type UserInMemoryRepository struct {
	Users []*entity.User
}

func NewUserInMemoryRepository() *UserInMemoryRepository {
	return &UserInMemoryRepository{
		Users: []*entity.User{
			{
				ID:   1,
				Name: "Yamada Tarou",
				Age:  23,
			},
			{
				ID:   2,
				Name: "Suzuki Shigeru",
				Age:  41,
			},
			{
				ID:   3,
				Name: "Ohashi Kanako",
				Age:  24,
			},
			{
				ID:   4,
				Name: "Yukishiro Kyoko",
				Age:  32,
			},
		},
	}
}

func (r *UserInMemoryRepository) List(ctx context.Context) ([]*entity.User, error) {
	return r.Users, nil
}

func (r *UserInMemoryRepository) Get(ctx context.Context, id int64) (*entity.User, error) {
	for _, u := range r.Users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, nil
}

func (r *UserInMemoryRepository) Insert(ctx context.Context, u *entity.User) (*entity.User, error) {
	r.Users = append(r.Users, u)
	return u, nil
}

func (r *UserInMemoryRepository) Update(ctx context.Context, u *entity.User) (*entity.User, error) {
	for _, ru := range r.Users {
		if ru.ID == u.ID {
			ru.Name = u.Name
			ru.Age = u.Age
			return ru, nil
		}
	}
	return nil, errors.New("User not found. ")
}

func (r *UserInMemoryRepository) Delete(ctx context.Context, id int64) error {
	var deleted []*entity.User
	count := len(r.Users)
	for _, u := range r.Users {
		if u.ID == id {
			continue
		}
		deleted = append(deleted, u)
	}
	if count == len(deleted) {
		return errors.New("User not found. ")
	}
	r.Users = deleted
	return nil
}

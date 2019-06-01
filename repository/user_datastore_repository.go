package repository

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/trewanek/repositoryPattern/entity"
	"google.golang.org/api/iterator"
	"os"
)

type UserDatastoreRepository struct {
	ds *datastore.Client
}

const kind = "User"

type UserKind struct {
	Key  *datastore.Key `datastore:"__key__"`
	Name string
	Age  int
}

func (uk *UserKind) toEntity() *entity.User {
	return &entity.User{
		ID:   uk.Key.ID,
		Name: uk.Name,
		Age:  uk.Age,
	}
}

func NewUserDatastoreRepository(ctx context.Context) *UserDatastoreRepository {
	ds, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		panic("panic: datastore client don't create. ")
	}
	return &UserDatastoreRepository{
		ds: ds,
	}
}

func (r *UserDatastoreRepository) List(ctx context.Context) ([]*entity.User, error) {
	var list []*entity.User
	var err error
	q := datastore.NewQuery(kind)
	it := r.ds.Run(ctx, q)
	for {
		var uk UserKind
		key, err := it.Next(&uk)
		if err != nil {
			break
		}
		uk.Key = key
		list = append(list, uk.toEntity())
	}
	if err != nil && err != iterator.Done {
		return nil, err
	}
	return list, nil
}

func (r *UserDatastoreRepository) Get(ctx context.Context, id int64) (*entity.User, error) {
	uk := &UserKind{}
	err := r.ds.Get(ctx, datastore.IDKey(kind, int64(id), nil), uk)
	if err != nil {
		return nil, err
	}
	return uk.toEntity(), nil
}

func (r *UserDatastoreRepository) Insert(ctx context.Context, u *entity.User) (*entity.User, error) {
	uk := toKind(u)
	key, err := r.ds.Put(ctx, datastore.IncompleteKey(kind, nil), uk)
	if err != nil {
		return nil, err
	}
	uk.Key = key
	return uk.toEntity(), nil
}

func (r *UserDatastoreRepository) Update(ctx context.Context, u *entity.User) (*entity.User, error) {
	uk := &UserKind{
		Name: u.Name,
		Age:  u.Age,
	}
	key, err := r.ds.Put(ctx, getKeyFromID(u.ID), uk)
	if err != nil {
		return nil, err
	}
	uk.Key = key
	return uk.toEntity(), nil
}

func (r *UserDatastoreRepository) Delete(ctx context.Context, id int64) error {
	return r.ds.Delete(ctx, getKeyFromID(id))
}

func getKeyFromID(id int64) *datastore.Key {
	return datastore.IDKey(kind, id, nil)
}

func toKind(u *entity.User) *UserKind {
	return &UserKind{
		Name: u.Name,
		Age:  u.Age,
	}
}

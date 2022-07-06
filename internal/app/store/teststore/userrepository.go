package teststore

import (
	"go_task/internal/app/store"
	"go_task/internal/app/model"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil{
		return err
	}

	r.users[u.Mobile] = u
	u.ID = len(r.users)

	return nil
}

func (r *UserRepository) FindByMobile(mobile string) (*model.User, error) {
	for _, u := range r.users {
		if u.Mobile == mobile {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}
package sqlstore

import (
	"go_task/internal/app/store"
	"database/sql"
	"go_task/internal/app/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil{
		return err
	}

	return r.store.db.QueryRow("INSERT INTO users (name, mobile) VALUES ($1, $2) RETURNING id", u.Name, u.Mobile).Scan(&u.ID)
}

func (r *UserRepository) FindByMobile(mobile string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, name, mobile FROM users WHERE mobile = $1", mobile).Scan(&u.ID, &u.Name, &u.Mobile,); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
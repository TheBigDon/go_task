package store

import (
	"go_task/internal/app/model"
)

type UserRepository interface {
	Create(*model.User) error
	FindByMobile(string) (*model.User, error)
}
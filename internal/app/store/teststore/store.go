package teststore

import (
	"go_task/internal/app/model"
	"go_task/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
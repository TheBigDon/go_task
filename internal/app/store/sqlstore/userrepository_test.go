package sqlstore_test

import (
	"github.com/stretchr/testify/assert"
	"go_task/internal/app/model"
	"go_task/internal/app/store"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}
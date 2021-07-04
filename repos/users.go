package repos

import (
	"github.com/aambayec/tut-grpc-go-web/types"
	"github.com/go-xorm/xorm"
)

// UsersRepo - the users repo interface
type UsersRepo interface {
	Create(*types.User) error
}

// NewUsersRepo - returns a new user's repo
func NewUsersRepo(db *xorm.Engine) UsersRepo {
	return &usersRepo{db:db}
}

type usersRepo struct {
	db *xorm.Engine
}

func (u usersRepo) Create(user *types.User) (err error) {
	if err = types.Validate(user); err != nil {
		return
	}

	if _, err = u.db.Insert(user); err != nil {
		return
	}

	return
}
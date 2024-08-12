package domain

import "context"


const (
	CollectionUser = "users"
)

type User struct {
	ID       string `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Role     string `bson:"role,omitempty" json:"role"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Login(c context.Context, user *User) (string, error)
	Promote(c context.Context, id string) error
}


type UserUseCase interface {
	Create(c context.Context, user *User) error
	Login(c context.Context, user *User)(string, error)
	Promote(c context.Context, id string) error
}
package requests

import "card-game/consts"

type (
	CreateUserRequest struct {
		Name   string            `validate:"required,min=3,max=191,user-exist" json:"name"`
		Email  string            `validate:"required,email,user-exist" json:"email"`
		Status consts.UserStatus `validate:"required,oneof=1 2" json:"status"`
	}
)

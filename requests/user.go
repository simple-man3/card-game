package requests

import "card-game/consts"

type (
	CreateUserRequest struct {
		Name     string            `validate:"required,min=3,max=191,user-not-exist" json:"name"`
		Email    string            `validate:"required,email,user-not-exist" json:"email"`
		Password string            `validate:"required,min=6" json:"password"`
		Status   consts.UserStatus `validate:"required,oneof=1 2" json:"status"`
	}
	PatchUserRequest struct {
		Name   string            `validate:"omitempty,min=3,max=191,user-not-exist" json:"name"`
		Email  string            `validate:"omitempty,email,user-not-exist" json:"email"`
		Status consts.UserStatus `validate:"omitempty,oneof=1 2" json:"status"`
	}
)

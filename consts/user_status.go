package consts

type UserStatus int

const (
	Active UserStatus = iota + 1
	Banned
)

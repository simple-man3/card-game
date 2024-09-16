package dto

import "github.com/golang-jwt/jwt/v4"

type JwtPayload struct {
	Email string `json:"email"`

	jwt.RegisteredClaims
}

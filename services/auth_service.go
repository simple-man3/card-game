package services

import (
	"card-game/config"
	"card-game/dto"
	"card-game/models"
	"card-game/session"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	ttlJwtToken = time.Hour / 2
)

type AuthService struct {
	UserService *UserService
}

func NewAuthService() *AuthService {
	userService := NewUserService()

	return &AuthService{
		UserService: userService,
	}
}

func (as AuthService) Auth(c *fiber.Ctx, email, password string) (string, error) {
	user := models.User{
		Email: email,
	}

	if err := as.UserService.GetUser(&user, []string{}); err != nil {
		return "", err
	}

	if !as.checkPasswordHash(password, user.Password) {
		return "", errors.New("логин или пароль не действителен")
	}

	token, err := as.generateToken(email)
	if err != nil {
		return "", err
	}

	sess, _ := session.Store.Get(c)
	sess.Set("authToken", token)

	return token, nil
}

func (as AuthService) generateToken(email string) (string, error) {
	claims := dto.JwtPayload{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttlJwtToken)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.EnvInstance.JwtSecret))
}

func (as AuthService) checkPasswordHash(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func (as AuthService) VerifyToken(token string) (*dto.JwtPayload, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &dto.JwtPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.EnvInstance.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	jwtPayload := parsedToken.Claims.(*dto.JwtPayload)

	if time.Now().After(jwtPayload.RegisteredClaims.ExpiresAt.Time) {
		return nil, errors.New("token expired")
	}

	return jwtPayload, nil
}

func (as AuthService) AuthFromToken(token string) error {
	jwtPayload, err := as.VerifyToken(token)
	if err != nil {
		return err
	}

	user := &models.User{Email: jwtPayload.Email}
	if err = as.UserService.GetUser(user, []string{}); err != nil {
		return err
	}

	models.AuthUser = user

	return nil
}

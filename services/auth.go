package services

import (
	"card-game/config"
	"card-game/models"
	"card-game/session"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
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

	sess, _ := session.Session.Get(c)
	sess.Set("authToken", token)

	return token, nil
}

func (as AuthService) generateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["exp"] = time.Now().Add(ttlJwtToken).Unix()

	env, _ := config.GetInstanceEnv()

	return token.SignedString([]byte(env.JwtSecret))
}

func (as AuthService) checkPasswordHash(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

package services

import (
	"card-game/config"
	"card-game/models"
	"card-game/session"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	ttlJwtToken = time.Hour / 2
)

func Auth(c *fiber.Ctx, email, password string) (string, error) {
	user := models.User{
		Email: email,
	}

	if err := GetUser(&user, []string{}); err != nil {
		return "", err
	}

	if !checkPasswordHash(password, user.Password) {
		return "", errors.New("логин или пароль не действителен")
	}

	token, err := generateToken(email)
	if err != nil {
		return "", err
	}

	sess, _ := session.Session.Get(c)
	sess.Set("authToken", token)

	return token, nil
}

func generateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["exp"] = time.Now().Add(ttlJwtToken).Unix()

	env, _ := config.GetInstanceEnv()

	t, err := token.SignedString([]byte(env.JwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func checkPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	fmt.Println(err)
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

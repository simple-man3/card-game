package session

import (
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Session *session.Store

func InitSession() {
	Session = session.New()
}

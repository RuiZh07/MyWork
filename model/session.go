package model

import "github.com/gofiber/fiber/v2/middleware/session"

//global variable for session data
var (
	Store      *session.Store
	AUTH_KEY   string = "authenticated"
	USER_EMAIL string = "user_email"
)

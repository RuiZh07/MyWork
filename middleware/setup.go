// package middleware

// import (
// 	"time"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/cors"
// 	"github.com/gofiber/fiber/v2/middleware/session"
// )

// var (
// 	store *session.Store
// 	AUTH_key string = "authenticated"
// 	USER_ID string = "user_id"
// )

// func Setup() {
// 	store := session.New(session.Config{
// 		CookieHTTPOnly: true,
// 		Expiration: time.Hour * 24,
// 	})
// }
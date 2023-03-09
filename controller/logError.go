package controller

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func UnexpectedError(c *fiber.Ctx, err error, source string) error {
	logFile, err2 := os.OpenFile("database/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err2 != nil {
		log.Print(err2)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)
	logger.Printf("Error in %s: %v", source, err)

	return c.Render("unexpectedError", fiber.Map{
		"Error": err,
	})
}

func UnexpectedErrorForFunction(err error) error {
	logFile, err2 := os.OpenFile("database/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err2 != nil {
		log.Print(err2)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println(err)

	return err
}

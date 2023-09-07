package middlewares

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	startTime := time.Now()

	err := c.Next()

	fmt.Printf("[%s] %s %s - %s\n", startTime.Format("2006-01-02 15:04:05"), c.Method(), c.Path(), time.Since(startTime))

	return err
}

package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func LogRequestsMiddleware() fiber.Handler {
	return  func(ctx *fiber.Ctx) error {
		method := ctx.Method()
		ip := ctx.IP()
		message := "ip:" + " " + ip + " " + "method" + " " + method
		log.Print(message)
		return ctx.Next()
	}
}
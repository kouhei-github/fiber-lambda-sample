package route

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	FiberApp *fiber.App
}

type RouterImp interface {
	GetRouter()
	GetAuthRouter()
}

func LoadRouter(router RouterImp) {
	router.GetRouter()
	router.GetAuthRouter()
}

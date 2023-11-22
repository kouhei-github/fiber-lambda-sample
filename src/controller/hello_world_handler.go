package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type response struct {
	Message string `json:"message"`
	Body    string `json:"body"`
	Status  int16  `json:"status"`
}

type request struct {
	UserName string `json:"userName"`
}

// HealthCheckHandler　POSTリクエストを受け取る
func HealthCheckHandler(c *fiber.Ctx) error {
	// リクエストヘッダーの受け取り
	agent := c.Get("User-Agent")

	// リクエストボディの受け取り
	var req request
	if err := c.BodyParser(&req); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Jsonでレスポンスを返す
	res := response{Message: agent, Status: fiber.StatusAccepted, Body: req.UserName}
	return c.Status(fiber.StatusCreated).JSON(res)
}

func HelloHandler(c *fiber.Ctx) error {
	res := c.Response()
	res.Header.SetStatusCode(fiber.StatusOK)
	return c.SendString("Hello, World 👋!")
}

func PathParamTestHandler(c *fiber.Ctx) error {
	// パスパラメータの取得
	id := c.Params("id")
	return c.Status(fiber.StatusOK).JSON(map[string]string{"pathParameter": id})
}

func QueryParamTestHandler(c *fiber.Ctx) error {
	// クエリパラメータの取得
	id := c.Query("id")
	return c.Status(fiber.StatusOK).JSON(map[string]string{"queryParameter": id})
}

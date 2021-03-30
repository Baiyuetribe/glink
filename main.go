package main

import (
	"Baiyuetribe/glink/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Static("/", "dist")              // 静态文件
	app.Get("api/*", service.ApiHandler) // 请求地址 http://127.0.0.1:3000/api/http://demo.com
	log.Fatal(app.Listen("127.0.0.1:3000"))
}

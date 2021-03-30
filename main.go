package main

import (
	"Baiyuetribe/glink/service"
	"log"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	service.PrintLogo()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(compress.New())              // 压缩静态资源未gzip或br
	app.Static("/", "web")               // 静态文件
	app.Get("api/*", service.ApiHandler) // 请求地址 http://127.0.0.1:3000/api/http://demo.com
	// log.Fatal(app.Listen(":3006"))
	// fmt.Println("应用访问地址：http://127.0.0.1:3006")
	log.Fatal(app.Listen("127.0.0.1:3006"))
}

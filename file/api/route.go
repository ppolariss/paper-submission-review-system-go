package api

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Route("/api", func(router fiber.Router) {
		router.Post("/file/upload", UploadFile)
		router.Get("/file/preview/:essayId", PreviewFile)
		router.Get("/file/download/:essayId", DownloadFile)
	})
}

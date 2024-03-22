package api

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	app.Route("/api", func(router fiber.Router) {
		router.Post("/register", Register)
		router.Post("/login", Login)
		router.Post("/logout", Logout)
		router.Get("/userInfo/self", GetSelfUserInfo)
		router.Get("/userInfo/byUsername/:username", GetUserInfoByUsername)
		router.Get("/userInfo/byRealName/:realName", GetUserInfoByRealName)
		router.Get("/userConferenceRole/self/:conferenceName", GetSelfUserConferenceRoles)
	})
}

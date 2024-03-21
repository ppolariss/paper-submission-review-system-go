package api

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Route("/api", func(router fiber.Router) {
		app.Post("/conference/apply", Apply)
		app.Get("/conference/getAllConferences", GetAllConferences)
		app.Get("/conference/getAllMyAppliedConference", GetAllMyAppliedConference)
		app.Get("/conference/getAllMyJoinedConference", GetAllMyJoinedConference)
		app.Get("/conference/getMyRoleInConference/:conferenceName", GetMyRoleInConference)
		app.Get("/conference/getConferenceInfoByName/:conferenceName", GetConferenceInfoByName)
		app.Post("/conference/startSubmission", StartSubmission)
		app.Get("/conference/getAllConferenceApplications", GetAllConferenceApplications)
		app.Post("/conference/auditConference", AuditConference)
	})
}

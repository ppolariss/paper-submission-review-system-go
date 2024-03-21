package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "github.com/ppolariss/paper-submission-review-system-go/conference/schema"
)

func Apply(c *fiber.Ctx) error {
	var applyConferenceRequest ApplyConferenceRequest
	err := common.ValidateBody(c, &applyConferenceRequest)
	if err != nil {
		return err
	}
	return nil
}

func GetAllConferences(c *fiber.Ctx) error {
	return nil
}

func GetAllMyAppliedConference(c *fiber.Ctx) error {
	return nil
}

func GetAllMyJoinedConference(c *fiber.Ctx) error {
	return nil
}

func GetMyRoleInConference(c *fiber.Ctx) error {
	conferenceName := c.Params("conferenceName")
	if conferenceName == "" {
		return common.BadRequest("Invalid conference name")
	}
	return nil
}

func GetConferenceInfoByName(c *fiber.Ctx) error {
	conferenceName := c.Params("conferenceName")
	if conferenceName == "" {
		return common.BadRequest("Invalid conference name")
	}
	return nil
}

func StartSubmission(c *fiber.Ctx) error {
	var applyConferenceRequest ApplyConferenceRequest
	err := common.ValidateBody(c, &applyConferenceRequest)
	if err != nil {
		return err
	}
	return nil
}

func GetAllConferenceApplications(c *fiber.Ctx) error {
	return nil
}

func AuditConference(c *fiber.Ctx) error {
	var applyConferenceRequest ApplyConferenceRequest
	err := common.ValidateBody(c, &applyConferenceRequest)
	if err != nil {
		return err
	}
	return nil
}

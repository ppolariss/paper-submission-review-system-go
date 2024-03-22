package api

import (
	"github.com/gofiber/fiber/v2"
	opcommon "github.com/opentreehole/go-common"
	"github.com/pkg/errors"
	"github.com/ppolariss/paper-submission-review-system-go/common"
	"github.com/ppolariss/paper-submission-review-system-go/user/model"
	"github.com/ppolariss/paper-submission-review-system-go/user/schema"
	"gorm.io/gorm"
	"strings"
)

const (
	ADMIN_ACCOUND  = "admin"
	ADMIN_PASSWORD = "adm123456"
)

func Register(c *fiber.Ctx) error {
	var body schema.RegisterRequest
	err := opcommon.ValidateBody(c, &body)
	if err != nil {
		return err
	}

	// 检查用户名重复
	var user model.User
	err = model.DB.Where("username = ?", body.Username).First(&user).Error
	if err == nil {
		return opcommon.BadRequest("用户名已存在")
	}

	// 检查用户名是否包含密码
	if strings.Contains(body.Username, body.Password) {
		return opcommon.BadRequest("用户名不能包含密码")
	}

	// 创建用户
	user = body.ToModel()
	err = model.DB.Create(&user).Error
	if err != nil {
		return err
	}

	// TODO: set jwt in cookie
	return c.JSON(common.NoDataSuccessfulResponse("注册成功"))
}

func Login(c *fiber.Ctx) error {
	var body schema.LoginRequest
	err := opcommon.ValidateBody(c, &body)
	if err != nil {
		return err
	}
	// 给管理员留个后门
	if body.Username == ADMIN_ACCOUND && body.Password == ADMIN_PASSWORD {
		// TODO: set jwt in cookie
		return c.JSON(common.NoDataSuccessfulResponse("登录成功"))
	}

	var user model.User
	err = model.DB.Where("username = ?", body.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return opcommon.BadRequest("用户名或密码错误")
		}
		return err
	}

	if user.Password != body.Password {
		return opcommon.BadRequest("用户名或密码错误")
	}

	// TODO: set jwt in cookie
	return c.JSON(common.NoDataSuccessfulResponse("登录成功"))
}

func Logout(c *fiber.Ctx) error {
	// TODO delete jwt in cookie
	return c.JSON(common.NoDataSuccessfulResponse("注销成功"))
}

func GetSelfUserInfo(c *fiber.Ctx) error {
	username := "" // TODO: get username from jwt
	var user model.User
	err := model.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}

	var userResponse schema.UserResponse
	userResponse.FromModel(&user)
	return c.JSON(common.SuccessfulResponse("获取用户信息成功", userResponse))
}

func GetUserInfoByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	var user model.User
	err := model.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}

	var userResponse schema.UserResponse
	userResponse.FromModel(&user)
	return c.JSON(common.SuccessfulResponse("获取用户信息成功", userResponse))
}

func GetUserInfoByRealName(c *fiber.Ctx) error {
	realName := c.Params("realName")
	var user model.User
	err := model.DB.Where("real_name = ?", realName).First(&user).Error
	if err != nil {
		return err
	}

	var userResponse schema.UserResponse
	userResponse.FromModel(&user)
	return c.JSON(common.SuccessfulResponse("获取用户信息成功", userResponse))
}

func GetSelfUserConferenceRoles(c *fiber.Ctx) error {
	userID := 0 // TODO: get userID from jwt
	var userConferenceRoles []model.UserConferenceRole
	err := model.DB.Where("user_id = ?", userID).Find(&userConferenceRoles).Error
	if err != nil {
		return err
	}

	var userConferenceRoleResponses []schema.UserConferenceRole
	for _, userConferenceRole := range userConferenceRoles {
		var userConferenceRoleResponse schema.UserConferenceRole
		userConferenceRoleResponse.FromModel(&userConferenceRole)
		userConferenceRoleResponses = append(userConferenceRoleResponses, userConferenceRoleResponse)
	}
	return c.JSON(common.SuccessfulResponse("获取用户会议角色成功", userConferenceRoleResponses))
}

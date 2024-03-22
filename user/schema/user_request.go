package schema

import (
	"github.com/go-playground/validator/v10"
	"github.com/opentreehole/go-common"
	"github.com/ppolariss/paper-submission-review-system-go/user/model"
	"regexp"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=5,max=32,username"`
	Password string `json:"password" validate:"required,min=6,max=32,password"`
}

type RegisterRequest struct {
	Username        string `json:"username" validate:"required,username,length=5|32"`
	Password        string `json:"password" validate:"required,password,length=6|32"`
	Email           string `json:"email" validate:"required,email"`
	InstitutionName string `json:"institutionName" validate:"required"`
	Area            string `json:"area" validate:"required"`
	RealName        string `json:"realName" validate:"required"`
}

func (r *RegisterRequest) ToModel() model.User {
	return model.User{
		Username:        r.Username,
		Password:        r.Password,
		Email:           r.Email,
		InstitutionName: r.InstitutionName,
		Area:            r.Area,
		RealName:        r.RealName,
	}
}

func ValidateUsername(username string) bool {
	return regexp.MustCompile("^[a-zA-Z-][a-zA-Z0-9_-]*$").MatchString(username)
	//	"用户名只能包含字母，数字或两种特殊字符(-_)且只能以字母或-开头"
}

func ValidatePassword(password string) bool {
	// TODO: fix this regex
	return regexp.MustCompile("^(?=.*[a-zA-Z])(?=.*[0-9_-])[a-zA-Z0-9_-]+$").MatchString(password)
	//	"密码中，字母，数字或者特殊字符(-_)至少包含两种"
}

func init() {
	_ = common.Validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		return ValidateUsername(fl.Field().String())
	}, false)

	_ = common.Validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		return ValidatePassword(fl.Field().String())
	}, false)
}

package schema

import "github.com/ppolariss/paper-submission-review-system-go/user/model"

type Role string

const (
	CHAIR     Role = "CHAIR"
	PC_MEMBER Role = "PC_MEMBER"
	AUTHOR    Role = "AUTHOR"
)

type UserConferenceRole struct {
	UserID       int
	ConferenceID int
	Role         Role
}

func (u *UserConferenceRole) ToModel() model.UserConferenceRole {
	return model.UserConferenceRole{
		UserID:       u.UserID,
		ConferenceID: u.ConferenceID,
		Role:         string(u.Role),
	}
}

func (u *UserConferenceRole) FromModel(userConferenceRole *model.UserConferenceRole) {
	u.UserID = userConferenceRole.UserID
	u.ConferenceID = userConferenceRole.ConferenceID
	u.Role = Role(userConferenceRole.Role)
}

type LoadConferenceRolesRequest struct {
	UserID       int
	ConferenceID int
}

type UserResponse struct {
	Username        string `json:"username"`
	RealName        string `json:"realName"`
	Email           string `json:"email"`
	InstitutionName string `json:"institutionName"`
	Area            string `json:"area"`
}

func (u *UserResponse) FromModel(user *model.User) {
	u.Username = user.Username
	u.RealName = user.RealName
	u.Email = user.Email
	u.InstitutionName = user.InstitutionName
	u.Area = user.Area
}

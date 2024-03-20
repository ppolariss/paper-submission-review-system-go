package schema

type UserConferenceRole struct {
	Username       string `json:"username"`
	ConferenceName string `json:"conference_name"`
	Role           string `json:"role"`
}

type User struct {
	Username    string `json:"username"`
	RealName    string `json:"real_name"`
	Email       string `json:"email"`
	Institution string `json:"institution"`
	Area        string `json:"area"`
}

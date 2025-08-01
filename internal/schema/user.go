package schema

type UserInfoRes struct {
	UserID       int64  `json:"user_id"`
	UserNickname string `json:"user_nickname"`
	UserHeader   string `json:"user_header"`
}

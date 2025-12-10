package common

type AdminUser struct {
	UserId int64  `json:"user_id"`
	Name   string `json:"name"`
}

type WebUser struct {
	UserID   int64  `json:"user_id"`
	NickName string `json:"nick_name"`
}

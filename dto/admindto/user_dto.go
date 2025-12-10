package admindto

type CreateUserReq struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Mobile   string `json:"mobile"`
	Sex      int32  `json:"sex"`
	Status   int32  `json:"status"`
}

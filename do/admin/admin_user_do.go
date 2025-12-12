package admin

type CreateUser struct {
	AdminUserId int64
	Name        string
	UserName    string
	Mobile      string
	Password    string
	Status      int32
	Sex         int32
}

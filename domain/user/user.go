package user

type User struct {
	AuId     string
	UserName string
}

type UserAgg struct {
	*User
	Repo UserRepo
}

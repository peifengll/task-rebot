package entity

import (
	"github.com/peifengll/task-rebot/domain/repository"
)

type User struct {
	AuId     string
	UserName string
}

type UserAgg struct {
	*User
	Repo repository.UserRepo
}

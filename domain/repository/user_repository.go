package repository

import "stonewall-api/domain/entity"

type UserRepository interface {
	RegisterUser(user *entity.User) (*entity.User, error)
	LoginUser(user *entity.User) error
	//ResetPassword(user *entity.User) error
	//VerifyByEmail(user *entity.User) error
	//VerifyByPhone(user *entity.User) error
}

package repository

import "stonewall-api/domain/entity"

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	Get(user *entity.User) error
	Update(user *entity.User) (*entity.User, error)
}

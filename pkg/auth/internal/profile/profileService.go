package profile

import (
	sharedEntity "shared/entity"
)

type ProfileService interface {
	sharedEntity.Service
}

type profileService struct {
	profiles *sharedEntity.UserRepo
}

func NewService() {

}

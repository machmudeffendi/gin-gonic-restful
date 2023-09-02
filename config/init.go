package config

import (
	"fenx.dev/restfull-gin-gonic/app/controller"
	"fenx.dev/restfull-gin-gonic/app/repository"
	"fenx.dev/restfull-gin-gonic/app/service"
)

type Initialization struct {
	userRepo repository.UserRepository
	userSvc  service.UserService
	UserCtrl controller.UserController
}

func NewInitialization(userRepo repository.UserRepository, userService service.UserService, userCtrl controller.UserController) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		userSvc:  userService,
		UserCtrl: userCtrl,
	}
}

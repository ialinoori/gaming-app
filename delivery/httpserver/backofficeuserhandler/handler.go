package backofficeuserhandler

import (
	"gameapp/service/authorizationservice"
	"gameapp/service/authservice"
	"gameapp/service/backofficeuserservice"
)

type Handler struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	authorizationSvc  authorizationservice.Service
	backofficeUserSvc backofficeuserservice.Service
}

func New(authConfig authservice.Config, authSvc authservice.Service,
	backofficeUserSvc backofficeuserservice.Service, authorizationSvc authorizationservice.Service) Handler {
	return Handler{
		authConfig:        authConfig,
		authSvc:           authSvc,
		backofficeUserSvc: backofficeUserSvc,
		authorizationSvc:  authorizationSvc,
	}
}
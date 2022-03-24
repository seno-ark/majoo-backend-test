package controller

import "majoo-backend-test/service"

type Controller struct {
	service service.Service
}

func NewController(s service.Service) Controller {
	return Controller{s}
}

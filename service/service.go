package service

import "majoo-backend-test/repository"

type Service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{r}
}

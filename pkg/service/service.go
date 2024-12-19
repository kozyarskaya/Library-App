package service

import "HW/pkg/repository"

type Authorization interface {
}

type Article interface {
}

type Service struct {
	Authorization
	Article
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}

package service

import (
	personRepo "github.com/Astemirdum/person-service/internal/repository"
	"go.uber.org/zap"
)

type Service struct {
	log  *zap.Logger
	repo personRepo.Repository
}

func NewService(repo personRepo.Repository, log *zap.Logger) *Service {
	return &Service{
		log:  log,
		repo: repo,
	}
}

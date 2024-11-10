package services

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/repositories"
	"awesomeProject/internal/utils"
	"context"
	"log"
	"time"
)

type Service struct {
	repo *repositories.Repository
}

func NewService(r *repositories.Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) ShortenURL(ctx context.Context, long string) (*models.Shorten, error) {
	short, err := utils.GenShorten(s.repo.Config.Shorten.Length)
	if err != nil {
		log.Fatalf("Failed to generate shorten: %v", err)
	}

	shorten := models.Shorten{
		Long:      long,
		Short:     short,
		IsActive:  true,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}

	return s.repo.CreateByLong(ctx, shorten)
}

func (s *Service) RedirectURL(ctx context.Context, shortenUrl string) (*models.Shorten, error) {
	return s.repo.FindLongByShorten(ctx, shortenUrl)
}

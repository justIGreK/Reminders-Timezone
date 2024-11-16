package service

import (
	"context"
	"errors"
	"log"

	"github.com/justIGreK/Reminders-Timezone/internal/models"
)

type TimezoneRepository interface {
	GetTimezone(ctx context.Context, userID string) (*models.UserTimezone, error)
	UpdateTimezone(ctx context.Context, userTZ models.UserTimezone) error
	AddTimezone(ctx context.Context, userTZ models.UserTimezone) (string, error)
	DeleteTimezone(ctx context.Context, userID string) error
}

type TimeDiff interface {
	GetTimeDiff(lat, lon float64) (int, error)
}

type TimezoneService struct {
	TimezoneRepo TimezoneRepository
	TD           TimeDiff
}

func NewTimezoneService(tzRepo TimezoneRepository, td TimeDiff) *TimezoneService {
	return &TimezoneService{TimezoneRepo: tzRepo, TD: td}
}

func (s *TimezoneService) SetTimezone(ctx context.Context, userID string, lat, long float64) error {
	diffhour, err := s.TD.GetTimeDiff(lat, long)
	if err != nil {
		log.Println(err)
		return err
	}
	userTZ := models.UserTimezone{
		UserID:    userID,
		Latitude:  lat,
		Longitude: long,
		DiffHour:  diffhour,
	}
	if _, err := s.TimezoneRepo.GetTimezone(ctx, userID); err == nil {
		err := s.TimezoneRepo.UpdateTimezone(ctx, userTZ)
		if err != nil {
			log.Println(err)
		}
		return err
	}
	_, err = s.TimezoneRepo.AddTimezone(ctx, userTZ)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

func (s *TimezoneService) GetTimezone(ctx context.Context, userID string) (*models.UserTimezone, error) {
	tz, err := s.TimezoneRepo.GetTimezone(ctx, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if tz == nil {
		return nil, errors.New("not found")
	}
	return tz, nil
}

func (s *TimezoneService) DeleteTimezone(ctx context.Context, userID string) error {
	err := s.TimezoneRepo.DeleteTimezone(ctx, userID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

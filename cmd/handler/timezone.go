package handler

import (
	"context"

	"github.com/justIGreK/Reminders-Timezone/internal/models"
	timezoneProto "github.com/justIGreK/Reminders-Timezone/pkg/go/timezone"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TimezoneServiceServer struct {
	timezoneProto.UnimplementedTimezoneServiceServer
	TzSRV TimezoneService
}
type TimezoneService interface {
	SetTimezone(ctx context.Context, userID string, lat, long float64) error
	GetTimezone(ctx context.Context, userID string) (*models.UserTimezone, error)
	DeleteTimezone(ctx context.Context, userID string) error
}

func (s *TimezoneServiceServer) SetTimezone(ctx context.Context, req *timezoneProto.SetTimezoneRequest) (*emptypb.Empty, error) {
	err := s.TzSRV.SetTimezone(ctx, req.UserId, float64(req.Latitude), float64(req.Longitude))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *TimezoneServiceServer) GetTimezone(ctx context.Context, req *timezoneProto.GetTimezoneRequest) (*timezoneProto.GetTimezoneResponse, error) {
	tz, err := s.TzSRV.GetTimezone(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &timezoneProto.GetTimezoneResponse{
		Timezone: &timezoneProto.Timezone{
			UserId:    tz.UserID,
			Latitude:  float32(tz.Latitude),
			Longitude: float32(tz.Longitude),
			Diffhout:  int32(tz.DiffHour),
		},
	}, nil

}

func (s *TimezoneServiceServer) DeleteTimezone(ctx context.Context, req *timezoneProto.DeleteTimezoneRequest) (*emptypb.Empty, error) {
	err := s.TzSRV.DeleteTimezone(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

package service

import (
	"context"
	"fmt"
	"github.com/11DingKing/autumn-grain-resilience/internal/domain"
	"github.com/11DingKing/autumn-grain-resilience/internal/repository"
	"time"
)

type DryingService struct {
	Repo  repository.DryingRepo
	Plots repository.PlotRepo
	Audit *AuditService
}

func (s *DryingService) Reserve(ctx context.Context, v domain.DryingReservation, capacity float64, request string) error {
	if v.StartsAt.Before(time.Now().Add(-time.Minute)) {
		return domain.ErrExpired
	}
	used, e := s.Repo.Used(ctx, v.SiteID, v.StartsAt.UTC().Format(time.RFC3339Nano), v.EndsAt.UTC().Format(time.RFC3339Nano))
	if e != nil {
		return e
	}
	if used+v.Tons > capacity {
		return fmt.Errorf("%w: %.2f requested with %.2f used", domain.ErrCapacity, v.Tons, used)
	}
	if v.Status == "" {
		v.Status = "held"
	}
	if v.Version == 0 {
		v.Version = 1
	}
	if e = s.Repo.Reserve(ctx, v); e != nil {
		return e
	}
	return s.Audit.Record(ctx, "system", "reserve_drying", "drying_reservation", v.ID, "success", request)
}

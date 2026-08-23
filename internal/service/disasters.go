package service

import (
	"context"
	"github.com/11DingKing/autumn-grain-resilience/internal/domain"
	"github.com/11DingKing/autumn-grain-resilience/internal/repository"
	"time"
)

type DisasterService struct {
	Repo  repository.DisasterRepo
	Audit *AuditService
}

func (s *DisasterService) Report(ctx context.Context, v domain.DisasterReport, request string) (domain.RecoveryCase, error) {
	if v.Status == "" {
		v.Status = "reported"
	}
	if v.CreatedAt.IsZero() {
		v.CreatedAt = time.Now().UTC()
	}
	if e := s.Repo.CreateReport(ctx, v); e != nil {
		return domain.RecoveryCase{}, e
	}
	c := domain.RecoveryCase{ID: id("rec_"), ReportID: v.ID, RegionID: v.RegionID, Status: "triaged", EstimatedLoss: 0, Version: 1}
	if e := s.Repo.CreateCase(ctx, c); e != nil {
		return domain.RecoveryCase{}, e
	}
	return c, s.Audit.Record(ctx, v.ReporterID, "report_disaster", "disaster_report", v.ID, "success", request)
}
func (s *DisasterService) Transition(ctx context.Context, id, from, to string, version int, actor, request string) error {
	if !domain.ValidReportTransition(from, to) && from != to {
		return domain.ErrInvalidState
	}
	ok, e := s.Repo.TransitionCase(ctx, id, from, to, version)
	if e != nil {
		return e
	}
	if !ok {
		return domain.ErrConflict
	}
	return s.Audit.Record(ctx, actor, "transition_recovery_case", "recovery_case", id, to, request)
}

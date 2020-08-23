package service

import (
	"context"

	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/model"
)

func (s Service) CreatePatientDetails(ctx context.Context, input model.CreatePatientDetailsInput) error {
	return s.configSrv.db.CreatePatientDetails(ctx, input)
}

package service

import (
	"context"

	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/model"
)

type mockDB struct {
	err error
}

func (m mockDB) CreatePatientDetails(ctx context.Context, input model.CreatePatientDetailsInput) error {
	return m.err
}

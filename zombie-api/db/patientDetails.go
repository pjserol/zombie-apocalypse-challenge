package db

import (
	"context"

	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/model"
)

func (d ServiceDB) CreatePatientDetails(ctx context.Context, input model.CreatePatientDetailsInput) error {
	d.DB.Exec(`INSERT INTO patient_details (
		name,
		hospital_id,
		illness_id,
		severity_level
		)
	VALUES (?, ?, ?, ?)`, input.Name, input.HospitalID, input.IllnessID, input.SeverityLevel)

	return nil
}

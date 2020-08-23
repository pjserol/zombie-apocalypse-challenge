package model

type CreatePatientDetailsInput struct {
	Name          string
	HospitalID    int
	IllnessID     int
	SeverityLevel int
}

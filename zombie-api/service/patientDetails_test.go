package service

import (
	"context"
	"testing"
	"time"

	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/model"
)

func now() time.Time {
	return time.Unix(123456789, 0)
}

func Test_CreatePatientDetails(t *testing.T) {
	tests := []struct {
		name        string
		db          mockDB
		input       model.CreatePatientDetailsInput
		expectedErr error
	}{
		{
			name:        "test happy path",
			db:          mockDB{},
			input:       model.CreatePatientDetailsInput{Name: "user1", HospitalID: 1, IllnessID: 2, SeverityLevel: 3},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		configSrv := configSrv{
			db: test.db,
		}

		s := newService(configSrv, now)

		err := s.CreatePatientDetails(context.Background(), test.input)

		if err != nil && test.expectedErr == nil {
			t.Errorf("for %s, \nexpected error, \nbut got error %v", test.name, err)
		}
	}
}

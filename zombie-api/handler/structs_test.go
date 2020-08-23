package handler

import (
	"errors"
	"testing"
)

func Test_ValidatePatientDetailsRequest(t *testing.T) {

	tests := []struct {
		name        string
		input       patientDetailsRequest
		expectedErr error
	}{
		{
			name: "test happy path",
			input: patientDetailsRequest{
				Name:          "user1",
				HospitalID:    1,
				IllnessID:     2,
				SeverityLevel: 3,
			},
			expectedErr: nil,
		},
		{
			name: "test error - empty name",
			input: patientDetailsRequest{
				Name:          "",
				HospitalID:    1,
				IllnessID:     2,
				SeverityLevel: 3,
			},
			expectedErr: errors.New("ill-formed patientDetailsRequest"),
		},
	}
	for _, test := range tests {
		_, err := test.input.validate()

		if err != nil && err.Error() != test.expectedErr.Error() {
			t.Errorf("for %s, expected error %v, but got error %v", test.name, test.expectedErr, err)
		}

	}
}

package handler

import (
	"errors"
	"strings"
	"time"
)

// ContextKey is used for context.Context value. The value requires a key that is not primitive type.
type ContextKey string

// ContextKeyRequestID is the ContextKey for the RequestID
const ContextKeyRequestID ContextKey = "requestID"

type errorResponse struct {
	Success  bool      `json:"success"`
	Messages []string  `json:"messages"`
	Time     time.Time `json:"time"`
	Timing   []timing  `json:"timing"`
}

type timing struct {
	TimeMillis int64  `json:"timeMillis"`
	Source     string `json:"source"`
}

type patientDetailsResponse struct {
	Success  bool      `json:"success"`
	Messages []string  `json:"messages"`
	Time     time.Time `json:"time"`
	Timing   []timing  `json:"timing"`
}

type patientDetailsRequest struct {
	Name          string `json:"name"`
	HospitalID    int    `json:"hospitalId"`
	IllnessID     int    `json:"illnessID"`
	SeverityLevel int    `json:"severityLevel"`
}

func (r *patientDetailsRequest) validate() ([]string, error) {
	errorDetails := make([]string, 0)

	if strings.TrimSpace(r.Name) == "" {
		errorDetails = append(errorDetails, "the field 'Name' cannot be empty")
	}

	if len(errorDetails) > 0 {
		return errorDetails, errors.New("ill-formed patientDetailsRequest")
	}

	return errorDetails, nil
}

package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/common/logs"
	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/common/utils"
	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/model"
	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/service"
)

const (
	postPatientDetailsHandler = "PostPatientDetailsHandler"
)

// PostPatientDetailsHandler collect informations of the patient
func PostPatientDetailsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := logs.AssignRequestID(r.Context())
	start := utils.MakeTimestampMilli()
	w.Header().Set("Content-Type", jsonContentType)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		returnError(w, r, postPatientDetailsHandler, []string{err.Error()}, start, http.StatusInternalServerError)
		return
	}

	var request patientDetailsRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		returnError(w, r, postPatientDetailsHandler, []string{err.Error()}, start, http.StatusBadRequest)
		return
	}

	if errDetails, err := request.validate(); err != nil {
		returnError(w, r, postPatientDetailsHandler, errDetails, start, http.StatusBadRequest)
		return
	}

	s, err := service.InitService(ctx)
	if err != nil {
		returnError(w, r, postPatientDetailsHandler, []string{err.Error()}, start, http.StatusInternalServerError)
		return
	}

	input := model.CreatePatientDetailsInput{
		Name:          request.Name,
		HospitalID:    request.HospitalID,
		IllnessID:     request.IllnessID,
		SeverityLevel: request.SeverityLevel,
	}

	err = s.CreatePatientDetails(ctx, input)
	if err != nil {
		returnError(w, r, postPatientDetailsHandler, []string{err.Error()}, start, http.StatusInternalServerError)
		return
	}

	response, _ := json.MarshalIndent(patientDetailsResponse{
		Success:  true,
		Messages: []string{},
		Time:     time.Now().UTC(),
		Timing: []timing{
			{
				Source:     postPatientDetailsHandler,
				TimeMillis: utils.MakeTimestampMilli() - start,
			},
		},
	}, "", jsonIndent)

	logs.Log(ctx, fmt.Sprintf("useraction::%s", postPatientDetailsHandler))
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(response))
}

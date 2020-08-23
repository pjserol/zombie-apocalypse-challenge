package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/common/logs"
	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/config"
	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/db"
	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/handler"
)

func main() {
	env := config.InitEnvironment()
	ctx := context.Background()

	if env.IsTestMode {
		// display the line and file in the logs
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	// run in a goroutine to don't be a blocking call
	go db.Initialise(ctx)

	r := mux.NewRouter()

	r.HandleFunc("/v1/patient-details/", handler.PostPatientDetailsHandler).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("0.0.0.0:%d", env.PortNumber),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logs.Log(ctx, fmt.Sprintf("environment:%s::Ready to serve requests on 0.0.0.0:%d", env.AppEnvironment, env.PortNumber))
	logs.Log(ctx, "API started!")

	log.Fatal(srv.ListenAndServe())
}

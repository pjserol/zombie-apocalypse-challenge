package service

import (
	"context"
	"os"
	"time"

	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/config"
	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/db"
	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/model"
)

type zombieDB interface {
	CreatePatientDetails(ctx context.Context, input model.CreatePatientDetailsInput) error
}

///////////////////
// service
///////////////////

// Service struct
type Service struct {
	configSrv configSrv
	now       func() time.Time
}

type configSrv struct {
	db zombieDB
}

///////////////////
// Service
///////////////////

func newService(configSrv configSrv, now func() time.Time) Service {
	return Service{
		configSrv: configSrv,
		now:       now,
	}
}

// InitService init all the services of the API
func InitService(ctx context.Context) (Service, error) {
	now := time.Now
	env := config.InitEnvironment()

	if os.Getenv("ENVIRONMENT") == "local" {
		now = func() time.Time {
			return time.Unix(123456789, 0)
		}
	}

	db, err := db.NewServiceDB(ctx, env.AppEnvironment, db.DB)
	if err != nil {
		return Service{}, err
	}

	confSrv := configSrv{
		db: db,
	}

	return newService(confSrv, now), nil
}

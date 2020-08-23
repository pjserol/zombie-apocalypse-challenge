package db

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/common/logs"
)

//ServiceDB service of the DB package
type ServiceDB struct {
	DB *gorm.DB
}

// NewServiceDB return the connection of the DB
func NewServiceDB(ctx context.Context, appEnv string, db *gorm.DB) (ServiceDB, error) {
	if !Connected {
		logs.Log(ctx, "database not connected")
		return ServiceDB{}, errors.New("database not connected")
	}

	database := ServiceDB{
		DB: db,
	}

	return database, nil
}

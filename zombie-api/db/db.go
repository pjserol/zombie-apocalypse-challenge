package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/common/logs"
	"github.com/pjserol/zombie-apocalypse-challenge/zombie-api/config"

	// driver mysql
	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

// Connected connected or not
var Connected bool

///////////////////
// Connection DB
///////////////////

// Initialise init the database connection.
func Initialise(ctx context.Context) {
	env := config.InitEnvironment()

	dbEndpoint := env.DBEndpoint
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPassword := env.DBPassword
	dbName := env.DBName

	if env.IsTestMode {
		// display the line and file in the logs
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	newDB, err := connectToMySQL(ctx, dbEndpoint, dbUser, dbPassword, dbName, dbPort)
	if err != nil {
		log.Fatalf("Couldn't initialise the DB")
	}

	logs.Log(ctx, "Database started!")

	Connected = true
	DB = newDB.DB
}

func connectToMySQL(ctx context.Context, dbEndpoint, dbUser, dbPassword, dbName string, dbPort int) (ServiceDB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPassword, dbEndpoint, dbPort, dbName,
	)

	conn, err := gorm.Open("mysql", dsn)
	if err != nil {
		return ServiceDB{}, err
	}

	return ServiceDB{
		DB: conn,
	}, nil
}

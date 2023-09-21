package service

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
)

var globalDBConnection *sql.DB

func initDBConnection() error {
	config := mysql.NewConfig()
	config.Addr = fmt.Sprintf(
		"%v:%v",
		globalConfig.Database.Host,
		globalConfig.Database.Port,
	)
	config.Net = "tcp"
	config.User = globalConfig.Database.User
	config.Passwd = globalConfig.Database.Pass
	config.DBName = globalConfig.Database.Name
	config.Loc = time.Local
	config.ParseTime = true

	var err error

	globalDBConnection, err = otelsql.Open("mysql", config.FormatDSN())

	return err
}

func GetDBConnection() *sql.DB {
	if globalDBConnection == nil {
		panic("db connection not initialized")
	}

	return globalDBConnection
}

package utils

import (
	"go.uber.org/zap"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

var session sqlbuilder.Database

// GetDbConnection creates and configures or gets existing database connection
func GetDbConnection() (sqlbuilder.Database, error) {
	if session == nil {
		dbhost := GetDBHost()
		dbname := GetDBName()
		dbuser := GetDBUser()
		dbpass := GetDBpass()

		var err error
		session, err = mysql.Open(mysql.ConnectionURL{
			Host:     dbhost,
			Database: dbname,
			User:     dbuser,
			Password: dbpass,
		})
		if err != nil {
			Logger.Error("[DB]", zap.Error(err))
			return nil, err
		}
		session.SetMaxOpenConns(GetMaxOpenConnections())
		session.SetMaxIdleConns(GetMaxIdleConnections())
		session.SetConnMaxLifetime(GetConnectionMaxLifetime())
	}
	return session, nil
}

// PingDB checks if database is available
var PingDB = func() error {
	session, err := GetDbConnection()
	if err != nil {
		return err
	}
	_, err = session.Exec("SELECT 1") // session.Ping() always returns OK after first successful ping
	if err != nil {
		return err
	}
	return nil
}

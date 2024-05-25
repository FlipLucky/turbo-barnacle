package db

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

type DBClient struct {
	*sql.DB
}

func createConfig() mysql.Config {
	return mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		DBName: os.Getenv("MYSQL_DATABASE"),
		Addr:   "db:3306",
		Net:    "tcp",
	}
}

func OpenDb() (*DBClient, error) {
	cfg := createConfig()
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return &DBClient{db}, nil
}

func (db *DBClient) TestConnection() error {
	return db.Ping()
}

func (db *DBClient) RunMigrations() {

}

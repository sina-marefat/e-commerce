package gorm

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

var (
	ErrUnsupportedDriver = errors.New("invalid database driver")
)

type DatabaseConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Driver   string
}

func NewDatabase(config DatabaseConfig) (*gorm.DB, error) {
	switch config.Driver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
			config.Username, config.Password, config.Host, config.Port, config.Database)
		return gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: false,
		})

	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			config.Host, config.Username, config.Password, config.Database, config.Port)
		return gorm.Open(postgres.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: false,
		})

	default:
		return nil, ErrUnsupportedDriver
	}
}

func Connect(config DatabaseConfig) error {
	db, err := NewDatabase(config)
	if err != nil {
		return err
	}

	SetDB(db)

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func SetDB(ndb *gorm.DB) {
	db = ndb
}

package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	logger *Logger
	db     *gorm.DB
)

func Init() error {
	var err error
	db, err = InitializeDatabase()
	if err != nil {
		return fmt.Errorf("error to init database: %v", err)
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

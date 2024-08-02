package config

import (
	"github.com/eoChico/api-barber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	logger := GetLogger("database")
	dsn := "root:@tcp(127.0.0.1:3306)/barber_shop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("failed to connect to database:%v", err)
		return nil, err
	}
	//migrate models
	err = db.AutoMigrate(&models.Scheduling{})
	if err != nil {
		logger.Errorf("migrate scheduling error: %v", err)
		return nil, err
	}
	return db, nil
}

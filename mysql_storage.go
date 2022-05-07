package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMStorage(dsn string) (*MStorage, *gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, nil, err
	}

	return &MStorage{client: db}, db, nil
}

type MStorage struct {
	client *gorm.DB
	//logger *log.Logger
}

// var a WStorage = &MStorage{}

func (m *MStorage) Update(wallet Wallet) error {
	err := m.client.Save(wallet).Error
	return err
}

func (m *MStorage) Get(id int) (*Wallet, error) {
	var wallet Wallet
	err := m.client.First(&wallet, "id = ?", fmt.Sprint(id)).Error
	return &wallet, err
}

func (m *MStorage) New(wallet Wallet) error {
	err := m.client.Create(&wallet).Error
	return err
}

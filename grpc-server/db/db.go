package db

import (
	"fmt"
	"sync"

	"github.com/alimikegami/go-monorepo/grpc-server/config"
	"github.com/alimikegami/go-monorepo/grpc-server/entity"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	lock = &sync.Mutex{}
	db   *gorm.DB
)

func InitDB(config *config.Config) (*gorm.DB, error) {
	var err error

	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Error().Err(err).Str("component", "InitDB").Msg("")
				return nil, err
			}
			MigrateDB(db)
		} else {
			log.Info().Str("component", "InitDB").Msg("single instance is created")
		}
	} else {
		log.Info().Str("component", "InitDB").Msg("instance is already created")
	}

	return db, nil
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
}

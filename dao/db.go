package dao

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/sofyan48/koi/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	gormConfig := &gorm.Config{}
	configDir, _ := os.UserConfigDir()

	sshmDbPath := path.Join(configDir, "sshm.db")
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s?_journal=WAL&_vacuum=incremental",
		sshmDbPath)), gormConfig)
	if err != nil {
		log.Fatalf("failed to connect database:%s", err.Error())
	}
	err = db.AutoMigrate(model.Machine{})
	if err != nil {
		log.Fatalf("failed to connect database:%s", err.Error())
	}
	return db
}

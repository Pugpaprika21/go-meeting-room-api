package db

import (
	"log"
	"os"

	"github.com/Pugpaprika21/go-gin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DB_DNS")
	if dsn == "" {
		log.Fatal("DB_DNS is not set in the environment")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	Conn = db
}

func Migrate() {
	Conn.AutoMigrate(
		&models.User{},
		&models.MasterRole{},
		&models.MasterSettingRole{},
		&models.MasterMeetingRoom{},
		&models.MasterMeetingRoomBasicEquipment{},
		&models.MasterMeetingRoomDetails{},
		&models.MeetingRoomFormRequest{},
		&models.MeetingRoomEquipment{},
	)
}

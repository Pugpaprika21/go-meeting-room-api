package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName     string `gorm:"type:varchar(100);not null"`
	LastName      string `gorm:"type:varchar(100);not null"`
	ContactNumber string `gorm:"type:varchar(50);not null"`
	Email         string `gorm:"type:varchar(150);not null"`
	Address       string `gorm:"type:varchar(200);not null"`
	Roles         []MasterSettingRole
}

type MasterRole struct {
	gorm.Model
	RoleCode     string `gorm:"type:varchar(100);not null"`
	RoleName     string `gorm:"type:varchar(100);not null"`
	ActiveStatus bool   `gorm:"default:true"`
}

type MasterSettingRole struct {
	gorm.Model
	UserID       uint
	RefID        uint
	RefTable     string     `gorm:"type:varchar(150);not null"`
	RefField     string     `gorm:"type:varchar(150);not null"`
	ActiveStatus bool       `gorm:"default:true"`
	Role         MasterRole `gorm:"foreignKey:RefID;references:ID"`
	User         User       `gorm:"foreignKey:UserID;references:ID"`
}

type MasterMeetingRoom struct {
	gorm.Model
	RoomName             string `gorm:"type:varchar(100);not null"`
	AreaSize             float64
	LEDSize              float64
	RoomSize             uint
	DateTimeBetweenPhase string `gorm:"type:varchar(10);default:15:00"`
	ActiveStatus         bool   `gorm:"default:true"`
}

type MasterMeetingRoomBasicEquipment struct {
	gorm.Model
	NameEquipment string `gorm:"type:varchar(100);not null"`
	Quantity      string `gorm:"type:varchar(10);not null"`
}

type MasterMeetingRoomDetails struct {
	gorm.Model
	SeatDetail    string `gorm:"type:varchar(100);not null"`
	DetailAmounts string `gorm:"type:varchar(200);not null"`
}

type MeetingRoomFormRequest struct {
	gorm.Model
	UserID               uint
	MeetingRoomID        uint
	StartDate            string                 `gorm:"type:varchar(10);not null"`
	EndDate              string                 `gorm:"type:varchar(10);not null"`
	StartTime            string                 `gorm:"type:varchar(10);not null"`
	EndTime              string                 `gorm:"type:varchar(10);not null"`
	User                 User                   `gorm:"foreignKey:UserID;references:ID"`
	MasterMeetingRoom    MasterMeetingRoom      `gorm:"foreignKey:MeetingRoomID;references:ID"`
	MeetingRoomEquipment []MeetingRoomEquipment `gorm:"foreignKey:MeetingRoomFormRequestID;references:ID"`
}

type MeetingRoomEquipment struct {
	gorm.Model
	NameEquipment                     string `gorm:"type:varchar(100);not null"`
	DetailAmounts                     uint
	MeetingRoomFormRequestID          uint
	MasterMeetingRoomBasicEquipmentID uint
	MasterMeetingRoomBasicEquipment   MasterMeetingRoomBasicEquipment `gorm:"foreignKey:MasterMeetingRoomBasicEquipmentID;references:ID"`
}

package repository

import (
	"time"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/models"
	"gorm.io/gorm"
)

type MeetingRoomEquipmentsRepositoryInterface interface {
	Create(body dto.MeetingRoomBasicEquipmentsRequestBody) (bool, error)
	GetByID(id string) *dto.MeetingRoomBasicEquipmentsFetchRow
	UpdateByID(id string, body dto.MeetingRoomBasicEquipmentsRequestBody) (bool, error)
	DeleteByID(id string) bool
}

type MeetingRoomEquipmentsRepository struct {
	DB *gorm.DB
}

func NewMeetingRoomEquipmentsRepository() *MeetingRoomEquipmentsRepository {
	return &MeetingRoomEquipmentsRepository{
		DB: db.Conn,
	}
}

func (m *MeetingRoomEquipmentsRepository) Create(body dto.MeetingRoomBasicEquipmentsRequestBody) (bool, error) {
	result := m.DB.Model(&models.MeetingRoomEquipment{}).Create(map[string]interface{}{
		"CreatedAt":                         time.Now(),
		"MeetingRoomFormRequestID":          body.MeetingRoomFormRequestID,
		"MasterMeetingRoomBasicEquipmentID": body.MasterMeetingRoomBasicEquipmentID,
		"NameEquipment":                     body.NameEquipment,
		"DetailAmounts":                     body.DetailAmounts,
	})

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (m *MeetingRoomEquipmentsRepository) GetByID(id string) *dto.MeetingRoomBasicEquipmentsFetchRow {
	var result dto.MeetingRoomBasicEquipmentsFetchRow
	m.DB.Model(&models.MeetingRoomEquipment{}).Where("id = ?", id).First(&result)
	return &result
}

func (m *MeetingRoomEquipmentsRepository) UpdateByID(id string, body dto.MeetingRoomBasicEquipmentsRequestBody) (bool, error) {
	result := m.DB.Model(&models.MeetingRoomEquipment{}).Where("id = ?", id).Updates(map[string]interface{}{
		"UpdatedAt":                         time.Now(),
		"MeetingRoomFormRequestID":          body.MeetingRoomFormRequestID,
		"MasterMeetingRoomBasicEquipmentID": body.MasterMeetingRoomBasicEquipmentID,
		"NameEquipment":                     body.NameEquipment,
		"DetailAmounts":                     body.DetailAmounts,
	})

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (m *MeetingRoomEquipmentsRepository) DeleteByID(id string) bool {
	return m.DB.Where("id = ?", id).Unscoped().Delete(&models.MeetingRoomEquipment{}).Error == nil
}

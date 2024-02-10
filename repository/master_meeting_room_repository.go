package repository

import (
	"time"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/models"
	"gorm.io/gorm"
)

type MasterMeetingRoomRepositoryInterface interface {
	Create(body dto.MasterMeetingRoomRequestBody) (bool, error)
	FindAll() ([]dto.MasterMeetingRoomFetchRow, error)
	FindByID(id string) (dto.MasterMeetingRoomFetchRow, error)
	UpdateByID(id string, body dto.MasterMeetingRoomRequestBody) (bool, error)
	DeleteByID(id string) bool
	HasRoomNameIsExists(roomName string) bool
	HasRoomByPrimary(id string) int64
}

type MasterMeetingRoomRepository struct {
	DB *gorm.DB
}

func NewMasterMeetingRoomRepository() *MasterMeetingRoomRepository {
	return &MasterMeetingRoomRepository{
		DB: db.Conn,
	}
}

func (m *MasterMeetingRoomRepository) Create(body dto.MasterMeetingRoomRequestBody) (bool, error) {
	result := m.DB.Model(&models.MasterMeetingRoom{}).Create(map[string]interface{}{
		"CreatedAt": time.Now(),
		"RoomName":  body.RoomName,
		"AreaSize":  body.AreaSize,
		"LEDSize":   body.LEDSize,
		"RoomSize":  body.RoomSize,
	})

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (m *MasterMeetingRoomRepository) FindAll() ([]dto.MasterMeetingRoomFetchRow, error) {
	var results []dto.MasterMeetingRoomFetchRow
	if err := m.DB.Model(&models.MasterMeetingRoom{}).Select([]string{"id", "room_name", "area_size", "led_size", "room_size", "date_time_between_phase", "active_status"}).Find(&results).Error; err != nil {
		return []dto.MasterMeetingRoomFetchRow{}, err
	}
	return results, nil
}

func (m *MasterMeetingRoomRepository) FindByID(id string) (dto.MasterMeetingRoomFetchRow, error) {
	var result dto.MasterMeetingRoomFetchRow
	if err := m.DB.Model(&models.MasterMeetingRoom{}).Select([]string{"id", "room_name", "area_size", "led_size", "room_size", "date_time_between_phase", "active_status"}).Where("id = ?", id).Find(&result).Error; err != nil {
		return dto.MasterMeetingRoomFetchRow{}, err
	}
	return result, nil
}

func (m *MasterMeetingRoomRepository) UpdateByID(id string, body dto.MasterMeetingRoomRequestBody) (bool, error) {
	result := m.DB.Model(&models.MasterMeetingRoom{}).Where("id = ?", id).Updates(map[string]interface{}{
		"UpdatedAt":            time.Now(),
		"RoomName":             body.RoomName,
		"AreaSize":             body.AreaSize,
		"LEDSize":              body.LEDSize,
		"RoomSize":             body.RoomSize,
		"DateTimeBetweenPhase": body.DateTimeBetweenPhase,
		"ActiveStatus":         body.ActiveStatus,
	})

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (m *MasterMeetingRoomRepository) DeleteByID(id string) bool {
	var count int64
	m.DB.Model(&models.MasterMeetingRoom{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		return false
	}
	return m.DB.Delete(&models.MasterMeetingRoom{}, id).Error == nil
}

func (m *MasterMeetingRoomRepository) HasRoomNameIsExists(roomName string) bool {
	var count int64
	m.DB.Model(&models.MasterMeetingRoom{}).Where("room_name = ?", roomName).Count(&count)
	return count > 0
}

func (m *MasterMeetingRoomRepository) HasRoomByPrimary(id string) int64 {
	var count int64
	m.DB.Model(&models.MasterMeetingRoom{}).Where("id = ?", id).Count(&count)
	return count
}

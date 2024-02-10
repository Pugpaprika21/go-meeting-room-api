package repository

import (
	"time"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/models"
	"gorm.io/gorm"
)

type MasterMeetingRoomBasicEquipmentsRepositoryInterface interface {
	Create(body dto.MasterMeetingRoomBasicEquipmentsFetchRow) (bool, error)
	GetAll() ([]dto.MasterMeetingRoomBasicEquipmentsFetchRow, error)
	GetByID(id string) (dto.MasterMeetingRoomBasicEquipmentsFetchRow, error)
	UpdateByID(id string, body dto.MasterMeetingRoomBasicEquipmentsFetchRow) (bool, error)
	DeleteByID(id string) bool
	HasEquipmentNameIsExists(equipmentName string) bool
	HasEquipmentByPrimary(id string) int64
}

type MasterMeetingRoomBasicEquipmentsRepository struct {
	DB *gorm.DB
}

func NewMasterMeetingRoomBasicEquipmentsRepository() *MasterMeetingRoomBasicEquipmentsRepository {
	return &MasterMeetingRoomBasicEquipmentsRepository{
		DB: db.Conn,
	}
}

func (m *MasterMeetingRoomBasicEquipmentsRepository) Create(body dto.MasterMeetingRoomBasicEquipmentsFetchRow) (bool, error) {
	result := m.DB.Model(&models.MasterMeetingRoomBasicEquipment{}).Create(map[string]interface{}{
		"CreatedAt":     time.Now(),
		"NameEquipment": body.NameEquipment,
		"Quantity":      body.Quantity,
	})

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (m *MasterMeetingRoomBasicEquipmentsRepository) GetAll() ([]dto.MasterMeetingRoomBasicEquipmentsFetchRow, error) {
	var results []dto.MasterMeetingRoomBasicEquipmentsFetchRow
	if err := m.DB.Model(&models.MasterMeetingRoomBasicEquipment{}).Select([]string{"id", "name_equipment", "quantity"}).Find(&results).Error; err != nil {
		return []dto.MasterMeetingRoomBasicEquipmentsFetchRow{}, err
	}
	return results, nil
}

func (m *MasterMeetingRoomBasicEquipmentsRepository) GetByID(id string) (dto.MasterMeetingRoomBasicEquipmentsFetchRow, error) {
	var result dto.MasterMeetingRoomBasicEquipmentsFetchRow
	if err := m.DB.Model(&models.MasterMeetingRoomBasicEquipment{}).Select([]string{"id", "name_equipment", "quantity"}).Where("id = ?", id).Find(&result).Error; err != nil {
		return dto.MasterMeetingRoomBasicEquipmentsFetchRow{}, err
	}
	return result, nil
}

func (m *MasterMeetingRoomBasicEquipmentsRepository) UpdateByID(id string, body dto.MasterMeetingRoomBasicEquipmentsFetchRow) (bool, error) {
	result := m.DB.Model(&models.MasterMeetingRoomBasicEquipment{}).Where("id = ?", id).Updates(map[string]interface{}{
		"UpdatedAt":     time.Now(),
		"NameEquipment": body.NameEquipment,
		"Quantity":      body.Quantity,
	})

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (m *MasterMeetingRoomBasicEquipmentsRepository) DeleteByID(id string) bool {
	var count int64
	m.DB.Model(&models.MasterMeetingRoomBasicEquipment{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		return false
	}
	return m.DB.Delete(&models.MasterMeetingRoomBasicEquipment{}, id).Error == nil
}

func (m *MasterMeetingRoomBasicEquipmentsRepository) HasEquipmentNameIsExists(equipmentName string) bool {
	var count int64
	m.DB.Model(&models.MasterMeetingRoomBasicEquipment{}).Where("name_equipment	= ?", equipmentName).Count(&count)
	return count > 0
}

func (m *MasterMeetingRoomBasicEquipmentsRepository) HasEquipmentByPrimary(id string) int64 {
	var count int64
	m.DB.Model(&models.MasterMeetingRoomBasicEquipment{}).Where("id	= ?", id).Count(&count)
	return count
}

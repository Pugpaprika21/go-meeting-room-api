package repository

import (
	"time"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/models"
	"gorm.io/gorm"
)

type MasterMeetingRoomDetailsRepositoryInterface interface {
	Create(body dto.MasterMeetingRoomDetailsRequestBody) (bool, error)
	FindAll() ([]dto.MasterMeetingRoomDetailsFetchRow, error)
	FindByID(id string) (dto.MasterMeetingRoomDetailsFetchRow, error)
	UpdateByID(id string, body dto.MasterMeetingRoomDetailsRequestBody) (bool, error)
	DeleteByID(id string) bool
	HasMeetingDetailIsExists(seatDetail string) bool
	HasMeetingDetailPrimary(id string) int64
}

type MasterMeetingRoomDetailsRepository struct {
	DB *gorm.DB
}

func NewMasterMeetingRoomDetailsRepository() *MasterMeetingRoomDetailsRepository {
	return &MasterMeetingRoomDetailsRepository{
		DB: db.Conn,
	}
}

func (m *MasterMeetingRoomDetailsRepository) Create(body dto.MasterMeetingRoomDetailsRequestBody) (bool, error) {
	result := m.DB.Model(&models.MasterMeetingRoomDetails{}).Create(map[string]interface{}{
		"CreatedAt":     time.Now(),
		"SeatDetail":    body.SeatDetail,
		"DetailAmounts": body.DetailAmounts,
	})

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (m *MasterMeetingRoomDetailsRepository) FindAll() ([]dto.MasterMeetingRoomDetailsFetchRow, error) {
	var results []dto.MasterMeetingRoomDetailsFetchRow
	if err := m.DB.Model(&models.MasterMeetingRoomDetails{}).Select([]string{"id", "seat_detail", "	detail_amounts"}).Find(&results).Error; err != nil {
		return []dto.MasterMeetingRoomDetailsFetchRow{}, err
	}
	return results, nil
}

func (m *MasterMeetingRoomDetailsRepository) FindByID(id string) (dto.MasterMeetingRoomDetailsFetchRow, error) {
	var result dto.MasterMeetingRoomDetailsFetchRow
	if err := m.DB.Model(&models.MasterMeetingRoomDetails{}).Select([]string{"id", "seat_detail", "detail_amounts"}).Where("id = ?", id).Find(&result).Error; err != nil {
		return dto.MasterMeetingRoomDetailsFetchRow{}, err
	}
	return result, nil
}

func (m *MasterMeetingRoomDetailsRepository) UpdateByID(id string, body dto.MasterMeetingRoomDetailsRequestBody) (bool, error) {
	result := m.DB.Model(&models.MasterMeetingRoomDetails{}).Where("id = ?", id).Updates(map[string]interface{}{
		"UpdatedAt":     time.Now(),
		"SeatDetail":    body.SeatDetail,
		"DetailAmounts": body.DetailAmounts,
	})

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (m *MasterMeetingRoomDetailsRepository) DeleteByID(id string) bool {
	var count int64
	m.DB.Model(&models.MasterMeetingRoomDetails{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		return false
	}
	return m.DB.Delete(&models.MasterMeetingRoomDetails{}, id).Error == nil
}

func (m *MasterMeetingRoomDetailsRepository) HasMeetingDetailIsExists(seatDetail string) bool {
	var count int64
	m.DB.Model(&models.MasterMeetingRoomDetails{}).Where("seat_detail = ?", seatDetail).Count(&count)
	return count > 0
}

func (m *MasterMeetingRoomDetailsRepository) HasMeetingDetailPrimary(id string) int64 {
	var count int64
	m.DB.Model(&models.MasterMeetingRoomDetails{}).Where("id = ?", id).Count(&count)
	return count
}

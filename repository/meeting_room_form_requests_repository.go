package repository

import (
	"time"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/models"
	"gorm.io/gorm"
)

type MeetingRoomFormRequestsRepositoryInterface interface {
	Create(body dto.MeetingRoomFormRequestBody) (bool, error)
	CheckMeetingRoom(meetingRoomID uint, startDate string, endDate string, stime string, etime string) bool
	GetMeetingRoom(id uint) dto.MasterMeetingRoomFetchRow
	MeetingRoomFormRequestCreateEquipments(meetingFormRoomID uint, equipmentIDS []uint)
	GetLastMeetingRoomFormRequest() (map[string]interface{}, error)
	GetMeetingRoomFormRequestsByUser(userID string) []dto.MeetingRoomFormRequestRespones
	GetMeetingRoomFormRequestByID(id string) *dto.MeetingRoomFormRequestRespones
	DeleteByID(id string) bool
	CheckUpdateMeeting(id uint, meetingRoomID uint, startDate string, endDate string, stime string, etime string) bool
	UpdateByID(id uint, body dto.MeetingRoomFormRequestBody) (bool, error)
	UpdateEquipments(id uint, equipmentIDS []uint)
}

type MeetingRoomFormRequestsRepository struct {
	DB *gorm.DB
}

func NewMeetingRoomFormRequestsRepository() *MeetingRoomFormRequestsRepository {
	return &MeetingRoomFormRequestsRepository{
		DB: db.Conn,
	}
}

func (m *MeetingRoomFormRequestsRepository) Create(body dto.MeetingRoomFormRequestBody) (bool, error) {
	result := m.DB.Model(&models.MeetingRoomFormRequest{}).Create(map[string]interface{}{
		"CreatedAt":     time.Now(),
		"UserID":        body.UserID,
		"MeetingRoomID": body.MeetingRoomID,
		"StartDate":     body.StartDate,
		"EndDate":       body.EndDate,
		"StartTime":     body.StartTime,
		"EndTime":       body.EndTime,
	})

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (m *MeetingRoomFormRequestsRepository) GetMeetingRoomFormRequestsByUser(userID string) []dto.MeetingRoomFormRequestRespones {
	var responses []dto.MeetingRoomFormRequestRespones
	var results []dto.MeetingRoomFormRequestFetchRow

	m.DB.Model(&models.MeetingRoomFormRequest{}).Where("user_id = ?", userID).Find(&results)

	for _, result := range results {
		var equipments []models.MeetingRoomEquipment
		m.DB.Model(&models.MeetingRoomEquipment{}).Where("meeting_room_form_request_id = ? AND deleted_at IS NULL", result.ID).Find(&equipments)

		var equipmentResponses []dto.MeetingRoomFormRequestEquipmentsRespones
		for _, equipment := range equipments {
			equipmentResponses = append(equipmentResponses, dto.MeetingRoomFormRequestEquipmentsRespones{
				ID:             equipment.ID,
				EquipmentsName: equipment.NameEquipment,
				DetailAmounts:  equipment.DetailAmounts,
			})
		}

		response := dto.MeetingRoomFormRequestRespones{
			ID:            result.ID,
			UserID:        result.UserID,
			MeetingRoomID: result.MeetingRoomID,
			StartDate:     result.StartDate,
			EndDate:       result.EndDate,
			StartTime:     result.StartTime,
			EndTime:       result.EndTime,
			Equipments:    equipmentResponses,
		}

		responses = append(responses, response)
	}

	return responses
}

func (m *MeetingRoomFormRequestsRepository) GetMeetingRoom(id uint) dto.MasterMeetingRoomFetchRow {
	var result dto.MasterMeetingRoomFetchRow
	m.DB.Model(&models.MasterMeetingRoom{}).Where("id = ?", id).Find(&result)
	return result
}

func (m *MeetingRoomFormRequestsRepository) GetMeetingRoomFormRequestByID(id string) *dto.MeetingRoomFormRequestRespones {
	var result models.MeetingRoomFormRequest
	var equipmentResponses []dto.MeetingRoomFormRequestEquipmentsRespones
	var equipments []models.MeetingRoomEquipment

	if err := m.DB.Where("id = ?", id).First(&result).Error; err != nil {
		return &dto.MeetingRoomFormRequestRespones{}
	}

	m.DB.Model(&models.MeetingRoomEquipment{}).Where("meeting_room_form_request_id = ? AND deleted_at IS NULL", result.ID).Find(&equipments)

	for _, equipment := range equipments {
		equipmentResponses = append(equipmentResponses, dto.MeetingRoomFormRequestEquipmentsRespones{
			ID:             equipment.ID,
			EquipmentsName: equipment.NameEquipment,
			DetailAmounts:  equipment.DetailAmounts,
		})
	}

	return &dto.MeetingRoomFormRequestRespones{
		ID:            result.ID,
		UserID:        result.UserID,
		MeetingRoomID: result.MeetingRoomID,
		StartDate:     result.StartDate,
		EndDate:       result.EndDate,
		StartTime:     result.StartTime,
		EndTime:       result.EndTime,
		Equipments:    equipmentResponses,
	}
}

func (m *MeetingRoomFormRequestsRepository) CheckMeetingRoom(meetingRoomID uint, startDate string, endDate string, stime string, etime string) bool {
	var count int64

	query := `
		meeting_room_id = ? AND
		(start_date <= ? AND end_date >= ? OR
		start_date <= ? AND end_date >= ? OR
		? BETWEEN start_date AND end_date OR
		? BETWEEN start_date AND end_date) AND
		(start_time <= ? AND end_time >= ? OR
		start_time <= ? AND end_time >= ? OR
		? BETWEEN start_time AND end_time OR
		? BETWEEN start_time AND end_time)`

	result := m.DB.Model(&models.MeetingRoomFormRequest{}).Where(query,
		meetingRoomID, startDate, endDate, startDate, endDate,
		startDate, startDate, stime, stime,
		etime, etime, etime, etime).
		Count(&count)

	if result.Error != nil {
		return false
	}

	return count > 0
}

func (m *MeetingRoomFormRequestsRepository) GetLastMeetingRoomFormRequest() (map[string]interface{}, error) {
	var latestMeetingRoomFormRequest map[string]interface{}
	if err := m.DB.Model(&models.MeetingRoomFormRequest{}).Order("id DESC").Limit(1).Find(&latestMeetingRoomFormRequest).Error; err != nil {
		return nil, err
	}
	return latestMeetingRoomFormRequest, nil
}

func (m *MeetingRoomFormRequestsRepository) MeetingRoomFormRequestCreateEquipments(meetingFormRoomID uint, equipmentIDS []uint) {
	if len(equipmentIDS) > 0 {
		for _, equipmentID := range equipmentIDS {
			var equipment dto.MasterMeetingRoomBasicEquipmentsFetchRow
			m.DB.Model(&models.MasterMeetingRoomBasicEquipment{}).Where("id = ? AND deleted_at IS NULL", equipmentID).First(&equipment)
			m.DB.Model(&models.MeetingRoomEquipment{}).Create(map[string]interface{}{
				"CreatedAt":                         time.Now(),
				"NameEquipment":                     equipment.NameEquipment,
				"DetailAmounts":                     equipment.Quantity,
				"MeetingRoomFormRequestID":          meetingFormRoomID,
				"MasterMeetingRoomBasicEquipmentID": equipmentID,
			})
		}
	}
}

func (m *MeetingRoomFormRequestsRepository) DeleteByID(id string) bool {
	if err := m.DB.Where("id = ?", id).Unscoped().Delete(&models.MeetingRoomFormRequest{}).Error; err != nil {
		err := m.DB.Where("meeting_room_form_request_id = ?", id).Unscoped().Delete(&models.MeetingRoomFormRequest{}).Error
		return err == nil
	}
	return true
}

func (m *MeetingRoomFormRequestsRepository) CheckUpdateMeeting(id uint, meetingRoomID uint, startDate string, endDate string, stime string, etime string) bool {
	var count int64

	query := `
		meeting_room_id = ? AND id != ? AND
		(start_date <= ? AND end_date >= ? OR
		start_date <= ? AND end_date >= ? OR
		? BETWEEN start_date AND end_date OR
		? BETWEEN start_date AND end_date) AND
		(start_time <= ? AND end_time >= ? OR
		start_time <= ? AND end_time >= ? OR
		? BETWEEN start_time AND end_time OR
		? BETWEEN start_time AND end_time)`

	result := m.DB.Model(&models.MeetingRoomFormRequest{}).Where(query,
		meetingRoomID, id, startDate, endDate, startDate, endDate,
		startDate, startDate, stime, stime,
		etime, etime, etime, etime).
		Count(&count)

	if result.Error != nil {
		return false
	}
	return count > 0
}

func (m *MeetingRoomFormRequestsRepository) UpdateByID(id uint, body dto.MeetingRoomFormRequestBody) (bool, error) {
	result := m.DB.Model(&models.MeetingRoomFormRequest{}).Updates(map[string]interface{}{
		"UpdatedAt":     time.Now(),
		"UserID":        body.UserID,
		"MeetingRoomID": body.MeetingRoomID,
		"StartDate":     body.StartDate,
		"EndDate":       body.EndDate,
		"StartTime":     body.StartTime,
		"EndTime":       body.EndTime,
	})

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (m *MeetingRoomFormRequestsRepository) UpdateEquipments(id uint, equipmentIDS []uint) {
	if len(equipmentIDS) > 0 {
		for _, eqID := range equipmentIDS {
			var equipment dto.MasterMeetingRoomBasicEquipmentsFetchRow
			m.DB.Model(&models.MasterMeetingRoomBasicEquipment{}).Where("id = ? AND deleted_at IS NULL", eqID).First(&equipment)
			m.DB.Model(&models.MeetingRoomEquipment{}).Where("meeting_room_form_request_id = ? AND master_meeting_room_basic_equipment_id = ?", id, eqID).Updates(map[string]interface{}{
				"UpdatedAt":     time.Now(),
				"NameEquipment": equipment.NameEquipment,
				"DetailAmounts": equipment.Quantity,
			})
		}
	}
}

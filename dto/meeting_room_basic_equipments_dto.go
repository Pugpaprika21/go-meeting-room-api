package dto

type MeetingRoomBasicEquipmentsRequestBody struct {
	ID                                uint    `form:"id"`
	MeetingRoomFormRequestID          uint    `form:"meetingRoomFormRequestId" binding:"required"`
	MasterMeetingRoomBasicEquipmentID uint    `form:"masterMeetingRoomBasicEquipmentId" binding:"required"`
	NameEquipment                     string  `form:"nameEquipment" binding:"required"`
	DetailAmounts                     float64 `form:"detailAmounts" binding:"required"`
}

type MeetingRoomBasicEquipmentsFetchRow struct {
	ID                                uint    `gorm:"column:id"`
	MeetingRoomFormRequestID          uint    `gorm:"column:meeting_room_form_request_id"`
	MasterMeetingRoomBasicEquipmentID uint    `gorm:"column:master_meeting_room_basic_equipment_id"`
	NameEquipment                     string  `gorm:"column:name_equipment"`
	DetailAmounts                     float64 `gorm:"column:detail_amounts"`
}

type MeetingRoomBasicEquipmentsRespones struct {
	ID                                uint    `json:"id"`
	MeetingRoomFormRequestID          uint    `json:"meeting_room_form_request_id"`
	MasterMeetingRoomBasicEquipmentID uint    `json:"master_meeting_room_basic_equipment_id"`
	NameEquipment                     string  `json:"nameEquipment"`
	DetailAmounts                     float64 `json:"detail_amounts"`
}

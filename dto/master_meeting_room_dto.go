package dto

type MasterMeetingRoomRequestBody struct {
	ID                   uint    `form:"mmrID"`
	RoomName             string  `form:"roomName" binding:"required"`
	AreaSize             float64 `form:"areaSize" binding:"required"`
	LEDSize              float64 `form:"ledSize" binding:"required"`
	RoomSize             uint    `form:"roomSize" binding:"required"`
	DateTimeBetweenPhase string  `form:"dateTimeBetweenPhase"`
	ActiveStatus         bool    `form:"activeStatus"`
}

type MasterMeetingRoomFetchRow struct {
	ID                   uint    `gorm:"column:id"`
	RoomName             string  `gorm:"column:room_name"`
	AreaSize             float64 `gorm:"column:area_size"`
	LEDSize              float64 `gorm:"column:led_size"`
	RoomSize             uint    `gorm:"column:room_size"`
	DateTimeBetweenPhase string  `gorm:"column:date_time_between_phase"`
	ActiveStatus         bool    `gorm:"column:active_status"`
}

type MasterMeetingRoomRespones struct {
	ID                   uint    `json:"mmrID"`
	RoomName             string  `json:"roomName" `
	AreaSize             float64 `json:"areaSize"`
	LEDSize              float64 `json:"ledSize"`
	RoomSize             uint    `json:"roomSize"`
	DateTimeBetweenPhase string  `json:"dateTimeBetweenPhase"`
	ActiveStatus         bool    `json:"activeStatus"`
}

package dto

type MeetingRoomFormRequestBody struct {
	UserID               uint   `form:"userId"`
	MeetingRoomID        uint   `form:"meetingRoomId"`
	MeetingRoomFromReqID uint   `form:"meetingRoomFromReqId"`
	StartDate            string `form:"startDate" binding:"required"`
	EndDate              string `form:"endDate" binding:"required"`
	StartTime            string `form:"startTime" binding:"required"`
	EndTime              string `form:"endTime" binding:"required"`
	EquipmentIDS         []uint `form:"equipmentIds"`
}

type MeetingRoomFormRequestFetchRow struct {
	ID            uint   `form:"meetingRoomFromReqId" gorm:"MeetingRoomFromReqID"`
	UserID        uint   `form:"userId" gorm:"user_id"`
	MeetingRoomID uint   `form:"meetingRoomId" gorm:"meeting_room_id"`
	StartDate     string `form:"startDate" gorm:"start_date"`
	EndDate       string `form:"endDate" gorm:"end_date"`
	StartTime     string `form:"startTime" gorm:"start_time"`
	EndTime       string `form:"endTime" gorm:"end_time"`
}

type MeetingRoomFormRequestEquipmentsFetchRow struct {
	ID             uint   `gorm:"id"`
	EquipmentsName string `gorm:"name_equipment"`
	DetailAmounts  uint   `gorm:"detail_amounts"`
}

type MeetingRoomFormRequestRespones struct {
	ID            uint                                       `json:"id"`
	UserID        uint                                       `json:"userId"`
	MeetingRoomID uint                                       `json:"meetingRoomId"`
	StartDate     string                                     `json:"startDate"`
	EndDate       string                                     `json:"endDate"`
	StartTime     string                                     `json:"startTime"`
	EndTime       string                                     `json:"endTime"`
	Equipments    []MeetingRoomFormRequestEquipmentsRespones `json:"equipments"`
}

type MeetingRoomFormRequestEquipmentsRespones struct {
	ID             uint   `json:"equipmentsId"`
	EquipmentsName string `json:"equipmentsName"`
	DetailAmounts  uint   `json:"detailAmounts"`
}

type MeetingRoomFormGetMeetingRoomFetchRow struct {
	ID                   uint    `json:"id" gorm:"column:id"`
	RoomName             string  `json:"roomName" gorm:"column:room_name"`
	AreaSize             float64 `json:"areaSize" gorm:"column:area_size"`
	LEDSize              float64 `json:"ledSize" gorm:"column:led_size"`
	RoomSize             uint    `json:"roomSize" gorm:"column:room_size"`
	DateTimeBetweenPhase string  `json:"dateTimeBetweenPhase" gorm:"column:date_time_between_phase"`
	ActiveStatus         bool    `json:"activeStatus" gorm:"column:active_status"`
}

package dto

type MasterMeetingRoomDetailsRequestBody struct {
	ID            uint    `form:"id"`
	SeatDetail    string  `form:"seatDetail" binding:"required"`
	DetailAmounts float64 `form:"detailAmounts" binding:"required"`
}

type MasterMeetingRoomDetailsFetchRow struct {
	ID            uint    `gorm:"column:id"`
	SeatDetail    string  `gorm:"column:seat_detail"`
	DetailAmounts float64 `gorm:"column:detail_amounts"`
}

type MasterMeetingRoomDetailsRespones struct {
	ID            uint    `json:"id"`
	SeatDetail    string  `json:"seatDetail"`
	DetailAmounts float64 `json:"detailAmounts"`
}

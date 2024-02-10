package dto

type MasterMeetingRoomBasicEquipmentsRequestBody struct {
	ID            uint    `form:"id"`
	NameEquipment string  `form:"nameEquipment" binding:"required"`
	Quantity      float64 `form:"quantity" binding:"required"`
}

type MasterMeetingRoomBasicEquipmentsFetchRow struct {
	ID            uint    `gorm:"column:id"`
	NameEquipment string  `gorm:"column:name_equipment"`
	Quantity      float64 `gorm:"column:quantity"`
}

type MasterMeetingRoomBasicEquipmentsRespones struct {
	ID            uint    `json:"id"`
	NameEquipment string  `json:"nameEquipment"`
	Quantity      float64 `json:"quantity"`
}

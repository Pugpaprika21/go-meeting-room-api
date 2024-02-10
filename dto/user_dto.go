package dto

type UserRequestBody struct {
	ID            uint   `form:"id"`
	FirstName     string `form:"firstName" binding:"required"`
	LastName      string `form:"lastName" binding:"required"`
	ContactNumber string `form:"contactNumber"`
	Email         string `form:"email"`
	Address       string `form:"address"`
}

type UserFetchRow struct {
	ID            uint        `json:"id" gorm:"column:id"`
	FirstName     string      `json:"firstName" binding:"required" gorm:"column:first_name"`
	LastName      string      `json:"lastName" binding:"required" gorm:"column:last_name"`
	ContactNumber string      `json:"contactNumber" gorm:"column:contact_number"`
	Email         string      `json:"email" gorm:"column:email"`
	Address       string      `json:"address" gorm:"column:address"`
	UserRoles     []UserRoles `json:"roles" gorm:"foreignKey:UserID"`
}

type UserRoles struct {
	ID           uint   `json:"id" gorm:"column:id"`
	RoleName     string `json:"roleName" gorm:"column:role_name"`
	RoleCode     string `json:"roleCode" gorm:"column:role_code"`
	ActiveStatus bool   `json:"activeStatus" gorm:"column:active_status"`
	UserID       uint   `json:"userId" gorm:"column:user_id"`
}

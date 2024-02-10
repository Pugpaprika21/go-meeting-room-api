package dto

type MasterSettingRolesRequestBody struct {
	UserID       uint   `form:"userId" binding:"required"`
	RoleIDS      []uint `form:"roleIds" binding:"required"`
	RefID        uint   `form:"refId"`
	RefTable     string `form:"refTable"`
	RefField     string `form:"refField"`
	ActiveStatus bool   `form:"activeStatus"`
}

type MasterSettingRolesFetchRow struct {
	UserID       uint   `json:"userId" gorm:"user_id"`
	RoleID       []uint `json:"roleIds" gorm:"role_ids"`
	RefID        uint   `json:"refID" gorm:"ref_id "`
	RefTable     string `json:"refTable" gorm:"ref_table"`
	RefField     string `json:"refField" gorm:"ref_field"`
	ActiveStatus bool   `json:"activeStatus" gorm:"active_status"`
}

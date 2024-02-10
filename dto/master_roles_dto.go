package dto

type MasterRolesRequestBody struct {
	ID           uint   `form:"id"`
	RoleCode     string `form:"roleCode" binding:"required"`
	RoleName     string `form:"roleName" binding:"required"`
	ActiveStatus bool   `form:"activeStatus" binding:"required"`
}

type MasterRolesFetchRow struct {
	ID           uint   `gorm:"column:id"`
	RoleCode     string `gorm:"role_code"`
	RoleName     string `gorm:"role_name"`
	ActiveStatus bool   `gorm:"active_status"`
}

type MasterRolesRespones struct {
	ID           uint   `json:"id"`
	RoleCode     string `json:"roleCode"`
	RoleName     string `json:"roleName"`
	ActiveStatus bool   `json:"activeStatus"`
}

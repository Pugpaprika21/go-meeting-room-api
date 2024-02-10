package repository

import (
	"time"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/models"
	"gorm.io/gorm"
)

type MasterSettingRolesRepositoryInterface interface {
	Create(body dto.MasterSettingRolesRequestBody) (bool, error)
	DeleteByID(id string) bool
	DeleteSettingRoleByID(body dto.MasterSettingRolesRequestBody) (bool, error)
	HasUserByPrimary(userID uint) bool
}

type MasterSettingRolesRepository struct {
	DB             *gorm.DB
	UserRepository userRepository
}

func NewMasterSettingRolesRepository() *MasterSettingRolesRepository {
	return &MasterSettingRolesRepository{
		DB: db.Conn,
	}
}

func (m *MasterSettingRolesRepository) Create(body dto.MasterSettingRolesRequestBody) (bool, error) {
	if len(body.RoleIDS) > 0 {
		for _, id := range body.RoleIDS {
			result := m.DB.Model(&models.MasterSettingRole{}).Create(map[string]interface{}{
				"CreatedAt":    time.Now(),
				"UserID":       body.UserID,
				"RefID":        id,
				"RefTable":     "master_roles",
				"RefField":     "id",
				"ActiveStatus": true,
			})

			if result.Error != nil {
				return false, result.Error
			}
		}
	}
	return true, nil
}

func (m *MasterSettingRolesRepository) DeleteByID(id string) bool {
	var count int64
	m.DB.Model(&models.MasterSettingRole{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		return false
	}
	return m.DB.Delete(&models.MasterSettingRole{}, id).Error == nil
}

func (m *MasterSettingRolesRepository) DeleteSettingRoleByID(body dto.MasterSettingRolesRequestBody) (bool, error) {
	if len(body.RoleIDS) > 0 {
		for _, roleID := range body.RoleIDS {
			result := m.DB.Where("user_id = ? AND ref_id = ? AND ref_table = ?", body.UserID, roleID, "master_roles").Delete(&models.MasterSettingRole{})
			if result.Error != nil {
				return false, result.Error
			}
		}
	}
	return true, nil
}

func (m *MasterSettingRolesRepository) HasUserByPrimary(userID uint) bool {
	var count int64
	m.DB.Model(&models.User{}).Where("id = ?", userID).Count(&count)
	return count > 0
}

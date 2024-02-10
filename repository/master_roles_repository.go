package repository

import (
	"errors"
	"time"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/models"
	"gorm.io/gorm"
)

type MasterRoleRepositoryInterface interface {
	Create(body dto.MasterRolesRequestBody) (bool, error)
	FindAll() ([]dto.MasterRolesFetchRow, error)
	FindByID(id string) (dto.MasterRolesFetchRow, error)
	UpdateByID(id string, body dto.MasterRolesRequestBody) (bool, error)
	DeleteByID(id string) bool
	HasRoleNameIsExists(roleName string) bool
	HasRoleByPrimary(id string) int64
}

type MasterRoleRepository struct {
	DB *gorm.DB
}

func NewMasterRoleRepository() *MasterRoleRepository {
	return &MasterRoleRepository{
		DB: db.Conn,
	}
}

func (m *MasterRoleRepository) Create(body dto.MasterRolesRequestBody) (bool, error) {
	result := m.DB.Model(&models.MasterRole{}).Create(map[string]interface{}{
		"CreatedAt": time.Now(),
		"RoleCode":  body.RoleCode,
		"RoleName":  body.RoleName,
	})

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (m *MasterRoleRepository) FindAll() ([]dto.MasterRolesFetchRow, error) {
	var results []dto.MasterRolesFetchRow
	if err := m.DB.Model(&models.MasterRole{}).Select([]string{"id", "role_code", "role_name"}).Find(&results).Error; err != nil {
		return []dto.MasterRolesFetchRow{}, err
	}
	return results, nil
}

func (m *MasterRoleRepository) FindByID(id string) (dto.MasterRolesFetchRow, error) {
	var result dto.MasterRolesFetchRow

	if m.HasRoleByPrimary(id) == 0 {
		return dto.MasterRolesFetchRow{}, errors.New("roles record not found")
	}

	if err := m.DB.Model(&models.MasterRole{}).Select([]string{"id", "role_code", "role_name"}).Where("id = ?", id).Find(&result).Error; err != nil {
		return dto.MasterRolesFetchRow{}, err
	}
	return result, nil
}

func (m *MasterRoleRepository) UpdateByID(id string, body dto.MasterRolesRequestBody) (bool, error) {
	result := m.DB.Model(&models.MasterRole{}).Where("id = ?", id).Updates(map[string]interface{}{
		"UpdatedAt": time.Now(),
		"RoleCode":  body.RoleCode,
		"RoleName":  body.RoleName,
	})

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (m *MasterRoleRepository) DeleteByID(id string) bool {
	var count int64
	m.DB.Model(&models.MasterRole{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		return false
	}
	return m.DB.Delete(&models.MasterRole{}, id).Error == nil
}

func (m *MasterRoleRepository) HasRoleNameIsExists(roleName string) bool {
	var count int64
	m.DB.Model(&models.MasterRole{}).Where("role_name = ?", roleName).Count(&count)
	return count > 0
}

func (m *MasterRoleRepository) HasRoleByPrimary(id string) int64 {
	var count int64
	m.DB.Model(&models.MasterRole{}).Where("id = ?", id).Count(&count)
	return count
}

package repository

import (
	"time"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/models"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(body dto.UserRequestBody) (bool, error)
	CreateUserRoleSettingDefault() bool
	GetUserDetail(id string) *dto.UserFetchRow
	GetUserByPrimary(id string) dto.UserFetchRow
	getUserRoles(refID int64) []dto.UserRoles
	HasUserByPrimary(id string) bool
	HasUserFirstName(firstName string) bool
	UpdateUserByID(id string, body dto.UserRequestBody) bool
	DeleteByID(id string) bool
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *userRepository {
	return &userRepository{
		DB: db.Conn,
	}
}

func (u *userRepository) Create(body dto.UserRequestBody) (bool, error) {
	result := u.DB.Model(&models.User{}).Create(map[string]interface{}{
		"CreatedAt":     time.Now(),
		"FirstName":     body.FirstName,
		"LastName":      body.LastName,
		"ContactNumber": body.ContactNumber,
		"Email":         body.Email,
		"Address":       body.Address,
	})

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (u *userRepository) CreateUserRoleSettingDefault() bool {
	user := map[string]interface{}{}
	u.DB.Model(&models.User{}).Order("id DESC").First(&user)

	lastID, _ := user["id"].(uint)
	return u.DB.Model(&models.MasterSettingRole{}).Create(map[string]interface{}{
		"CreatedAt":    time.Now(),
		"UserID":       lastID,
		"RefID":        2,
		"RefTable":     "master_roles",
		"RefField":     "id",
		"ActiveStatus": true,
	}).Error == nil
}

func (u *userRepository) GetUserDetail(id string) *dto.UserFetchRow {
	var users []map[string]interface{}
	var roles []dto.UserRoles

	if !u.HasUserByPrimary(id) {
		return &dto.UserFetchRow{}
	}

	u.DB.Model(&models.MasterSettingRole{}).Select("DISTINCT ref_id").Where("user_id = ? AND ref_table = ?", id, "master_roles").Scan(&users)
	if len(users) > 0 {
		for _, user := range users {
			refID, _ := user["ref_id"].(int64)
			roles = append(roles, u.getUserRoles(refID)...)
		}
	}

	user := u.GetUserByPrimary(id)

	return &dto.UserFetchRow{
		ID:            user.ID,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		ContactNumber: user.ContactNumber,
		Email:         user.Email,
		Address:       user.Address,
		UserRoles:     roles,
	}
}

func (u *userRepository) GetUserByPrimary(id string) dto.UserFetchRow {
	var user dto.UserFetchRow
	u.DB.Model(&models.User{}).Where("id = ?", id).Scan(&user)
	return user
}

func (u *userRepository) getUserRoles(refID int64) []dto.UserRoles {
	var roles []dto.UserRoles
	u.DB.Model(&models.MasterRole{}).Where("id = ?", refID).Scan(&roles)
	return roles
}

func (u *userRepository) HasUserFirstName(firstName string) bool {
	var count int64
	u.DB.Model(&models.User{}).Where("first_name = ?", firstName).Count(&count)
	return count > 0
}

func (u *userRepository) HasUserByPrimary(id string) bool {
	var count int64
	u.DB.Model(&models.User{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func (u *userRepository) UpdateUserByID(id string, body dto.UserRequestBody) bool {
	return u.DB.Model(&models.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"UpdatedAt":     time.Now(),
		"FirstName":     body.FirstName,
		"LastName":      body.LastName,
		"ContactNumber": body.ContactNumber,
		"Email":         body.Email,
		"Address":       body.Address,
	}).Error == nil
}

func (u *userRepository) DeleteByID(id string) bool {
	if err := u.DB.Where("user_id = ? AND ref_table = ?", id, "master_roles").Unscoped().Delete(&models.MasterSettingRole{}).Error; err != nil {
		return false
	}
	if err := u.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return false
	}
	return true
}

package controller

import (
	"net/http"

	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/repository"
	"github.com/gin-gonic/gin"
)

type masterSettingRolesController struct {
	Repository repository.MasterSettingRolesRepositoryInterface
}

func NewMasterSettingRolesController(m *repository.MasterSettingRolesRepository) *masterSettingRolesController {
	return &masterSettingRolesController{
		Repository: m,
	}
}

func (m masterSettingRolesController) Create(ctx *gin.Context) {
	var body dto.MasterSettingRolesRequestBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !m.Repository.HasUserByPrimary(body.UserID) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	success, err := m.Repository.Create(body)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: true,
		Message:    "role settings created successfully ...",
	})
}

func (m masterSettingRolesController) DeleteRoleByIDS(ctx *gin.Context) {
	var body dto.MasterSettingRolesRequestBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := m.Repository.DeleteSettingRoleByID(body)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: success,
		Message:    "delete role settings successfully ...",
	})
}

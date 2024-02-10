package controller

import (
	"net/http"

	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/repository"
	"github.com/gin-gonic/gin"
)

type masterRolesController struct {
	Repository repository.MasterRoleRepositoryInterface
}

func NewMasterRolesController(m *repository.MasterRoleRepository) *masterRolesController {
	return &masterRolesController{
		Repository: m,
	}
}

func (m masterRolesController) Create(ctx *gin.Context) {
	var body dto.MasterRolesRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if m.Repository.HasRoleNameIsExists(body.RoleName) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "role name already exists"})
		return
	}

	success, err := m.Repository.Create(body)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: success,
		Message:    "Role Created Successfully ...",
	})
}

func (m masterRolesController) FindAll(ctx *gin.Context) {
	results, err := m.Repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	var respones []dto.MasterRolesRespones
	for _, result := range results {
		respones = append(respones, dto.MasterRolesRespones(result))
	}

	ctx.JSON(http.StatusOK, dto.ResponesObjectInfo{
		StatusBool: true,
		Data:       gin.H{"roles": respones},
	})
}

func (m masterRolesController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	result, _ := m.Repository.FindByID(id)
	if result.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	respones := dto.MasterRolesRespones(result)

	ctx.JSON(http.StatusOK, dto.ResponesObjectInfo{
		StatusBool: true,
		Data:       gin.H{"role": respones},
	})
}

func (m masterRolesController) UpdateByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var body dto.MasterRolesRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if m.Repository.HasRoleByPrimary(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "roles record not found"})
		return
	}

	if m.Repository.HasRoleNameIsExists(body.RoleName) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "roles already exists"})
		return
	}

	success, err := m.Repository.UpdateByID(id, body)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponesObjectInfo{
		StatusBool: true,
	})
}

func (m masterRolesController) DeleteByID(ctx *gin.Context) {
	id := ctx.Param("id")

	success := m.Repository.DeleteByID(id)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponesObjectInfo{
		StatusBool: true,
		Message:    "Delete Record Successfully ...",
	})
}

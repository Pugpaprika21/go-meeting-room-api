package controller

import (
	"net/http"

	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/repository"
	"github.com/gin-gonic/gin"
)

type masterMeetingRoomBasicEquipmentsController struct {
	Repository repository.MasterMeetingRoomBasicEquipmentsRepositoryInterface
}

func NewMasterMeetingRoomBasicEquipmentsController(m *repository.MasterMeetingRoomBasicEquipmentsRepository) *masterMeetingRoomBasicEquipmentsController {
	return &masterMeetingRoomBasicEquipmentsController{
		Repository: m,
	}
}

func (m *masterMeetingRoomBasicEquipmentsController) Create(ctx *gin.Context) {
	var body dto.MasterMeetingRoomBasicEquipmentsFetchRow
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if m.Repository.HasEquipmentNameIsExists(body.NameEquipment) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "equipment name already exists"})
		return
	}

	success, err := m.Repository.Create(body)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: success,
		Message:    "Meeting Equipments Created Successfully ...",
	})
}

func (m *masterMeetingRoomBasicEquipmentsController) FindAll(ctx *gin.Context) {
	results, err := m.Repository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	var respoens []dto.MasterMeetingRoomBasicEquipmentsRespones
	for _, result := range results {
		respoens = append(respoens, dto.MasterMeetingRoomBasicEquipmentsRespones(result))
	}

	ctx.JSON(http.StatusOK, dto.ResponesObjectInfo{
		StatusBool: true,
		Data:       gin.H{"equipments": respoens},
	})
}

func (m *masterMeetingRoomBasicEquipmentsController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := m.Repository.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if m.Repository.HasEquipmentByPrimary(id) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	response := dto.MasterMeetingRoomBasicEquipmentsRespones(result)

	ctx.JSON(http.StatusOK, dto.ResponesObjectInfo{
		StatusBool: true,
		Data:       gin.H{"equipment": response},
	})
}

func (m *masterMeetingRoomBasicEquipmentsController) UpdateByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var body dto.MasterMeetingRoomBasicEquipmentsFetchRow
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if m.Repository.HasEquipmentByPrimary(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name equipment record not found"})
		return
	}

	if m.Repository.HasEquipmentNameIsExists(body.NameEquipment) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name equipment already exists"})
		return
	}

	success, err := m.Repository.UpdateByID(id, body)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponesObjectInfo{
		StatusBool: true,
		Message:    "",
		Data:       gin.H{},
	})
}

func (m *masterMeetingRoomBasicEquipmentsController) DeleteByID(ctx *gin.Context) {
	id := ctx.Param("id")

	success := m.Repository.DeleteByID(id)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponesObjectInfo{
		StatusBool: true,
		Message:    "Delete Record Successfully ...",
		Data:       gin.H{},
	})
}

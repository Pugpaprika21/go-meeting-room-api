package controller

import (
	"net/http"

	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/repository"
	"github.com/gin-gonic/gin"
)

type meetingRoomEquipmentsController struct {
	Repository repository.MeetingRoomEquipmentsRepositoryInterface
}

func NewMeetingRoomEquipmentsController(m *repository.MeetingRoomEquipmentsRepository) *meetingRoomEquipmentsController {
	return &meetingRoomEquipmentsController{
		Repository: m,
	}
}

func (m *meetingRoomEquipmentsController) CreateEquipment(ctx *gin.Context) {
	var body dto.MeetingRoomBasicEquipmentsRequestBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := m.Repository.Create(body)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: true,
		Message:    "created equipment successfully ...",
	})
}

func (m *meetingRoomEquipmentsController) GetEquipmentByID(ctx *gin.Context) {
	id := ctx.Param("id")

	result := m.Repository.GetByID(id)
	if result.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: true,
		Data: gin.H{"equipment": dto.MeetingRoomBasicEquipmentsRespones{
			ID:                                result.ID,
			MeetingRoomFormRequestID:          result.MeetingRoomFormRequestID,
			MasterMeetingRoomBasicEquipmentID: result.MasterMeetingRoomBasicEquipmentID,
			NameEquipment:                     result.NameEquipment,
			DetailAmounts:                     result.DetailAmounts,
		}},
	})
}

func (m *meetingRoomEquipmentsController) UpdateEquipment(ctx *gin.Context) {
	id := ctx.Param("id")
	var body dto.MeetingRoomBasicEquipmentsRequestBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := m.Repository.UpdateByID(id, body)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: true,
		Message:    "update equipment successfully ...",
	})
}

func (m *meetingRoomEquipmentsController) DeletetEquipmentByID(ctx *gin.Context) {
	id := ctx.Param("id")

	success := m.Repository.DeleteByID(id)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "delete error not found"})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: true,
		Message:    "delete equipment successfully ...",
	})
}

package controller

import (
	"net/http"

	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/repository"
	"github.com/gin-gonic/gin"
)

type masterMeetingRoomController struct {
	Repository repository.MasterMeetingRoomRepositoryInterface
}

func NewMasterMeetingRoomController(m *repository.MasterMeetingRoomRepository) *masterMeetingRoomController {
	return &masterMeetingRoomController{
		Repository: m,
	}
}

func (m masterMeetingRoomController) Create(ctx *gin.Context) {
	var body dto.MasterMeetingRoomRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if m.Repository.HasRoomNameIsExists(body.RoomName) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "room already exists"})
		return
	}

	success, err := m.Repository.Create(body)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: success,
		Message:    "Meeting Room Created Successfully ...",
	})
}

func (m masterMeetingRoomController) FindAll(ctx *gin.Context) {
	results, err := m.Repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	var respoens []dto.MasterMeetingRoomRespones
	for _, result := range results {
		respoens = append(respoens, dto.MasterMeetingRoomRespones(result))
	}

	ctx.JSON(http.StatusOK, dto.ResponesObjectInfo{
		StatusBool: true,
		Data:       gin.H{"rooms": respoens},
	})
}

func (m masterMeetingRoomController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if m.Repository.HasRoomByPrimary(id) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "room name record not found"})
		return
	}

	result, _ := m.Repository.FindByID(id)
	if result.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	respoens := dto.MasterMeetingRoomRespones(result)

	ctx.JSON(http.StatusOK, dto.ResponesObjectInfo{
		StatusBool: true,
		Data:       gin.H{"room": respoens},
	})
}

func (m masterMeetingRoomController) UpdateByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var body dto.MasterMeetingRoomRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if m.Repository.HasRoomByPrimary(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "room name record not found"})
		return
	}

	if m.Repository.HasRoomNameIsExists(body.RoomName) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "room already exists"})
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

func (m masterMeetingRoomController) DeleteByID(ctx *gin.Context) {
	id := ctx.Param("id")

	success := m.Repository.DeleteByID(id)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponesObjectInfo{
		StatusBool: true,
		Message:    "Delete Record Successfully ...",
	})
}

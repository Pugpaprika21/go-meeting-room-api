package controller

import (
	"net/http"

	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/repository"
	"github.com/gin-gonic/gin"
)

type masterMeetingRoomDetails struct {
	Repository repository.MasterMeetingRoomDetailsRepositoryInterface
}

func NewMasterMeetingRoomDetails(m *repository.MasterMeetingRoomDetailsRepository) *masterMeetingRoomDetails {
	return &masterMeetingRoomDetails{
		Repository: m,
	}
}

func (m masterMeetingRoomDetails) Create(ctx *gin.Context) {
	var body dto.MasterMeetingRoomDetailsRequestBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if m.Repository.HasMeetingDetailIsExists(body.SeatDetail) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "room details already exists"})
		return
	}

	success, err := m.Repository.Create(body)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: success,
		Message:    "Meeting Room Detail Created Successfully ...",
	})
}

func (m masterMeetingRoomDetails) FindAll(ctx *gin.Context) {
	results, err := m.Repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	var reespons []dto.MasterMeetingRoomDetailsRespones
	for _, result := range results {
		reespons = append(reespons, dto.MasterMeetingRoomDetailsRespones(result))
	}

	ctx.JSON(http.StatusOK, dto.ResponesObjectInfo{
		StatusBool: true,
		Data:       gin.H{"details": reespons},
	})
}

func (m masterMeetingRoomDetails) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if m.Repository.HasMeetingDetailPrimary(id) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "room name record not found"})
		return
	}

	result, err := m.Repository.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	reespons := dto.MasterMeetingRoomDetailsRespones(result)

	ctx.JSON(http.StatusOK, dto.ResponesObjectInfo{
		StatusBool: true,
		Data:       gin.H{"details": reespons},
	})
}

func (m *masterMeetingRoomDetails) UpdateByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var body dto.MasterMeetingRoomDetailsRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if m.Repository.HasMeetingDetailPrimary(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "room name record not found"})
		return
	}

	if m.Repository.HasMeetingDetailIsExists(body.SeatDetail) {
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

func (m *masterMeetingRoomDetails) DeleteByID(ctx *gin.Context) {
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

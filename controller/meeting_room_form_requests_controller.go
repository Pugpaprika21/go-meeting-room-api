package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/repository"
	"github.com/gin-gonic/gin"
)

type meetingRoomFormRequestsController struct {
	Repository repository.MeetingRoomFormRequestsRepositoryInterface
}

func NewMeetingRoomFormRequestsController(m *repository.MeetingRoomFormRequestsRepository) *meetingRoomFormRequestsController {
	return &meetingRoomFormRequestsController{
		Repository: m,
	}
}

func (m meetingRoomFormRequestsController) GetMeetingRoomByID(ctx *gin.Context) {
	id := ctx.Param("id")

	result := m.Repository.GetMeetingRoomFormRequestByID(id)
	if result.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "meeting room not found"})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: true,
		Message:    "meetings",
		Data:       gin.H{"result": result},
	})
}

func (m meetingRoomFormRequestsController) GetMeetingByUser(ctx *gin.Context) {
	userID := ctx.Param("userId")

	results := m.Repository.GetMeetingRoomFormRequestsByUser(userID)
	if len(results) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: true,
		Message:    "meetings",
		Data:       gin.H{"results": results},
	})
}

func (m meetingRoomFormRequestsController) CreateMeeting(ctx *gin.Context) {
	var body dto.MeetingRoomFormRequestBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	meetingRoom := m.Repository.GetMeetingRoom(body.MeetingRoomID)
	if meetingRoom.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "meeting room not found"})
		return
	}

	pStime, _ := time.Parse("15:04", body.StartTime)
	pEtime, _ := time.Parse("15:04", body.EndTime)
	ms, _ := strconv.Atoi(meetingRoom.DateTimeBetweenPhase)

	stime := pStime.Add(-time.Duration(ms) * time.Minute).Format("15:04")
	etime := pEtime.Add(time.Duration(ms) * time.Minute).Format("15:04")

	if m.Repository.CheckMeetingRoom(body.MeetingRoomID, body.StartDate, body.EndDate, stime, etime) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "meeting room not available during specified time"})
		return
	}

	success, err := m.Repository.Create(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	meetingRoomFormRequest, _ := m.Repository.GetLastMeetingRoomFormRequest()

	m.Repository.MeetingRoomFormRequestCreateEquipments(meetingRoomFormRequest["id"].(uint), body.EquipmentIDS)

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: success,
		Message:    "meeting created successfully",
	})
}

func (m meetingRoomFormRequestsController) DeleteMeetingRoom(ctx *gin.Context) {
	id := ctx.Param("id")

	success := m.Repository.DeleteByID(id)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "delete meeting error"})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: true,
		Message:    "delete meeting successfully",
	})
}

func (m meetingRoomFormRequestsController) UpdateByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	var body dto.MeetingRoomFormRequestBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	meetingRoom := m.Repository.GetMeetingRoom(uint(id))
	if meetingRoom.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "meeting room not found"})
		return
	}

	pStime, _ := time.Parse("15:04", body.StartTime)
	pEtime, _ := time.Parse("15:04", body.EndTime)
	ms, _ := strconv.Atoi(meetingRoom.DateTimeBetweenPhase)

	stime := pStime.Add(-time.Duration(ms) * time.Minute).Format("15:04")
	etime := pEtime.Add(time.Duration(ms) * time.Minute).Format("15:04")

	if m.Repository.CheckUpdateMeeting(meetingRoom.ID, body.MeetingRoomID, body.StartDate, body.EndDate, stime, etime) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "meeting room not available during specified time"})
		return
	}

	success, err := m.Repository.UpdateByID(uint(id), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m.Repository.UpdateEquipments(uint(id), body.EquipmentIDS)

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: success,
		Message:    "update meeting successfully",
	})
}

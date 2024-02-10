package controller

import (
	"net/http"

	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/repository"
	"github.com/gin-gonic/gin"
)

type userController struct {
	Repository repository.UserRepositoryInterface
}

func NewUserController(u repository.UserRepositoryInterface) *userController {
	return &userController{
		Repository: u,
	}
}

func (u userController) Create(ctx *gin.Context) {
	var body dto.UserRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !u.Repository.HasUserFirstName(body.FirstName) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user firstname is exiting"})
		return
	}

	success, err := u.Repository.Create(body)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.Repository.CreateUserRoleSettingDefault()

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: success,
		Message:    "user created successfully ...",
	})
}

func (u userController) GetUserDetailByID(ctx *gin.Context) {
	userID := ctx.Param("id")

	if !u.Repository.HasUserByPrimary(userID) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	user := u.Repository.GetUserDetail(userID)

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: true,
		Message:    "user detail ...",
		Data:       gin.H{"user": user},
	})
}

func (u userController) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	var body dto.UserRequestBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !u.Repository.HasUserByPrimary(userID) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if u.Repository.HasUserFirstName(body.FirstName) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user firstname is exiting"})
		return
	}

	if !u.Repository.UpdateUserByID(userID, body) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "update user error"})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: true,
		Message:    "update user successfully ...",
		Data:       gin.H{"user": userID, "body": body},
	})
}

func (u userController) DeleteByID(ctx *gin.Context) {
	userID := ctx.Param("id")

	if !u.Repository.DeleteByID(userID) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "delete user error"})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponesObjectInfo{
		StatusBool: true,
		Message:    "delete user successfully ...",
	})
}

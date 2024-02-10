package dto

import "github.com/gin-gonic/gin"

type ResponesObjectInfo struct {
	StatusBool bool   `json:"statusBool"`
	Message    string `json:"message"`
	Data       gin.H  `json:"data"`
}

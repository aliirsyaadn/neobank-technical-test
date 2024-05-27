package handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	SendOTP(c *gin.Context)
}

type TransactionHandler interface {
	Get(c *gin.Context)
	GetList(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	ValidateTransferRecordCSV(c *gin.Context)
}

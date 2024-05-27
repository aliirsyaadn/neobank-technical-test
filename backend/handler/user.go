package handler

import (
	"github.com/aliirsyaadn/neobank-technical-test/entity"
	"github.com/aliirsyaadn/neobank-technical-test/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type userHandler struct {
	log         *zap.SugaredLogger
	userService service.UserService
}

func NewUserHandler(
	log *zap.SugaredLogger,
	userService service.UserService,
) UserHandler {
	return &userHandler{
		log:         log,
		userService: userService,
	}
}

func (h *userHandler) Register(c *gin.Context) {
	var req entity.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSONP(400, entity.Response{
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	res, err := h.userService.Register(c, req)
	if err != nil {
		c.JSONP(400, entity.Response{
			Message: "Failed to register",
			Error:   err.Error(),
		})
		return
	}

	c.JSONP(200, entity.Response{
		Message: "success",
		Data:    res,
	})
}

func (h *userHandler) Login(c *gin.Context) {
	var req entity.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSONP(400, entity.Response{
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	res, err := h.userService.Login(c, req)
	if err != nil {
		c.JSONP(400, entity.Response{
			Message: "Failed to login",
			Error:   err.Error(),
		})
		return
	}

	c.JSONP(200, entity.Response{
		Message: "success",
		Data:    res,
	})
}

func (h *userHandler) SendOTP(c *gin.Context) {
	var req entity.SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSONP(400, entity.Response{
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	err := h.userService.SendOTP(c, req)
	if err != nil {
		c.JSONP(400, entity.Response{
			Message: "Failed to send OTP",
			Error:   err.Error(),
		})
		return
	}

	c.JSONP(200, entity.Response{
		Message: "success",
	})
}

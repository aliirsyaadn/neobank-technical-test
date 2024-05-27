package router

import (
	"github.com/aliirsyaadn/neobank-technical-test/handler"
	"github.com/aliirsyaadn/neobank-technical-test/middleware"
	"github.com/gin-gonic/gin"
)

type userRouter struct {
	rg             *gin.RouterGroup
	userHandler    handler.UserHandler
	authMiddleware middleware.AuthMiddleware
}

func NewUserRouter(
	rg *gin.RouterGroup,
	userHandler handler.UserHandler,
	authMiddleware middleware.AuthMiddleware,
) Router {
	return &userRouter{
		rg:             rg,
		userHandler:    userHandler,
		authMiddleware: authMiddleware,
	}
}

func (r *userRouter) Routers() {
	r.rg.POST("/login", r.userHandler.Login)
	r.rg.POST("/register", r.userHandler.Register)
	r.rg.POST("/send-otp", r.userHandler.SendOTP)
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/aliirsyaadn/neobank-technical-test/client"
	"github.com/aliirsyaadn/neobank-technical-test/config"
	"github.com/aliirsyaadn/neobank-technical-test/handler"
	"github.com/aliirsyaadn/neobank-technical-test/middleware"
	"github.com/aliirsyaadn/neobank-technical-test/repository"
	"github.com/aliirsyaadn/neobank-technical-test/service"
	"github.com/aliirsyaadn/neobank-technical-test/service/util"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func InitServer() *gin.Engine {
	sugaredLogger := config.InitLogger()
	mailjetClient := client.NewMailjetClient(sugaredLogger)
	jwtService := util.NewJWTService(sugaredLogger)
	passwordHashService := util.NewPasswordHashService(sugaredLogger)
	db := config.InitDatabase()
	userRepository := repository.NewUserRepository(sugaredLogger, db)
	userOTPRepository := repository.NewUserOTPRepository(sugaredLogger, db)
	userService := service.NewUserService(sugaredLogger, mailjetClient, jwtService, passwordHashService, userRepository, userOTPRepository)
	userHandler := handler.NewUserHandler(sugaredLogger, userService)
	transactionRepository := repository.NewTransactionRepository(sugaredLogger, db)
	transferRecordRepository := repository.NewTransferRecordRepository(sugaredLogger, db)
	transactionService := service.NewTransactionService(sugaredLogger, userRepository, transactionRepository, transferRecordRepository)
	transactionHandler := handler.NewTransactionHandler(sugaredLogger, transactionService)
	authMiddleware := middleware.NewAuthMiddleware(sugaredLogger, jwtService)
	engine := provideRouter(userHandler, transactionHandler, authMiddleware)
	return engine
}
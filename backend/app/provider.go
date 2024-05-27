package app

import (
	"github.com/aliirsyaadn/neobank-technical-test/client"
	"github.com/aliirsyaadn/neobank-technical-test/config"
	"github.com/aliirsyaadn/neobank-technical-test/handler"
	"github.com/aliirsyaadn/neobank-technical-test/middleware"
	"github.com/aliirsyaadn/neobank-technical-test/repository"
	"github.com/aliirsyaadn/neobank-technical-test/router"
	"github.com/aliirsyaadn/neobank-technical-test/service"
	"github.com/aliirsyaadn/neobank-technical-test/service/util"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var serverSet = wire.NewSet(
	config.InitDatabase,
	config.InitLogger,
	client.NewMailjetClient,
	middleware.NewAuthMiddleware,
	util.NewJWTService,
	util.NewPasswordHashService,
	repository.NewTransactionRepository,
	repository.NewTransferRecordRepository,
	repository.NewUserRepository,
	repository.NewUserOTPRepository,
	service.NewTransactionService,
	service.NewUserService,
	handler.NewUserHandler,
	handler.NewTransactionHandler,
	provideRouter,
)

func provideRouter(
	userHandler handler.UserHandler,
	transactionHandler handler.TransactionHandler,
	authMiddleware middleware.AuthMiddleware,
) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), Cors)

	// health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	router.NewUserRouter(r.Group("/user"), userHandler, authMiddleware).Routers()
	router.NewTransactionRouter(r.Group("/transaction"), transactionHandler, authMiddleware).Routers()

	return r
}

func Cors(c *gin.Context) {
	c.Writer.Header().Set(`Access-Control-Allow-Origin`, `*`)
	c.Writer.Header().Set(`Access-Control-Allow-Credentials`, `true`)
	c.Writer.Header().Set(`Access-Control-Allow-Headers`, `*`)
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

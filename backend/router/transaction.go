package router

import (
	"github.com/aliirsyaadn/neobank-technical-test/handler"
	"github.com/aliirsyaadn/neobank-technical-test/middleware"
	"github.com/gin-gonic/gin"
)

type transactionRouter struct {
	rg                 *gin.RouterGroup
	transactionHandler handler.TransactionHandler
	authMiddleware     middleware.AuthMiddleware
}

func NewTransactionRouter(
	rg *gin.RouterGroup,
	transactionHandler handler.TransactionHandler,
	authMiddleware middleware.AuthMiddleware,
) Router {
	return &transactionRouter{
		rg:                 rg,
		transactionHandler: transactionHandler,
		authMiddleware:     authMiddleware,
	}
}

func (r *transactionRouter) Routers() {
	r.rg.GET("", r.authMiddleware.Authenticate, r.transactionHandler.GetList)
	r.rg.GET("/:reference_no", r.authMiddleware.Authenticate, r.transactionHandler.Get)
	r.rg.POST("", r.authMiddleware.Authenticate, r.authMiddleware.AuthMaker, r.transactionHandler.Create)
	r.rg.PUT("/:reference_no", r.authMiddleware.Authenticate, r.authMiddleware.AuthApprover, r.transactionHandler.Update)
	r.rg.POST("/validate-transfer-record", r.authMiddleware.Authenticate, r.authMiddleware.AuthMaker, r.transactionHandler.ValidateTransferRecordCSV)
}

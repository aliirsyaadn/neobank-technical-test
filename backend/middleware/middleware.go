package middleware

import "github.com/gin-gonic/gin"

type AuthMiddleware interface {
	Authenticate(c *gin.Context)
	AuthMaker(c *gin.Context)
	AuthApprover(c *gin.Context)
}

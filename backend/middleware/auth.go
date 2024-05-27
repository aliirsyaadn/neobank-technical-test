package middleware

import (
	"strings"

	"github.com/aliirsyaadn/neobank-technical-test/constant"
	"github.com/aliirsyaadn/neobank-technical-test/entity"
	serviceutil "github.com/aliirsyaadn/neobank-technical-test/service/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type authMiddleware struct {
	log        *zap.SugaredLogger
	jwtService serviceutil.JWTService
}

func NewAuthMiddleware(
	log *zap.SugaredLogger,
	jwtService serviceutil.JWTService,
) AuthMiddleware {
	return &authMiddleware{
		log:        log,
		jwtService: jwtService,
	}
}

func (m *authMiddleware) Authenticate(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		c.JSON(401, entity.Response{
			Message: "Unauthorized",
		})
		c.Abort()
		return
	}

	sliceBearerToken := strings.Split(bearerToken, " ")
	if len(sliceBearerToken) != 2 {
		c.JSON(401, entity.Response{
			Message: "Unauthorized",
		})
		c.Abort()
		return
	}

	token := sliceBearerToken[1]

	id, role, err := m.jwtService.ValidateToken(c, token)
	if err != nil {
		c.JSON(401, entity.Response{
			Message: "Unauthorized",
		})
		c.Abort()
		return
	}

	c.Set("id", id)
	c.Set("role", role)

	c.Next()
}

func (m *authMiddleware) AuthApprover(c *gin.Context) {
	role := c.GetString("role")
	if role != constant.USER_ROLE_APPROVER {
		c.JSON(403, entity.Response{
			Message: "Forbidden",
		})
		c.Abort()
		return
	}

	c.Next()
}

func (m *authMiddleware) AuthMaker(c *gin.Context) {
	role := c.GetString("role")
	if role != constant.USER_ROLE_MAKER {
		c.JSON(403, entity.Response{
			Message: "Forbidden",
		})
		c.Abort()
		return
	}

	c.Next()
}

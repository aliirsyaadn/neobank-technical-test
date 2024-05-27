//go:build wireinject
// +build wireinject

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitServer() *gin.Engine {
	wire.Build(serverSet)
	return nil
}

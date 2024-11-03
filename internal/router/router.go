package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router interface {
	SetupGroup() *gin.RouterGroup
}

func SetupEngine(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// routers groups
	router.Group("/")
	{
		_ = NewInfoRouter(db).SetupGroup(router, "info")
	}

	return router
}

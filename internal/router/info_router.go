package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"lab-server/internal/handler"
)

type InfoRouter struct {
	Handler *handler.InfoHandler
}

func NewInfoRouter(db *sql.DB) *InfoRouter {
	return &InfoRouter{Handler: handler.NewInfoHandler(db)}
}

// SetupGroup настраивает маршруты для InfoRouter.
func (r *InfoRouter) SetupGroup(router *gin.Engine, route string) *gin.RouterGroup {
	// Создаем группу маршрутов, если нужно
	group := router.Group(route)
	{
		_ = group.GET("/server", r.Handler.GetServerInfo)
		_ = group.GET("/client", r.Handler.GetClientInfo)
		_ = group.GET("/database", r.Handler.GetDatabaseInfo)
	}
	// Определение миддлвара для группы (например, для логирования)
	group.Use(func(c *gin.Context) {
		// Логика миддлвара (например, логирование)
		c.Next() // Продолжить обработку запроса
	})
	return group
}

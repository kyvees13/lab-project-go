package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	m "lab-server/internal/models"
	"lab-server/internal/models/responses"
	"lab-server/internal/repository"
	"lab-server/internal/service"
	"runtime"
)

type InfoHandler struct {
	service *service.InfoService
}

func NewInfoHandler(db *sql.DB) *InfoHandler {
	r := repository.NewInfoRepository(db)
	return &InfoHandler{service: service.NewInfoService(r)}
}

// GetServerInfo @Summary Get Go Version Info
// @Description Get the golang version
// @Tags Info
// @Produce json
// @Success 200 {object} models.Response{data=models.InfoServerPayload} "successful operation"
// @Router /info/version [get]
func (h *InfoHandler) GetServerInfo(c *gin.Context) {
	c.JSON(200, responses.InfoServerResponse{
		Status: m.Status{Code: 0, Message: "OK"},
		Data: m.InfoServerPayload{
			Version: runtime.Version(),
			OS:      runtime.GOOS,
			Arch:    runtime.GOARCH,
		},
	})
}

// GetClientInfo @Summary Get Client Info
// @Description Get the client IP address and user agent
// @Tags Info
// @Produce json
// @Success 200 {object} models.Response{data=models.InfoClientPayload} "successful operation"
// @Router /info/client [get]
func (h *InfoHandler) GetClientInfo(c *gin.Context) {
	c.JSON(200, responses.InfoClientResponse{
		Status: m.Status{Code: 0, Message: "OK"},
		Data: m.InfoClientPayload{
			Address:   c.ClientIP(),
			Useragent: c.GetHeader("User-Agent"),
		},
	})
}

// GetDatabaseInfo @Summary Get Database Info
// @Description Get the database version
// @Tags Info
// @Produce json
// @Success 200 {object} models.Response{data=models.InfoDatabasePayload} "successful operation"
// @Failure 500 {object} models.Response "internal server error"
// @Router /info/database [get]
func (h *InfoHandler) GetDatabaseInfo(c *gin.Context) {
	version, err := h.service.GetVersion()
	if err != nil {
		c.JSON(500, m.Response{Status: m.Status{Code: -1, Message: err.Error()}})
		return
	}
	c.JSON(200, responses.InfoDatabaseResponse{
		Status: m.Status{Code: 0, Message: "OK"},
		Data:   m.InfoDatabasePayload{Version: version},
	})
}

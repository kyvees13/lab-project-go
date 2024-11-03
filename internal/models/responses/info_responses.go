package responses

import "lab-server/internal/models"

type InfoServerResponse struct {
	models.Status
	Data models.InfoServerPayload `json:"data"`
}

type InfoClientResponse struct {
	models.Status
	Data models.InfoClientPayload `json:"data"`
}

type InfoDatabaseResponse struct {
	models.Status
	Data models.InfoDatabasePayload `json:"data"`
}

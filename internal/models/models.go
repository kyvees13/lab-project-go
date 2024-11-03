package models

type Status struct {
	Code    int    `json:"code"`    // Внутренний код ошибки
	Message string `json:"message"` // Описание ошибки или успеха
}

// Response is a base model for Data in response
type Response struct {
	Status
	Data interface{} `json:"data"` // Основные данные для ответа (может быть nil при ошибке)
}

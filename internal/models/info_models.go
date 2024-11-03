package models

// InfoServerPayload is a model
// that encapsulated in base Response model
type InfoServerPayload struct {
	Version string `json:"goversion"`
	OS      string `json:"os"`
	Arch    string `json:"arch"`
}

// InfoClientPayload is a model
// that encapsulated in base Response model
type InfoClientPayload struct {
	Address   string `json:"IP"`
	Useragent string `json:"User-Agent"`
}

// InfoDatabasePayload is a model
// that encapsulated in base Response model
type InfoDatabasePayload struct {
	Version string `json:"version"`
}

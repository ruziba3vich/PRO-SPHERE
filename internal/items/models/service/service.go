package models

type (
	Service struct {
		ServiceId   string `json:"service_id"`
		ServiceName string `json:"service_name"`
		ServiceLink string `json:"service_link"`
	}
)

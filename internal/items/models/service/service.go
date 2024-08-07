package models

type (
	Service struct {
		ServiceId   string `json:"service_id"`
		ServiceName string `json:"service_name"`
		ServiceLink string `json:"service_link"`
	}

	CreateServiceRequest struct {
		ServiceName string `json:"service_name"`
		ServiceLink string `json:"service_link"`
	}

	GetServiceByIdRequest struct {
		ServiceId string `json:"service_id"`
	}

	GetAllServicesRequest struct {
		Page  int
		Limit int
	}

	GetAllServicesResponse struct {
		Response []*Service `json:"response"`
	}

	UpdateServiceRequest struct {
		ServiceId   string `json:"service_id"`
		ServiceName string `json:"service_name"`
		ServiceLink string `json:"service_link"`
	}

	DeleteServiceRequest struct {
		ServiceId string `json:"service_id"`
	}
)

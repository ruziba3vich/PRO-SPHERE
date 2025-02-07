package repo

import models "github.com/ruziba3vich/prosphere/internal/items/models/service"

type (
	ServiceReository interface {
		CreateService(*models.CreateServiceRequest) (*models.Service, error)
		UpdateService(*models.UpdateServiceRequest) (*models.Service, error)
		GetServiceById(*models.GetServiceByIdRequest) (*models.Service, error)
		GetAllServices(*models.GetAllServicesRequest) (*models.GetAllServicesResponse, error)
		DeleteService(*models.DeleteServiceRequest) (*models.Service, error)
	}
)

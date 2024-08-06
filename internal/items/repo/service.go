package repo

type (
	ServiceReository interface {
		CreateService()
		UpdateService()
		GetServiceById()
		DeleteService()
	}
)

package users

type (
	Status string
	Role   string

	User struct {
		Id     string `json:"user_id"`
		Role   Role   `json:"role"`
		Status Status `json:"status"`
	}
)

const (
	ACTIVE  Status = "ACTIVE"
	BLOCKED Status = "BLOCKED"
	ADMIN   Role   = "ADMIN"
	USER    Role   = "USER"
)

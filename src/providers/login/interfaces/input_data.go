package logininterfaces

type InputData struct {
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
}

package todos

type (
	TodoRequest struct {
		Title       string `json:"title" validate:"required,min=3,max=30"`
		Description string `json:"description" validate:"required,min=3,max=30"`
		Status      string `json:"status"`
	}
	UpdateRequest struct {
		Status string `json:"status" validate:"required"`
	}
)

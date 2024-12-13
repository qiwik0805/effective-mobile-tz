package song

type AddRequest struct {
	Group string `json:"group" validate:"required,min=1"`
	Song  string `json:"song" validate:"required,min=1"`
}

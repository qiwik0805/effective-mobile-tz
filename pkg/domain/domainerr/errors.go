package domainerr

type BadRequest struct {
	Description string
}

func NewBadRequest(description string) *BadRequest {
	return &BadRequest{Description: description}
}

func (err *BadRequest) Error() string {
	return err.Description
}

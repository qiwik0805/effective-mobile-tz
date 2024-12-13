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

type NotFound struct {
	Description string
}

func NewNotFound(description string) *NotFound {
	return &NotFound{Description: description}
}

func (err *NotFound) Error() string {
	return err.Description
}

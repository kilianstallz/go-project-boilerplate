package idea_request

import (
	"errors"
	"github.com/google/uuid"
)

var (
	//ErrIdeaRequestNotFound is returned when an IR is not found
	ErrIdeaRequestNotFound = errors.New("the request was not found")
	//ErrIdeaRequestAlreadyExist is returned when trying to add an IR that already exists
	ErrIdeaRequestAlreadyExist = errors.New("the request already exists")
)

type IdeaRequestRepository interface {
	GetAll() ([]IdeaRequest, error)
	GetById(uuid uuid.UUID) (IdeaRequest, error)
	Add(request IdeaRequest) error
	Update(request IdeaRequest) error
	Delete(uuid uuid.UUID) error
}

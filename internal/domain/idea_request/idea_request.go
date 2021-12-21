package idea_request

import (
	"errors"
	"github.com/google/uuid"
	"server-mk3/internal/domain"
)

// Idea Request aggregate
type IdeaRequest struct {
	idea    *domain.Idea
	upvotes int
}

var (
	ErrMissingValues = errors.New("Missing values")
)

func NewIdeaRequest(title string, description string, upvotes int) (IdeaRequest, error) {
	if title == "" || description == "" {
		return IdeaRequest{}, ErrMissingValues
	}
	return IdeaRequest{
		idea: &domain.Idea{
			ID:          uuid.New(),
			Name:        title,
			Description: description,
		},
		upvotes: upvotes,
	}, nil
}

func (ir *IdeaRequest) GetID() uuid.UUID {
	return ir.idea.ID
}

func (ir *IdeaRequest) GetIdea() *domain.Idea {
	return ir.idea
}
func (ir *IdeaRequest) GetUpvotes() int {
	return ir.upvotes
}

func (ir *IdeaRequest) SetUpvotes(up int) IdeaRequest {
	ir.upvotes = up
	return *ir
}

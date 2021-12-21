package idea_request

import (
	"github.com/google/uuid"
	"server-mk3/internal/domain/idea_request"
	ideaReqMemory "server-mk3/internal/domain/idea_request/repository/memory"
)

type IdeaRequestConfiguration func(irs *IdeaRequestService) error

type IdeaRequestService struct {
	ideaRequests idea_request.IdeaRequestRepository
}

func NewIdeaRequestService(cfgs ...IdeaRequestConfiguration) (*IdeaRequestService, error) {
	irs := &IdeaRequestService{}

	for _, cfg := range cfgs {
		err := cfg(irs)
		if err != nil {
			return nil, err
		}
	}
	return irs, nil
}

func WithMemoryIdeaRequestRepository() IdeaRequestConfiguration {
	return func(irs *IdeaRequestService) error {
		ideasReqs := ideaReqMemory.New()
		irs.ideaRequests = ideasReqs
		return nil
	}
}

// CreateIdeaRequest creates a new idea request
func (is *IdeaRequestService) CreateIdeaRequest(name string, description string) (idea_request.IdeaRequest, error) {
	ir, err := idea_request.NewIdeaRequest(name, description, 0)
	if err != nil {
		return idea_request.IdeaRequest{}, err
	}
	err = is.ideaRequests.Add(ir)
	if err != nil {
		return idea_request.IdeaRequest{}, err
	}
	final, err := is.ideaRequests.GetById(ir.GetID())
	if err != nil {
		return idea_request.IdeaRequest{}, err
	}
	return final, nil
}

// UpvoteIdeaRequest upvotes and idea
func (is *IdeaRequestService) UpvoteIdeaRequest(id uuid.UUID) (idea_request.IdeaRequest, error) {
	ir, err := is.ideaRequests.GetById(id)
	if err != nil {
		return idea_request.IdeaRequest{}, err
	}
	ir.SetUpvotes(ir.GetUpvotes() + 1)
	err = is.ideaRequests.Update(ir)
	if err != nil {
		return idea_request.IdeaRequest{}, err
	}
	return ir, nil
}

func (is *IdeaRequestService) ListAllIdeas() ([]idea_request.IdeaRequest, error) {
	ir, err := is.ideaRequests.GetAll()
	if err != nil {
		return []idea_request.IdeaRequest{}, err
	}
	return ir, nil

}

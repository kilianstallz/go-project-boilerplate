package memory

import (
	"github.com/google/uuid"
	"server-mk3/internal/domain/idea_request"
	"sync"
)

type MemoryIdeaRequestRepository struct {
	ideaRequests map[uuid.UUID]idea_request.IdeaRequest
	sync.Mutex
}

func New() *MemoryIdeaRequestRepository {
	return &MemoryIdeaRequestRepository{
		ideaRequests: make(map[uuid.UUID]idea_request.IdeaRequest),
	}
}

func (irr *MemoryIdeaRequestRepository) GetAll() ([]idea_request.IdeaRequest, error) {
	var ideaRequests []idea_request.IdeaRequest
	for _, ir := range irr.ideaRequests {
		ideaRequests = append(ideaRequests, ir)
	}
	return ideaRequests, nil
}

func (irr *MemoryIdeaRequestRepository) GetById(id uuid.UUID) (idea_request.IdeaRequest, error) {
	if ir, ok := irr.ideaRequests[uuid.UUID(id)]; ok {
		return ir, nil
	}
	return idea_request.IdeaRequest{}, idea_request.ErrIdeaRequestNotFound
}

func (irr *MemoryIdeaRequestRepository) Add(newReq idea_request.IdeaRequest) error {
	irr.Lock()
	defer irr.Unlock()

	if _, ok := irr.ideaRequests[newReq.GetID()]; ok {
		return idea_request.ErrIdeaRequestAlreadyExist
	}
	irr.ideaRequests[newReq.GetID()] = newReq
	return nil
}

func (irr *MemoryIdeaRequestRepository) Update(newReq idea_request.IdeaRequest) error {
	irr.Lock()
	defer irr.Unlock()
	if _, ok := irr.ideaRequests[newReq.GetID()]; !ok {
		return idea_request.ErrIdeaRequestNotFound
	}
	irr.ideaRequests[newReq.GetID()] = newReq
	return nil
}

func (irr *MemoryIdeaRequestRepository) Delete(id uuid.UUID) error {
	irr.Lock()
	defer irr.Unlock()
	if _, ok := irr.ideaRequests[id]; ok {
		return idea_request.ErrIdeaRequestAlreadyExist
	}
	delete(irr.ideaRequests, id)
	return nil
}

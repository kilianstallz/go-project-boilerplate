package idea_request

import (
	"testing"
)

func TestNewIdeaRequestService(t *testing.T) {
	irs, err := NewIdeaRequestService(
		WithMemoryIdeaRequestRepository())

	if err != nil {
		t.Error(err)
	}

	idea, err := irs.CreateIdeaRequest("Hello", "Desc")
	if err != nil {
		t.Error(err)
	}

	// Perform upvote
	_, err = irs.UpvoteIdeaRequest(idea.GetID())
	if err != nil {
		t.Error(err)
	}
	getIR, err := irs.ideaRequests.GetById(idea.GetID())
	if err != nil {
		t.Error(err)
	}
	if getIR.GetUpvotes() != idea.GetUpvotes()+1 {
		t.Error("Upvote did not upvote")
	}
}

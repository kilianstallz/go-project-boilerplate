package memory

import (
	"github.com/google/uuid"
	"server-mk3/internal/domain/idea_request"
	"testing"
)

func TestMemoryIdeaRequestRepository_Add(t *testing.T) {
	repo := New()

	idea, err := idea_request.NewIdeaRequest("Test", "The wow effect", 134)
	if err != nil {
		t.Error(err)
	}
	repo.Add(idea)
	if len(repo.ideaRequests) != 1 {
		t.Errorf("Expected 1 idea, got %d", len(repo.ideaRequests))
	}
}

func TestMemoryIdeaRequestRepository_GetById(t *testing.T) {
	repo := New()
	idea, err := idea_request.NewIdeaRequest("Test", "wow effect", 1)
	if err != nil {
		t.Error(err)
	}
	repo.Add(idea)
	if len(repo.ideaRequests) != 1 {
		t.Errorf("Expected 1 idea, got %d", len(repo.ideaRequests))
	}
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}
	testCases := []testCase{
		{
			name:        "Get idea by id",
			id:          idea.GetID(),
			expectedErr: nil,
		},
		{
			name:        "Get non-existing product by id",
			id:          uuid.New(),
			expectedErr: idea_request.ErrIdeaRequestNotFound,
		},
	}
	for _, ir := range testCases {
		t.Run(ir.name, func(t *testing.T) {

			_, err := repo.GetById(ir.id)
			if err != ir.expectedErr {
				t.Errorf("Expected error, got %v", err)
			}
		})
	}

}

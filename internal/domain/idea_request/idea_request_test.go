package idea_request

import (
	"testing"
)

func TestNewIdeaRequest(t *testing.T) {

	t.Run("Should throw error when name is empty", func(t *testing.T) {
		_, err := NewIdeaRequest("", "desc", 0)
		if err != ErrMissingValues {
			t.Errorf("Expected error, got nil")
		}
	})
}

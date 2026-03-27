package store

import (
	"context"
	"testing"

	tracemodel "github.com/KevinAbura/backend-power/internal/trace/model"
)

func TestInMemoryStore(t *testing.T) {
	ctx := context.Background()
	s := &InMemoryStore{}

	if err := s.SaveWorkItem(ctx, tracemodel.WorkItem{ID: "w1"}); err != nil {
		t.Fatalf("SaveWorkItem returned error: %v", err)
	}
	if err := s.SaveWorkItemDesignLink(ctx, tracemodel.WorkItemDesignLink{WorkItemID: "w1"}); err != nil {
		t.Fatalf("SaveWorkItemDesignLink returned error: %v", err)
	}
	if err := s.SaveImplementationRun(ctx, tracemodel.ImplementationRun{ID: "r1"}); err != nil {
		t.Fatalf("SaveImplementationRun returned error: %v", err)
	}
	if err := s.SavePullRequestRef(ctx, tracemodel.PullRequestRef{PRNumber: 1}); err != nil {
		t.Fatalf("SavePullRequestRef returned error: %v", err)
	}

	if len(s.WorkItems) != 1 || len(s.WorkItemDesignLinks) != 1 || len(s.ImplementationRuns) != 1 || len(s.PullRequests) != 1 {
		t.Fatalf("unexpected store state: %+v", s)
	}
}

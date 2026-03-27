package flow

import (
	"context"
	"strings"
	"testing"

	designmodel "github.com/KevinAbura/backend-power/internal/design/model"
	tracestore "github.com/KevinAbura/backend-power/internal/trace/store"
)

func TestCreateIssueFlowRun(t *testing.T) {
	store := &tracestore.InMemoryStore{}
	flow := CreateIssueFlow{Store: store}

	out, err := flow.Run(context.Background(), CreateIssueInput{
		Repository: "KevinAbura/backend-power",
		WorkItemID: "w1",
		WorkItemKey: "BP-1",
		Title:      "Create checkout API issue",
		Summary:    "Draft the backend issue for checkout API support.",
		DesignReference: designmodel.DesignReference{
			ID:        "d1",
			Provider:  "figma",
			FileKey:   "file-1",
			NodeID:    "node-1",
			FrameName: "Checkout",
			DesignURL: "https://figma.example/checkout",
		},
		Scope:       []string{"API", "Validation"},
		Constraints: []string{"Reuse existing error model"},
	})
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	if len(store.WorkItems) != 1 {
		t.Fatalf("expected one work item, got %d", len(store.WorkItems))
	}
	if len(store.WorkItemDesignLinks) != 1 {
		t.Fatalf("expected one design link, got %d", len(store.WorkItemDesignLinks))
	}
	if !strings.Contains(out.IssueBody, "Create checkout API issue") && !strings.Contains(out.IssueBody, "Draft the backend issue") {
		t.Fatalf("expected issue body content, got %q", out.IssueBody)
	}
}

func TestCreateIssueFlowValidation(t *testing.T) {
	flow := CreateIssueFlow{}
	_, err := flow.Run(context.Background(), CreateIssueInput{})
	if err == nil {
		t.Fatalf("expected validation error")
	}
}

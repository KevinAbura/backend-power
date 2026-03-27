package flow

import (
	"context"
	"fmt"
	"time"

	designmodel "github.com/KevinAbura/backend-power/internal/design/model"
	ghtracker "github.com/KevinAbura/backend-power/internal/tracker/github"
	tracemodel "github.com/KevinAbura/backend-power/internal/trace/model"
	tracestore "github.com/KevinAbura/backend-power/internal/trace/store"
)

type CreateIssueInput struct {
	Repository      string
	WorkItemID      string
	WorkItemKey     string
	Title           string
	Summary         string
	DesignReference designmodel.DesignReference
	Scope           []string
	Constraints     []string
}

type CreateIssueOutput struct {
	IssueBody string
	WorkItem  tracemodel.WorkItem
	Link      tracemodel.WorkItemDesignLink
}

type CreateIssueFlow struct {
	Store tracestore.Store
}

func (f CreateIssueFlow) Run(ctx context.Context, input CreateIssueInput) (CreateIssueOutput, error) {
	if input.Repository == "" {
		return CreateIssueOutput{}, fmt.Errorf("repository is required")
	}
	if input.Title == "" {
		return CreateIssueOutput{}, fmt.Errorf("title is required")
	}
	if f.Store == nil {
		return CreateIssueOutput{}, fmt.Errorf("store is required")
	}

	payload := ghtracker.IssuePayload{
		Title:       input.Title,
		Summary:     input.Summary,
		DesignURL:   input.DesignReference.DesignURL,
		Scope:       input.Scope,
		Constraints: input.Constraints,
	}

	workItem := tracemodel.WorkItem{
		ID:          input.WorkItemID,
		Provider:    "github",
		ExternalKey: input.WorkItemKey,
		Title:       input.Title,
		Type:        "issue",
		Status:      "draft",
	}
	link := tracemodel.WorkItemDesignLink{
		WorkItemID:        input.WorkItemID,
		DesignReferenceID: input.DesignReference.ID,
		RelationshipType:  "source",
	}

	if err := f.Store.SaveWorkItem(ctx, workItem); err != nil {
		return CreateIssueOutput{}, err
	}
	if !input.DesignReference.IsZero() {
		if err := f.Store.SaveWorkItemDesignLink(ctx, link); err != nil {
			return CreateIssueOutput{}, err
		}
	}

	_ = time.Now()

	return CreateIssueOutput{
		IssueBody: ghtracker.RenderIssueBody(payload),
		WorkItem:  workItem,
		Link:      link,
	}, nil
}

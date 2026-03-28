package github

import (
	"context"
	"fmt"
)

type IssuePayload struct {
	Title       string
	Summary     string
	DesignURL   string
	Scope       []string
	Constraints []string
}

type IssueAdapter interface {
	CreateIssue(ctx context.Context, repo string, payload IssuePayload) error
}

type Client struct{}

func (c Client) CreateIssue(ctx context.Context, repo string, payload IssuePayload) error {
	if repo == "" {
		return fmt.Errorf("repo is required")
	}
	if payload.Title == "" {
		return fmt.Errorf("issue title is required")
	}

	// TODO: wire this adapter to the GitHub issue creation flow once the
	// repository command integration is ready.
	return nil
}

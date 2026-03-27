package github

import (
	"context"
	"testing"
)

func TestClientCreateIssueValidation(t *testing.T) {
	client := Client{}
	ctx := context.Background()

	if err := client.CreateIssue(ctx, "", IssuePayload{Title: "x"}); err == nil {
		t.Fatalf("expected repo validation error")
	}

	if err := client.CreateIssue(ctx, "KevinAbura/backend-power", IssuePayload{}); err == nil {
		t.Fatalf("expected title validation error")
	}

	if err := client.CreateIssue(ctx, "KevinAbura/backend-power", IssuePayload{Title: "valid"}); err != nil {
		t.Fatalf("expected valid payload to pass, got %v", err)
	}
}

package github

import (
	"strings"
	"testing"
)

func TestRenderIssueBody(t *testing.T) {
	body := RenderIssueBody(IssuePayload{
		Summary:     "Implement checkout API",
		DesignURL:   "https://figma.example/design",
		Scope:       []string{"API", "Database"},
		Constraints: []string{"Reuse existing error model"},
	})

	checks := []string{
		"## Summary",
		"Implement checkout API",
		"## Design Reference",
		"https://figma.example/design",
		"## Backend Scope",
		"- API",
		"## Constraints",
		"Reuse existing error model",
	}

	for _, check := range checks {
		if !strings.Contains(body, check) {
			t.Fatalf("expected body to contain %q, got %q", check, body)
		}
	}
}

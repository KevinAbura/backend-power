package github

import "strings"

func RenderIssueBody(payload IssuePayload) string {
	var b strings.Builder

	if payload.Summary != "" {
		b.WriteString("## Summary\n")
		b.WriteString(payload.Summary)
		b.WriteString("\n\n")
	}

	if payload.DesignURL != "" {
		b.WriteString("## Design Reference\n")
		b.WriteString("- Design URL: ")
		b.WriteString(payload.DesignURL)
		b.WriteString("\n\n")
	}

	if len(payload.Scope) > 0 {
		b.WriteString("## Backend Scope\n")
		for _, item := range payload.Scope {
			b.WriteString("- ")
			b.WriteString(item)
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}

	if len(payload.Constraints) > 0 {
		b.WriteString("## Constraints\n")
		for _, item := range payload.Constraints {
			b.WriteString("- ")
			b.WriteString(item)
			b.WriteString("\n")
		}
	}

	return strings.TrimSpace(b.String())
}

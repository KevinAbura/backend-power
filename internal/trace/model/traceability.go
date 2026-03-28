package model

import "time"

type WorkItem struct {
	ID          string
	Provider    string
	ExternalKey string
	Title       string
	Type        string
	Status      string
}

type WorkItemDesignLink struct {
	WorkItemID        string
	DesignReferenceID string
	RelationshipType  string
}

type ImplementationRun struct {
	ID          string
	WorkItemID  string
	Repository  string
	BranchName  string
	BaseBranch  string
	Status      string
	StartedAt   time.Time
	CompletedAt *time.Time
}

type PullRequestRef struct {
	ImplementationRunID string
	Provider            string
	PRNumber            int
	URL                 string
	Status              string
}

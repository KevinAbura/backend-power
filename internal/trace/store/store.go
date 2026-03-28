package store

import (
	"context"

	tracemodel "github.com/KevinAbura/backend-power/internal/trace/model"
)

type Store interface {
	SaveWorkItem(ctx context.Context, item tracemodel.WorkItem) error
	SaveWorkItemDesignLink(ctx context.Context, link tracemodel.WorkItemDesignLink) error
	SaveImplementationRun(ctx context.Context, run tracemodel.ImplementationRun) error
	SavePullRequestRef(ctx context.Context, ref tracemodel.PullRequestRef) error
}

type InMemoryStore struct {
	WorkItems           []tracemodel.WorkItem
	WorkItemDesignLinks []tracemodel.WorkItemDesignLink
	ImplementationRuns  []tracemodel.ImplementationRun
	PullRequests        []tracemodel.PullRequestRef
}

func (s *InMemoryStore) SaveWorkItem(ctx context.Context, item tracemodel.WorkItem) error {
	_ = ctx
	s.WorkItems = append(s.WorkItems, item)
	return nil
}

func (s *InMemoryStore) SaveWorkItemDesignLink(ctx context.Context, link tracemodel.WorkItemDesignLink) error {
	_ = ctx
	s.WorkItemDesignLinks = append(s.WorkItemDesignLinks, link)
	return nil
}

func (s *InMemoryStore) SaveImplementationRun(ctx context.Context, run tracemodel.ImplementationRun) error {
	_ = ctx
	s.ImplementationRuns = append(s.ImplementationRuns, run)
	return nil
}

func (s *InMemoryStore) SavePullRequestRef(ctx context.Context, ref tracemodel.PullRequestRef) error {
	_ = ctx
	s.PullRequests = append(s.PullRequests, ref)
	return nil
}

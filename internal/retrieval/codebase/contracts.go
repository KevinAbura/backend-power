package codebase

import "context"

type ReferenceResult struct {
	Endpoints  []string
	Services   []string
	Migrations []string
	Tests      []string
	Guidance   []string
}

type Retriever interface {
	FindReferences(ctx context.Context, repo string, query string) (ReferenceResult, error)
}

type InMemoryRetriever struct {
	Results map[string]ReferenceResult
}

func (r InMemoryRetriever) FindReferences(ctx context.Context, repo string, query string) (ReferenceResult, error) {
	_ = ctx
	if r.Results == nil {
		return ReferenceResult{}, nil
	}
	return r.Results[query], nil
}

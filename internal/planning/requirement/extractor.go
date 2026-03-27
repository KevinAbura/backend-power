package requirement

import "github.com/KevinAbura/backend-power/internal/design/model"

type ExtractedRequirement struct {
	Entities       []string
	Actions        []string
	Validations    []string
	Capabilities   []string
	CandidateTasks []string
}

type Extractor interface {
	Extract(model.DesignReference, string) ExtractedRequirement
}

type BasicExtractor struct{}

func (b BasicExtractor) Extract(ref model.DesignReference, notes string) ExtractedRequirement {
	result := ExtractedRequirement{
		Capabilities: []string{"analyze design", "derive backend work items"},
	}

	if ref.FrameName != "" {
		result.CandidateTasks = append(result.CandidateTasks, "map frame to backend scope: "+ref.FrameName)
	}
	if ref.FlowName != "" {
		result.CandidateTasks = append(result.CandidateTasks, "derive work items for flow: "+ref.FlowName)
	}
	if notes != "" {
		result.Actions = append(result.Actions, "incorporate product notes into planning")
	}

	return result
}

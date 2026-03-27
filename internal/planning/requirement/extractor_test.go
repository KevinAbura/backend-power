package requirement

import (
	"testing"

	designmodel "github.com/KevinAbura/backend-power/internal/design/model"
)

func TestBasicExtractorExtract(t *testing.T) {
	extractor := BasicExtractor{}
	ref := designmodel.DesignReference{
		FrameName: "Checkout",
		FlowName:  "Place Order",
	}

	result := extractor.Extract(ref, "support coupon validation")

	if len(result.Capabilities) == 0 {
		t.Fatalf("expected capabilities to be populated")
	}
	if len(result.CandidateTasks) < 2 {
		t.Fatalf("expected candidate tasks for frame and flow, got %v", result.CandidateTasks)
	}
	if len(result.Actions) != 1 {
		t.Fatalf("expected notes to produce one action, got %v", result.Actions)
	}
}

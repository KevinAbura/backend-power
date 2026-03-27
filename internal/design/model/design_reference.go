package model

import "time"

// DesignReference identifies a design artifact that can be linked to work items
// and implementation evidence.
type DesignReference struct {
	ID          string
	Provider    string
	FileKey     string
	NodeID      string
	FrameName   string
	FlowName    string
	DesignURL   string
	ExtractedAt time.Time
}

func (r DesignReference) IsZero() bool {
	return r.Provider == "" && r.FileKey == "" && r.NodeID == ""
}

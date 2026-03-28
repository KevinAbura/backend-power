package router

import "testing"

func TestRoute(t *testing.T) {
	tests := []struct {
		input   string
		wantOK  bool
		wantCmd Command
	}{
		{input: "bp:find-reference", wantOK: true, wantCmd: CommandFindRef},
		{input: "  BP:IMPLEMENT  ", wantOK: true, wantCmd: CommandImplement},
		{input: "unknown", wantOK: false},
	}

	for _, tt := range tests {
		got, ok := Route(tt.input)
		if ok != tt.wantOK {
			t.Fatalf("Route(%q) ok = %v, want %v", tt.input, ok, tt.wantOK)
		}
		if got != tt.wantCmd {
			t.Fatalf("Route(%q) command = %q, want %q", tt.input, got, tt.wantCmd)
		}
	}
}

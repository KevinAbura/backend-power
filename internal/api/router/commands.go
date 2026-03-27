package router

import "strings"

type Command string

const (
	CommandAnalyzeFigma Command = "bp:analyze-figma"
	CommandCreateIssue  Command = "bp:create-issue"
	CommandFindRef      Command = "bp:find-reference"
	CommandImplement    Command = "bp:implement"
	CommandOpenPR       Command = "bp:open-pr"
	CommandReviewPR     Command = "bp:review-pr"
)

func Route(input string) (Command, bool) {
	trimmed := strings.TrimSpace(strings.ToLower(input))
	switch trimmed {
	case string(CommandAnalyzeFigma):
		return CommandAnalyzeFigma, true
	case string(CommandCreateIssue):
		return CommandCreateIssue, true
	case string(CommandFindRef):
		return CommandFindRef, true
	case string(CommandImplement):
		return CommandImplement, true
	case string(CommandOpenPR):
		return CommandOpenPR, true
	case string(CommandReviewPR):
		return CommandReviewPR, true
	default:
		return "", false
	}
}

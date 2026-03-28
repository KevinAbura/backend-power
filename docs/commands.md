# Commands

Backend Power can expose explicit keyword-driven workflows inspired by skill-based developer tooling.

## V1 commands

### `bp:analyze-figma`
Parse a Figma file or node into backend-relevant signals.

Typical outputs:
- entities
- actions
- validations
- candidate backend capabilities
- candidate work items

### `bp:create-issue`
Create or enrich a GitHub Issue from design and planning context.

Typical outputs:
- issue title and summary
- design linkage
- backend scope
- acceptance criteria
- verification notes

### `bp:find-reference`
Retrieve similar implementations from the repository.

Typical outputs:
- similar endpoints
- similar services
- similar migrations
- similar tests
- repository style guidance

### `bp:implement`
Generate a patch-oriented implementation plan or code proposal grounded in repository patterns.

Typical outputs:
- implementation plan
- target files
- proposed patch
- test additions

### `bp:open-pr`
Open a pull request with work item linkage and verification context.

### `bp:review-pr`
Review a PR for correctness, repository alignment, and missing coverage.

## Rule of thumb

Implementation in Backend Power should be repository-grounded.
In practice, `bp:find-reference` should happen before or inside `bp:implement`.

## Future directions

Later versions may add commands for:
- staging verification
- issue closure
- release evidence
- Jira sync

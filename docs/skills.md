# Skills

Backend Power organizes its workflow around a small set of composable skills.

## V1 skills

### `design-to-work-item`
Transforms design references into backend-oriented work items.

Responsibilities:
- parse design metadata
- extract backend-relevant requirements
- identify entities, actions, validations, and state transitions
- produce candidate issues or tasks

### `repository-grounded-implementation`
Produces implementation guidance and patches that align with the existing repository.

Responsibilities:
- retrieve similar code
- infer repository patterns
- generate patch-oriented changes
- recommend tests and verification scope

### `pull-request-review`
Reviews proposed changes for correctness, consistency, and maintainability.

Responsibilities:
- inspect diffs against linked work items
- compare changes with repository patterns
- flag missing tests and risky deviations
- summarize review findings

## Triggering model

A skill can be triggered in two ways:
- explicit command usage, such as `bp:find-reference`
- natural language routing, such as asking to find similar code or review a PR

## Important constraint

Backend Power should avoid greenfield-style generation inside existing repositories.

That means implementation-related skills should use repository retrieval and pattern inference before producing patches.

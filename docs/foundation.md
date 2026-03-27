# Foundation progress

This document tracks the first implementation slice of the V1 roadmap.

## Included in this foundation PR

- Go module bootstrap
- design reference model
- traceability model
- GitHub issue adapter skeleton
- design-to-work-item extractor skeleton

## Why start here?

These pieces establish the minimum internal contracts needed for later steps:
- design ingestion needs a stable reference model
- issue creation needs a work item payload contract
- implementation and review need traceability objects
- planning needs a requirement extraction interface

## Next suggested slices

1. add persistence interfaces for design and traceability models
2. add repository retrieval contracts for reference lookup
3. add a small command router for V1 commands
4. add tests for the extractor and adapter validation behavior

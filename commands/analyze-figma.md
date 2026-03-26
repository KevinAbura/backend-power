# bp:analyze-figma

## Purpose

Parse a Figma file or node into backend-relevant signals.

## Inputs

- Figma file key
- Figma node ID
- Optional product notes

## Outputs

- extracted entities
- user actions
- candidate validations
- inferred backend capabilities
- candidate work items

## Usage

Use this command when starting from design and needing a backend-oriented view of the work.

## Notes

This command should not treat Figma as a complete backend specification. It should produce structured signals for later planning and issue creation.

# bp:review-pr

## Purpose

Review a pull request for correctness, repository alignment, and missing coverage.

## Inputs

- PR diff
- linked issue
- reference implementation set

## Outputs

- review summary
- style drift warnings
- missing test warnings
- architecture mismatch warnings

## Usage

Use this command after a PR is open or when a review is requested.

## Notes

Reviews should check whether the patch follows existing project patterns and avoids unnecessary abstractions.

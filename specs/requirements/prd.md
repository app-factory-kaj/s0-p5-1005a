# greeter — PRD

## Problem Statement

Teams building and integrating services on the platform need a small, predictable reference endpoint to greet a caller by name — useful as a smoke-test target, an onboarding example, and a conformance check against the organization's own service conventions (`app-factory-kaj/e2e-reference`). Today there is no minimal, standards-following service to point at for this. E2E marker s0p5-1005a.

## Solution

Greeter is a small Go HTTP service exposing a single endpoint, `GET /hello`, that returns a JSON greeting for the name supplied in the query string. It is built to the structure and conventions demonstrated in `app-factory-kaj/e2e-reference`, so it also serves as a clean, minimal example of those conventions in practice.

## Actors

- **API Consumer** — any client (a developer, a script, or another service) that calls `GET /hello` to receive a JSON greeting.

## User Stories

1. As an API Consumer, I want to call `GET /hello?name=X`, so that I receive a JSON greeting addressed to "X".
2. As an API Consumer, I want to call `GET /hello` without a `name`, so that I still receive a usable, generic JSON greeting instead of an error.

## Product Decisions

- Missing `name`: the service returns a generic greeting (e.g. "Hello, World!") rather than rejecting the request. *assumed*
- The endpoint is open — it requires no authentication and is not placed behind the platform API gateway's auth. *assumed*
- The service follows the structure and conventions shown in `app-factory-kaj/e2e-reference`.

## Out of Scope

- Persistence of any kind — the service is stateless.
- Additional endpoints beyond `GET /hello`.
- Rate limiting, quotas, or per-caller usage tracking.
- Localization of the greeting text.

## Open Questions

None at this time.
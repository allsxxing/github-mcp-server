# Test Coverage GitHub Issues

This document contains issue templates for improving test coverage. Create these issues in GitHub to track the testing work.

## Priority 1 Issues (Critical)

### Issue 1: Add test coverage for pkg/utils/result.go

**Title:** Add test coverage for pkg/utils/result.go (Priority 1)

**Labels:** `testing`, `priority-1`, `good-first-issue`

**Body:**
```markdown
## Overview
The `pkg/utils/result.go` file currently has **0% test coverage** and contains widely-used utility functions for creating MCP tool results.

## Priority
**Priority 1 (Critical)** - High usage across codebase, simple to test

## Current Coverage
- Coverage: 0%
- Functions to test: 4
- Estimated effort: 2-3 hours

## Functions Requiring Tests
1. `NewToolResultText()` - Creates text tool results
2. `NewToolResultError()` - Creates error tool results
3. `NewToolResultErrorFromErr()` - Creates error from error object
4. `NewToolResultResource()` - Creates resource-embedded results

## Test Requirements
- Test various message formats (empty, long, special characters, unicode)
- Verify `IsError` flag is set correctly
- Test error message concatenation
- Test resource embedding
- Edge cases: nil errors, empty messages

## Test Skeleton
A complete test skeleton has been created at `pkg/utils/result_test.go` with all necessary test cases ready to be implemented.

## Related
- See `docs/test-coverage-analysis.md` for full coverage analysis
- Part of Phase 1 (Quick Wins) in the testing roadmap

## Definition of Done
- [ ] All 4 functions have unit tests
- [ ] Coverage reaches 100%
- [ ] All edge cases covered
- [ ] Tests pass in CI
```

---

### Issue 2: Add test coverage for pkg/buffer/buffer.go

**Title:** Add test coverage for pkg/buffer/buffer.go (Priority 1)

**Labels:** `testing`, `priority-1`

**Body:**
```markdown
## Overview
The `pkg/buffer/buffer.go` file currently has **0% test coverage** and contains critical ring buffer logic for processing HTTP response logs.

## Priority
**Priority 1 (Critical)** - Complex data processing logic, critical for correctness

## Current Coverage
- Coverage: 0%
- Functions to test: 1 (ProcessResponseAsRingBufferToEnd)
- Estimated effort: 4-6 hours

## Test Requirements

### Core Functionality
- Empty HTTP response body
- Single line response
- Response with exactly maxJobLogLines
- Response with < maxJobLogLines
- Response with > maxJobLogLines

### Ring Buffer Behavior
- Multiple wraparounds with large input
- Verify only last N lines are retained
- Test with various buffer sizes (1, 3, 10, 100, 1000)

### Edge Cases
- Response with > 100,000 lines (cap enforcement)
- Very long individual lines (> 1MB)
- Empty lines and newline handling
- Special characters and unicode
- maxLines = 0, 1, negative values

### Error Handling
- Scanner errors during reading
- Buffer capacity limits

## Test Skeleton
A comprehensive test skeleton has been created at `pkg/buffer/buffer_test.go` with all necessary test cases.

## Related
- See `docs/test-coverage-analysis.md` for full coverage analysis
- Part of Phase 2 (Critical Components) in the testing roadmap

## Definition of Done
- [ ] All test cases implemented
- [ ] Coverage reaches 90%+
- [ ] Ring buffer correctness verified
- [ ] Edge cases covered
- [ ] Performance acceptable for large inputs
- [ ] Tests pass in CI
```

---

### Issue 3: Improve test coverage for internal/ghmcp/server.go

**Title:** Improve test coverage for internal/ghmcp/server.go (Priority 1)

**Labels:** `testing`, `priority-1`, `core`

**Body:**
```markdown
## Overview
The `internal/ghmcp/server.go` file currently has **28.6% test coverage**. This is the core MCP server implementation and requires significantly better coverage.

## Priority
**Priority 1 (Critical)** - Core server logic

## Current Coverage
- Coverage: 28.6%
- Target coverage: 65%+
- Estimated effort: 8-10 hours

## Areas Requiring Tests

### Client Creation (`createGitHubClients`)
- Valid configuration
- Invalid token
- API host parsing for github.com
- API host parsing for GitHub Enterprise
- Bearer auth transport setup
- GraphQL client initialization
- Raw client initialization

### Server Configuration
- MCPServerConfig validation
- EnabledToolsets filtering
- EnabledTools registration
- Feature flag handling
- Read-only mode enforcement
- DynamicToolsets behavior
- LockdownMode setup

### Server Lifecycle
- Server initialization
- Configuration loading from environment
- Signal handling (SIGTERM, SIGINT)
- Graceful shutdown
- Error handling and recovery

### Translation and Logging
- Translator initialization
- Logger setup and usage
- RepoAccessTTL configuration
- TokenScopes filtering

## Related
- Current tests in `internal/ghmcp/server_test.go`
- See `docs/test-coverage-analysis.md` for full coverage analysis
- Part of Phase 3 (Core Server) in the testing roadmap

## Definition of Done
- [ ] Coverage increases from 28.6% to 65%+
- [ ] All critical paths tested
- [ ] Error handling verified
- [ ] Configuration validation covered
- [ ] Tests pass in CI
```

---

### Issue 4: Improve security test coverage for pkg/lockdown

**Title:** Improve security test coverage for pkg/lockdown/lockdown.go (Priority 1)

**Labels:** `testing`, `priority-1`, `security`

**Body:**
```markdown
## Overview
The `pkg/lockdown/lockdown.go` file currently has **47.8% test coverage**. This is security-critical code that requires comprehensive testing.

## Priority
**Priority 1 (Critical)** - Security component

## Current Coverage
- Coverage: 47.8%
- Target coverage: 75%+
- Estimated effort: 6-8 hours

## Critical Untested Functions

### IsSafeContent (0% coverage) 🔴
- **Security Critical**: Validates content safety
- Test with various content types
- Test permission levels
- Test edge cases and boundary conditions

### isTrustedBot (0% coverage) 🔴
- **Security Critical**: Bot validation
- Test known bots (GitHub, Dependabot, etc.)
- Test unknown/malicious bots
- Test bot name variations

### SetLogger (0% coverage)
- Logger initialization
- Nil logger handling

## Areas Requiring Improvement

### Cache Management
- Cache expiration and renewal
- TTL behavior
- Concurrent access patterns
- Cache key generation

### Repository Access
- `getRepoAccessInfo` error paths (46.4% coverage)
- Permission level validation
- Public vs private repository handling
- Collaborator permission checking

### Error Handling
- Network errors during GraphQL queries
- Invalid repository references
- Permission denied scenarios

## Test Requirements
- Unit tests for all security-critical functions
- Integration tests for cache behavior
- Concurrent access testing
- Permission boundary testing
- Mock GraphQL responses for various scenarios

## Related
- Current tests in `pkg/lockdown/lockdown_test.go`
- See `docs/test-coverage-analysis.md` for full coverage analysis
- Part of Phase 2 (Critical Components) in the testing roadmap

## Security Implications
This code controls repository access and content safety. Inadequate testing could lead to:
- Unauthorized repository access
- Security policy bypasses
- Data exposure

## Definition of Done
- [ ] Coverage increases from 47.8% to 75%+
- [ ] All security-critical functions tested (IsSafeContent, isTrustedBot)
- [ ] Cache behavior verified
- [ ] Concurrent access tested
- [ ] Error paths covered
- [ ] Security review of tests
- [ ] Tests pass in CI
```

---

## Priority 2 Issues

### Issue 5: Add test coverage for pkg/translations

**Title:** Add test coverage for pkg/translations/translations.go (Priority 2)

**Labels:** `testing`, `priority-2`, `i18n`

**Body:**
```markdown
## Overview
The `pkg/translations/translations.go` file currently has **0% test coverage** and handles internationalization and configuration.

## Priority
**Priority 2 (Medium)** - Important for i18n support

## Current Coverage
- Coverage: 0%
- Target coverage: 80%+
- Estimated effort: 3-4 hours

## Functions Requiring Tests

### NullTranslationHelper
- Returns default values
- Ignores key parameter

### TranslationHelper
- Translation key lookup
- Environment variable overrides (GITHUB_MCP_*)
- Config file loading from JSON
- Default value fallback
- Case insensitivity (keys converted to uppercase)
- Caching behavior

### DumpTranslationKeyMap
- File creation
- JSON marshaling
- Error handling for file I/O
- Overwrites existing files

## Test Requirements
- Config file loading (valid/invalid JSON)
- Environment variable precedence
- Key normalization (case insensitivity)
- Cleanup function behavior
- Error scenarios
- Missing config file handling

## Test Skeleton
A complete test skeleton has been created at `pkg/translations/translations_test.go` with all necessary test cases.

## Related
- See `docs/test-coverage-analysis.md` for full coverage analysis
- Part of Phase 2 (Critical Components) in the testing roadmap

## Definition of Done
- [ ] All functions have unit tests
- [ ] Coverage reaches 80%+
- [ ] Environment variable override tested
- [ ] Config file loading verified
- [ ] Error handling covered
- [ ] Tests pass in CI
```

---

### Issue 6: Improve test coverage for pkg/inventory/errors.go

**Title:** Improve test coverage for pkg/inventory/errors.go (Priority 2)

**Labels:** `testing`, `priority-2`, `good-first-issue`

**Body:**
```markdown
## Overview
The `pkg/inventory/errors.go` file currently has **20% test coverage**. This file defines custom error types used throughout the inventory package.

## Priority
**Priority 2 (Medium)** - Error handling

## Current Coverage
- Coverage: 20%
- Target coverage: 90%+
- Estimated effort: 1-2 hours

## Areas Requiring Tests

### ToolsetDoesNotExistError
- Error message formatting ✓ (likely covered)
- `Is()` method - **Currently untested** 🔴
- Type assertions

### ToolDoesNotExistError
- Error message formatting
- Constructor function

## Test Requirements
- Test `ToolsetDoesNotExistError.Is()` with various targets
- Test error message content
- Test with `errors.Is()` function
- Test type assertions
- Test nil handling in `Is()` method

## Related
- Current tests may exist in `pkg/inventory/registry_test.go`
- See `docs/test-coverage-analysis.md` for full coverage analysis
- Part of Phase 1 (Quick Wins) in the testing roadmap

## Definition of Done
- [ ] `Is()` method fully tested
- [ ] Coverage reaches 90%+
- [ ] Error matching behavior verified
- [ ] Tests pass in CI
```

---

### Issue 7: Improve test coverage for pkg/github/workflow_prompts.go

**Title:** Improve test coverage for pkg/github/workflow_prompts.go (Priority 2)

**Labels:** `testing`, `priority-2`

**Body:**
```markdown
## Overview
The `pkg/github/workflow_prompts.go` file currently has **5.3% test coverage**. This file defines workflow prompts for GitHub operations.

## Priority
**Priority 2 (Medium)** - User-facing workflow

## Current Coverage
- Coverage: 5.3%
- Target coverage: 70%+
- Estimated effort: 2-3 hours

## Functions Requiring Tests

### IssueToFixWorkflowPrompt
- Prompt structure validation
- Argument requirements
- Translation key usage
- Prompt message generation
- GetPrompt callback behavior

## Test Requirements
- Verify all required arguments are defined
- Test argument validation (required vs optional)
- Test translation helper integration
- Test prompt metadata (name, description)
- Verify prompt can be called successfully

## Related
- See `docs/test-coverage-analysis.md` for full coverage analysis
- Part of Phase 2 (Critical Components) in the testing roadmap

## Definition of Done
- [ ] Prompt structure tested
- [ ] Coverage reaches 70%+
- [ ] Argument validation verified
- [ ] Translation integration tested
- [ ] Tests pass in CI
```

---

## Priority 3 Issues

### Issue 8: Add integration tests for main entry points

**Title:** Add integration tests for cmd packages (Priority 3)

**Labels:** `testing`, `priority-3`, `integration`

**Body:**
```markdown
## Overview
The main entry points (`cmd/github-mcp-server/main.go`, `cmd/mcpcurl/main.go`) currently have **0% test coverage**.

## Priority
**Priority 3** - Integration testing

## Current Coverage
- cmd/github-mcp-server/main.go: 0%
- cmd/github-mcp-server/generate_docs.go: 0%
- cmd/mcpcurl/main.go: 0%
- Target coverage: 40%+
- Estimated effort: 6-8 hours

## Test Requirements

### cmd/github-mcp-server/main.go
- Server startup and initialization
- Configuration loading from environment variables
- Signal handling (SIGTERM, SIGINT)
- Graceful shutdown
- Error handling and exit codes
- Command-line flag parsing

### cmd/github-mcp-server/generate_docs.go
- Documentation generation
- Output formatting
- Error handling

### cmd/mcpcurl/main.go
- MCP client initialization
- Request/response handling
- Error scenarios

## Approach
- Use integration tests rather than unit tests
- Test actual binary execution where possible
- Mock external dependencies (GitHub API)
- Test configuration scenarios

## Related
- See `docs/test-coverage-analysis.md` for full coverage analysis
- Part of Phase 4 (Integration) in the testing roadmap

## Definition of Done
- [ ] Server startup/shutdown tested
- [ ] Configuration loading verified
- [ ] Signal handling works
- [ ] Coverage reaches 40%+
- [ ] Tests pass in CI
```

---

### Issue 9: Add test coverage for internal/profiler

**Title:** Add test coverage for internal/profiler/profiler.go (Priority 3)

**Labels:** `testing`, `priority-3`

**Body:**
```markdown
## Overview
The `internal/profiler/profiler.go` file currently has **0% test coverage**. This is an optional debugging feature.

## Priority
**Priority 3** - Optional feature, lower priority

## Current Coverage
- Coverage: 0%
- Target coverage: 70%+
- Estimated effort: 3-4 hours

## Functions Requiring Tests

### Profiler Methods
- `ProfileFunc()` - Function profiling
- `ProfileFuncWithMetrics()` - Profiling with metrics
- `Start()` - Manual profiling start
- Memory delta calculation

### Global Functions
- `IsProfilingEnabled()` - Environment check
- `Init()` / `InitFromEnv()` - Initialization
- `ProfileFunc()` / `ProfileFuncWithMetrics()` - Global profiler usage

### Edge Cases
- `safeMemoryDelta()` - Overflow handling
- Disabled profiler behavior
- Nil logger handling

## Test Requirements
- Verify timing accuracy (within reasonable bounds)
- Test memory delta calculation
- Test overflow protection in `safeMemoryDelta()`
- Verify profiler enable/disable
- Test with nil logger
- Benchmark tests for profiler overhead

## Related
- See `docs/test-coverage-analysis.md` for full coverage analysis
- Part of Phase 4 (Integration) in the testing roadmap

## Definition of Done
- [ ] All profiler functions tested
- [ ] Coverage reaches 70%+
- [ ] Overflow handling verified
- [ ] Benchmarks added
- [ ] Tests pass in CI
```

---

## Creating These Issues

To create these issues in GitHub:

1. **Using GitHub CLI:**
```bash
gh issue create --title "TITLE" --body "BODY" --label "labels"
```

2. **Using GitHub Web UI:**
- Go to Issues → New Issue
- Copy the title and body from above
- Add the specified labels
- Submit

3. **Using GitHub API:**
```bash
curl -X POST \
  -H "Authorization: token YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  https://api.github.com/repos/OWNER/REPO/issues \
  -d '{"title":"TITLE","body":"BODY","labels":["label1","label2"]}'
```

## Tracking Progress

- Create a GitHub Project or Milestone to track these issues
- Use the labels to filter and prioritize
- Link PRs to issues as they are resolved
- Update `docs/test-coverage-analysis.md` as coverage improves

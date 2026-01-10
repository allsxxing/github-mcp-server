# Fork Setup and Maintenance Guide

This guide walks you through setting up a fork of the GitHub MCP Server repository and keeping it synchronized with the upstream repository.

## Initial Fork Setup

### Prerequisites

1. A GitHub account
2. Git installed on your local machine
3. [GitHub CLI](https://cli.github.com/) installed (optional, but recommended)
4. Go 1.24 or later installed
5. [golangci-lint v2](https://golangci-lint.run/welcome/install/#local-installation) installed

### Step 1: Sync Fork (Choose One Method)

#### Option A: Sync via GitHub UI (Recommended)

If you have an open sync PR in your fork:

1. Navigate to your fork's pull requests: `https://github.com/YOUR_USERNAME/github-mcp-server/pulls`
2. Find the sync PR (usually titled "Sync fork")
3. Click "Merge pull request"
4. Confirm the merge

#### Option B: Sync via GitHub CLI

```bash
gh pr merge 1 --repo YOUR_USERNAME/github-mcp-server --merge
```

Replace `YOUR_USERNAME` with your GitHub username and adjust the PR number if needed.

### Step 2: Clone Your Fork

```bash
# Clone your fork to your local machine
git clone https://github.com/YOUR_USERNAME/github-mcp-server.git
cd github-mcp-server
```

### Step 3: Configure Remotes

Set up the upstream repository to track the original project:

```bash
# Add the upstream remote
git remote add upstream https://github.com/github/github-mcp-server.git

# Verify remotes are configured correctly
git remote -v
```

Expected output:
```
origin    https://github.com/YOUR_USERNAME/github-mcp-server.git (fetch)
origin    https://github.com/YOUR_USERNAME/github-mcp-server.git (push)
upstream  https://github.com/github/github-mcp-server.git (fetch)
upstream  https://github.com/github/github-mcp-server.git (push)
```

### Step 4: Pull Latest Changes

```bash
# Fetch all branches from upstream
git fetch upstream

# Switch to your main branch
git checkout main

# Merge upstream changes
git merge upstream/main

# Push updates to your fork
git push origin main
```

## Building and Testing

### Step 5: Install Dependencies

```bash
# Download dependencies
go mod download

# Clean up and verify
go mod tidy
go mod verify
```

### Step 6: Build the Binary

```bash
# Create bin directory if it doesn't exist
mkdir -p bin

# Build the server binary
go build -o bin/github-mcp-server ./cmd/github-mcp-server
```

### Step 7: Verify the Build

```bash
# Check the binary version
./bin/github-mcp-server --version

# Verify the binary exists and check its size
ls -lh bin/github-mcp-server
```

### Step 8: Run Unit Tests

```bash
# Run all tests with verbose output
go test -v ./...

# Or use the test script
script/test
```

## Running the Server

### Step 9: Set Up Environment Variables

```bash
# Required: Set your GitHub Personal Access Token
export GITHUB_PERSONAL_ACCESS_TOKEN="ghp_YOUR_TOKEN_HERE"

# Optional: Configure specific toolsets
export GITHUB_TOOLSETS="repos,issues,pull_requests,actions"
```

**Note:** To create a GitHub Personal Access Token:
1. Go to GitHub Settings → Developer settings → Personal access tokens → Tokens (classic)
2. Click "Generate new token"
3. Select the necessary scopes based on your needs
4. Copy the token immediately (you won't be able to see it again)

For more details on token security, see the [Token Security Best Practices](../README.md#token-security-best-practices) section in the main README.

### Step 10: Test the Server

```bash
# Run the server in stdio mode
./bin/github-mcp-server stdio
```

The server should start and wait for MCP protocol messages on stdin.

## Advanced Testing

### Step 11: Run E2E Tests (Optional)

End-to-end tests require a GitHub token with appropriate permissions:

```bash
# Set the E2E test token (using GitHub CLI)
export GITHUB_MCP_SERVER_E2E_TOKEN=$(gh auth token)

# Or set it manually
export GITHUB_MCP_SERVER_E2E_TOKEN="ghp_YOUR_TOKEN_HERE"

# Run E2E tests
go test -v --tags e2e ./e2e
```

**Note:** E2E tests interact with the live GitHub API and may take longer to complete.

## Ongoing Fork Maintenance

### Syncing with Upstream

Keep your fork up-to-date with the upstream repository by running these commands periodically:

```bash
# Fetch latest changes from upstream
git fetch upstream

# Switch to main branch
git checkout main

# Merge upstream changes
git merge upstream/main

# Push updates to your fork
git push origin main
```

### Creating Feature Branches

When working on new features or fixes:

```bash
# Create and switch to a new feature branch
git checkout -b feature/your-feature-name

# Make your changes...
# Then stage and commit
git add .
git commit -m "feat: description of changes"

# Push your feature branch to your fork
git push origin feature/your-feature-name

# Create a pull request using GitHub CLI
gh pr create --base main --head feature/your-feature-name

# Or create a pull request manually through GitHub's web interface
```

### Development Workflow

When making changes to the codebase:

1. **Run linter** before committing:
   ```bash
   script/lint
   ```

2. **Update snapshots** if you modified MCP tools:
   ```bash
   UPDATE_TOOLSNAPS=true go test ./...
   ```

3. **Update documentation** if you changed tools:
   ```bash
   script/generate-docs
   ```

4. **Run tests** to ensure nothing broke:
   ```bash
   script/test
   ```

For more detailed contribution guidelines, see [CONTRIBUTING.md](../CONTRIBUTING.md).

## Troubleshooting

### Common Issues

**Build Failures:**
- Ensure you're using Go 1.24 or later: `go version`
- Try cleaning the build cache: `go clean -cache`
- Verify dependencies: `go mod verify`

**Test Failures:**
- Some tests may require network access
- E2E tests require a valid GitHub token
- Check if your token has the necessary permissions

**Merge Conflicts:**
- If you encounter merge conflicts when syncing:
  ```bash
  git fetch upstream
  git checkout main
  git merge upstream/main
  # Resolve conflicts in your editor
  git add .
  git commit -m "Merge upstream changes"
  git push origin main
  ```

**Outdated Dependencies:**
- Update dependencies with: `go get -u ./...`
- Run `go mod tidy` after updating
- Run `script/licenses` if you updated dependencies

## Additional Resources

- [Contributing Guidelines](../CONTRIBUTING.md)
- [Main README](../README.md)
- [Testing Documentation](testing.md)
- [Toolsets and Icons Guide](toolsets-and-icons.md)
- [GitHub's Fork Documentation](https://docs.github.com/en/get-started/quickstart/fork-a-repo)

## Quick Reference

```bash
# One-time setup
git clone https://github.com/YOUR_USERNAME/github-mcp-server.git
cd github-mcp-server
git remote add upstream https://github.com/github/github-mcp-server.git
go mod download
go build -o bin/github-mcp-server ./cmd/github-mcp-server

# Regular sync
git fetch upstream && git checkout main && git merge upstream/main && git push origin main

# Before committing
script/lint && script/test

# Feature development
git checkout -b feature/name
# ... make changes ...
git add . && git commit -m "feat: description"
git push origin feature/name
gh pr create
```

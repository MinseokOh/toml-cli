# Branch Strategy

## 🌳 Branch Naming Convention

### Main Branches
- `main` - Production-ready code
- `develop` - Development branch (if needed for complex features)

### Feature Branches
Use the following prefixes for feature branches:

#### Primary Types
- `feat/` - New features
  - `feat/add-json-export`
  - `feat/implement-bulk-operations`

- `fix/` - Bug fixes
  - `fix/handle-empty-toml-files`
  - `fix/memory-leak-in-parser`

- `hotfix/` - Critical production fixes
  - `hotfix/security-vulnerability`
  - `hotfix/crash-on-invalid-input`

#### Secondary Types
- `docs/` - Documentation updates
  - `docs/update-readme`
  - `docs/add-api-examples`

- `refactor/` - Code refactoring
  - `refactor/optimize-parser`
  - `refactor/simplify-cli-interface`

- `perf/` - Performance improvements
  - `perf/optimize-large-file-handling`
  - `perf/reduce-memory-usage`

- `test/` - Adding or updating tests
  - `test/add-integration-tests`
  - `test/improve-coverage`

- `chore/` - Maintenance tasks
  - `chore/update-dependencies`
  - `chore/add-ci-pipeline`

- `style/` - Code style changes
  - `style/format-code`
  - `style/fix-linting-issues`

## 🔄 Workflow

1. **Create Branch** from `main`
   ```bash
   git checkout main
   git pull origin main
   git checkout -b feat/your-feature-name
   ```

2. **Work on Feature**
   - Make commits with conventional commit messages
   - Keep commits focused and atomic

3. **Push Branch**
   ```bash
   git push -u origin feat/your-feature-name
   ```

4. **Create Pull Request**
   - Use the PR template
   - Fill out all required sections
   - Request review

5. **Merge**
   - Squash and merge for clean history
   - Delete branch after merge

## 📝 Branch Naming Rules

### ✅ Good Examples
- `feat/add-yaml-support`
- `fix/handle-unicode-characters`
- `docs/update-installation-guide`
- `refactor/modernize-error-handling`
- `chore/upgrade-go-version`

### ❌ Bad Examples
- `my-feature` (no prefix)
- `feat/Fix-Bug` (mixed case, wrong type)
- `feature/new-stuff` (vague description)
- `fix/various-fixes` (too generic)

## 🚀 Quick Commands

```bash
# Start new feature
git checkout -b feat/feature-name

# Start bug fix
git checkout -b fix/bug-description

# Start documentation update
git checkout -b docs/doc-update

# Start refactoring
git checkout -b refactor/component-name

# Start maintenance task
git checkout -b chore/task-description
```

## 🤖 AI-Assisted Development

When working with Claude Code:
- Branch names should be descriptive enough for AI context
- Include relevant context in commit messages
- Update `CLAUDE.md` when needed
- Mark PR as AI-assisted in the template
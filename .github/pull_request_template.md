# Pull Request

## 📋 Description

<!-- Provide a clear and concise description of what this PR does -->

### Type of Change

<!-- Mark the relevant option with an "x" -->

- [ ] 🐛 Bug fix (non-breaking change that fixes an issue)
- [ ] ✨ New feature (non-breaking change that adds functionality)
- [ ] 💥 Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] 📚 Documentation update (changes to documentation only)
- [ ] 🔧 Refactoring (code changes that neither fix bugs nor add features)
- [ ] ⚡ Performance improvement
- [ ] 🧪 Test improvements
- [ ] 🔨 Build/CI changes
- [ ] 🎨 Code style/formatting changes

## 🔗 Related Issues

<!-- Link to related issues using keywords like "Fixes", "Closes", "Resolves" -->
<!-- Example: Fixes #123, Closes #456 -->

- Fixes #<!-- issue number -->
- Related to #<!-- issue number -->

## 🧪 Testing

<!-- Describe how you tested your changes -->

### Test Coverage

- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] Benchmark tests added/updated (if performance-related)
- [ ] Manual testing completed

### Test Results

<!-- Provide test results or commands to verify the changes -->

```bash
# Commands used for testing
make test
make test-benchmark
```

**Test Output:**
```
<!-- Paste relevant test output here -->
```

## 📊 Performance Impact

<!-- If this PR affects performance, please provide benchmark results -->

- [ ] No performance impact
- [ ] Performance improvement (provide benchmarks)
- [ ] Performance regression (explain why it's acceptable)
- [ ] Performance impact unknown (requires further testing)

## 🔄 API Changes

<!-- If this PR changes the public API, describe the changes -->

- [ ] No API changes
- [ ] Backward compatible API additions
- [ ] Breaking API changes (provide migration guide)

### Migration Guide (if applicable)

<!-- Provide migration instructions for breaking changes -->

```go
// Before
// old code example

// After
// new code example
```

## 🏁 Code Quality

### Self-Review Checklist

- [ ] Code follows the project's coding standards
- [ ] Self-review of code completed
- [ ] Code is well-documented with clear comments
- [ ] No debugging/console statements left in code
- [ ] All functions have appropriate error handling
- [ ] Memory leaks and resource cleanup verified
- [ ] Thread safety considerations addressed (if applicable)

### Code Quality Tools

- [ ] `make lint` passes
- [ ] `make format` applied
- [ ] `make security-scan` passes
- [ ] No new warnings introduced

## 📚 Documentation

<!-- Check all that apply -->

- [ ] Code changes are self-documenting
- [ ] Godoc comments added/updated for public APIs
- [ ] README updated (if needed)
- [ ] API documentation updated (if needed)
- [ ] Examples added/updated (if needed)
- [ ] Changelog entry added (for significant changes)

## 🔐 Security Considerations

<!-- Describe any security implications of your changes -->

- [ ] No security implications
- [ ] Security review required
- [ ] Potential security improvement
- [ ] May introduce security risks (explain mitigation)

**Security Analysis:**
<!-- Describe any security considerations, potential vulnerabilities, or improvements -->

## 🖥️ Platform Compatibility

<!-- Check platforms where you've tested the changes -->

**Tested On:**

- [ ] Linux (Ubuntu 22.04)
- [ ] Linux (Ubuntu 20.04)
- [ ] macOS (Intel)
- [ ] macOS (Apple Silicon)
- [ ] Windows (WSL2)
- [ ] Docker container

**Go Versions:**

- [ ] Go 1.19
- [ ] Go 1.20
- [ ] Go 1.21

**Python Versions:**

- [ ] Python 3.8
- [ ] Python 3.9
- [ ] Python 3.10
- [ ] Python 3.11

## 🌍 Multi-language Impact

<!-- If your changes affect multi-language support -->

- [ ] No impact on multi-language support
- [ ] Tested with English models only
- [ ] Tested with multiple language models
- [ ] May affect non-English language processing (explain)

**Languages Tested:**
<!-- List language models tested if applicable -->

## 📦 Dependencies

<!-- Describe any dependency changes -->

- [ ] No new dependencies
- [ ] New Go dependencies added (list below)
- [ ] New system dependencies required (list below)
- [ ] Python dependencies updated (list below)

**New Dependencies:**
<!-- List any new dependencies and justify their inclusion -->

## 🚀 Deployment Considerations

<!-- Any special considerations for deployment -->

- [ ] No special deployment requirements
- [ ] Requires database migration
- [ ] Requires configuration changes
- [ ] Requires system-level changes
- [ ] Requires documentation updates

## 🎯 Reviewer Notes

<!-- Any specific areas you'd like reviewers to focus on -->

**Areas for Special Attention:**
<!-- Highlight complex logic, performance-critical sections, or areas you're uncertain about -->

**Questions for Reviewers:**
<!-- Any specific questions or concerns you have -->

## 📸 Screenshots/Examples

<!-- If applicable, add screenshots or example outputs -->

<!--
Example:
![Before](url-to-before-image)
![After](url-to-after-image)
-->

## ✅ Final Checklist

<!-- Verify before submitting -->

- [ ] PR title is clear and descriptive
- [ ] PR description explains the change and why it's needed
- [ ] All tests pass locally
- [ ] Code has been reviewed by the author
- [ ] Documentation is updated
- [ ] No merge conflicts with target branch
- [ ] Commits are squashed appropriately
- [ ] Commit messages follow conventional format

---

<!--
Additional Notes for Reviewers:
- Please check that the CI pipeline passes before approving
- For breaking changes, ensure migration path is clear
- For performance changes, verify benchmarks
- For new features, ensure adequate test coverage
-->

**Thank you for contributing to Go-Spacy! 🙏**
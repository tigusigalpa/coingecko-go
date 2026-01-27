# Contributing to CoinGecko Go Client

First off, thank you for considering contributing to the CoinGecko Go Client! It's people like you that make this project better.

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the existing issues to avoid duplicates. When you create a bug report, include as many details as possible:

- **Use a clear and descriptive title**
- **Describe the exact steps to reproduce the problem**
- **Provide specific examples**
- **Describe the behavior you observed and what you expected**
- **Include Go version and OS information**

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, include:

- **Use a clear and descriptive title**
- **Provide a detailed description of the suggested enhancement**
- **Explain why this enhancement would be useful**
- **List any similar features in other projects**

### Pull Requests

1. Fork the repository
2. Create a new branch from `main`:
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. Make your changes
4. Run tests:
   ```bash
   go test ./...
   ```
5. Run formatting:
   ```bash
   go fmt ./...
   ```
6. Run linting:
   ```bash
   go vet ./...
   ```
7. Commit your changes:
   ```bash
   git commit -m 'Add some amazing feature'
   ```
8. Push to your fork:
   ```bash
   git push origin feature/amazing-feature
   ```
9. Open a Pull Request

## Development Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/tigusigalpa/coingecko-go.git
   cd coingecko-go
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Run tests:**
   ```bash
   go test ./...
   ```

## Coding Standards

- Follow standard Go conventions and idioms
- Use `gofmt` for formatting
- Write clear, concise comments
- Add tests for new features
- Keep functions small and focused
- Use meaningful variable and function names

## Testing

- Write unit tests for all new code
- Ensure all tests pass before submitting PR
- Aim for high test coverage
- Use table-driven tests where appropriate

## Documentation

- Update README.md if adding new features
- Add code comments for exported functions
- Include examples for complex functionality
- Update CHANGELOG.md with your changes

## Commit Messages

- Use present tense ("Add feature" not "Added feature")
- Use imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit first line to 72 characters
- Reference issues and pull requests liberally

## Questions?

Feel free to open an issue with your question or contact the maintainers directly.

Thank you for contributing! 🎉

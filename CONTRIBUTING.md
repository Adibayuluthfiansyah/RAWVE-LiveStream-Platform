# Contributing to RAWVE

Thank you for your interest in contributing to **RAWVE**.

RAWVE is an open-source project built to create a fair and creator-first live streaming platform. Contributions from developers around the world are welcome and appreciated.

---

## Ways to Contribute

There are many ways you can help improve RAWVE:

* Reporting bugs
* Suggesting new features
* Improving documentation
* Submitting pull requests
* Reviewing code
* Improving performance or architecture
* Writing tests

Even small improvements are valuable.

---

## Getting Started

1. Fork the repository.

2. Clone your fork to your local machine.

```bash
git clone https://github.com/YOUR_USERNAME/RAWVE-LiveStream-Platform.git
```

3. Move into the project directory.

```bash
cd RAWVE-LiveStream-Platform
```

4. Add the upstream repository.

```bash
git remote add upstream https://github.com/Adibayuluthfiansyah/RAWVE-LiveStream-Platform.git
```

5. Create a new branch for your contribution.

```bash
git checkout -b feature/your-feature-name
```

---

## Development Setup

Make sure the following tools are installed:

* Go 1.25+
* Docker
* Docker Compose

Start the database:

```bash
docker-compose up -d
```

Run the application:

```bash
go run cmd/api/main.go
```

The server will start at:

```
http://localhost:8080
```

---

## Commit Convention

RAWVE follows the **Conventional Commits** standard.

Examples:

```
feat: add stream analytics endpoint
fix: resolve websocket reconnect issue
docs: update installation guide
refactor: simplify authentication middleware
test: add unit tests for user usecase
chore: update dependencies
```

| Prefix   | Description           |
| -------- | --------------------- |
| feat     | new feature           |
| fix      | bug fix               |
| docs     | documentation changes |
| refactor | code refactoring      |
| test     | adding tests          |
| chore    | maintenance tasks     |

---

## Code Guidelines

Please follow these guidelines when contributing:

* Follow Go best practices and idiomatic Go code
* Maintain **Clean Architecture separation**
* Use clear and meaningful function names
* Write proper error handling
* Avoid unnecessary dependencies
* Keep functions small and focused
* Document complex logic when needed

Directory responsibilities:

* `domain` → entities and interfaces
* `usecase` → business logic
* `repository` → database or external services
* `delivery` → HTTP handlers and WebSocket layer

---

## Pull Request Process

Before submitting a pull request:

1. Ensure the project builds successfully.
2. Make sure your changes are focused and minimal.
3. Write a clear description of what your PR does.
4. Reference related issues if applicable.

Steps to submit:

```
git add .
git commit -m "feat: add your feature"
git push origin feature/your-feature-name
```

Then open a **Pull Request** on GitHub.

---

## Reporting Bugs

If you find a bug, please open an issue and include:

* Description of the problem
* Steps to reproduce
* Expected behavior
* Screenshots or logs if applicable

GitHub Issues:
https://github.com/Adibayuluthfiansyah/RAWVE-LiveStream-Platform/issues

---

## Feature Requests

Feature ideas are welcome.

Please open an issue and describe:

* The problem you want to solve
* Your proposed solution
* Possible alternatives

---

## Community Guidelines

Please be respectful and constructive when interacting with other contributors.

We aim to create a positive and welcoming open-source community.

By participating in this project, you agree to follow the **Code of Conduct**.

---

## Maintainer

RAWVE is maintained by:

Adibayu Luthfiansyah

GitHub
https://github.com/adibayuluthfiansyah

---

Thank you for helping make RAWVE better.

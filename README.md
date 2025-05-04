# ZenithPlans Golang Packages

- [ZenithPlans Golang Packages](#zenithplans-golang-packages)
  - [📦 Available Packages](#📦-available-packages)
  - [🧱 Architecture Philosophy](#🧱-architecture-philosophy)
  - [🔧 Getting Started](#🔧-getting-started)
    - [Release Versioning](#release-versioning)
    - [Usage](#usage)
      - [🧪 Pre-release (Unstable)](#🧪-pre-release-unstable)
      - [✅ Stable](#stable)

A collection of modular Go packages powering the core infrastructure of the [ZenithPlans](https://github.com/zenithplans) ecosystem — designed for performance, composability, and clarity.

## 📦 Available Packages

| Package              | Description                                                                               |
| -------------------- | ----------------------------------------------------------------------------------------- |
| [`env`](./env)       | Lightweight environment variable loader with safe fallbacks and strict mode (`MustGet`)   |
| [`chain`](./chain)   | Middleware chaining utility for `http.HandlerFunc`                                        |
| [`logger`](./logger) | Structured logging wrapper over `log/slog` with context propagation support               |
| [`pg`](./pg)         | Configurable PostgreSQL client using `pgxpool`, with lifecycle control and health checks  |
| [`http`](./http)     | Thin abstraction over Go’s `http.Server` with configurable timeouts and graceful shutdown |

## 🧱 Architecture Philosophy

- ✅ **Explicit, not magical** – no hidden behavior or global state.
- 🧩 **Composable** – designed to plug into event-driven and RESTful architectures.
- 🛡️ **Minimal dependencies** – favors the standard library and well-maintained community packages.

## 🔧 Getting Started

### Release Versioning

We follow [Semantic Versioning](https://semver.org/) for all public tags:

- `v1.1.0-beta.1`
  Pre-release version — for testing upcoming features.

- `v1.0.0`
  First stable release, suitable for production use.

### Usage

#### 🧪 Pre-release (Unstable)

To use the latest unreleased features (from the `release` branch):

```bash
go get github.com/zenithplans/pkg/pg@release

# or to target a specific beta version:
go get github.com/zenithplans/pkg/pg@v1.2.0-beta.1
```

#### ✅ Stable

To install a stable version:

```bash
go get github.com/zenithplans/pkg/pg

# or explicitly pin a version:
go get github.com/zenithplans/pkg/pg@v1.0.0
```

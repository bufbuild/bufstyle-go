# bufstyle-go

[![Build](https://github.com/bufbuild/bufstyle-go/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/bufbuild/bufstyle-go/actions/workflows/ci.yaml)
[![Report Card](https://goreportcard.com/badge/buf.build/go/bufstyle)](https://goreportcard.com/report/buf.build/go/bufstyle)
[![GoDoc](https://pkg.go.dev/badge/buf.build/go/bufstyle.svg)](https://pkg.go.dev/buf.build/go/bufstyle)
[![Slack](https://img.shields.io/badge/slack-buf-%23e01563)](https://buf.build/links/slack)

A linter that enforces Buf-specific Go style.

### Usage

```
go install buf.build/go/bufstyle
bufstyle ./...
```

To automatically fix issues that `bufstyle` knows how to fix:

```
bufstyle -fix ./...
```

A `.bufstyle.yaml` file can be added in the working directory to disable and ignore analyzers:

```yaml
version: v1
disable:
  - BEHAVIOUR
ignore:
  PACKAGE_FILENAME:
    - internal/foo/foo.go
    - internal/pkg
```

## Status: Alpha

This will never be generally available and is not stable. By design, you should not use this unless you are developing libraries within `github.com/bufbuild` or `buf.build/go`.

## Legal

Offered under the [Apache 2 license](https://github.com/bufbuild/bufstyle-go/blob/main/LICENSE).

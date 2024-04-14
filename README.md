# twu

<!-- [![Keep a Changelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735)](CHANGELOG.md)
[![GitHub Release](https://img.shields.io/github/v/release/golang-templates/seed)](https://github.com/golang-templates/seed/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/golang-templates/seed.svg)](https://pkg.go.dev/github.com/golang-templates/seed)
[![go.mod](https://img.shields.io/github/go-mod/go-version/golang-templates/seed)](go.mod)
[![LICENSE](https://img.shields.io/github/license/golang-templates/seed)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/golang-templates/seed/build.yml?branch=main)](https://github.com/golang-templates/seed/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-templates/seed)](https://goreportcard.com/report/github.com/golang-templates/seed)
[![Codecov](https://codecov.io/gh/golang-templates/seed/branch/main/graph/badge.svg)](https://codecov.io/gh/golang-templates/seed)

⭐ `Star` this repository if you find it valuable and worth maintaining.

👁 `Watch` this repository to get notified about new releases, issues, etc. -->

`twu` consumes the [OpenWeatherMap API](https://openweathermap.org/current) to provide weather information for a given city.

It's built with Go and uses [HTMX](https://htmx.org/) to make the application interactive without writing _any_ JavaScript.

## Quickstart

* To run the application, execute the following commands:

    ```bash
    # run the application
    go run main.go

    # quit the application
    Ctrl + C
    ```

* Once it's running, you can access the application at `http://localhost:8080`.
* To build the application, execute the following commands:

    ```bash
    # build the application
    go build -o ./build/twu .
    ```

* To run the built application, execute the following commands:

    ```bash
    # run the application
    ./build/twu
    ```

* To run the tests, execute the following commands:

    ```bash
    # run the tests
    go test ./...
    ```

## Setup

1. Install [Go](https://golang.org/doc/install).
2. Install [Visual Studio Code](https://code.visualstudio.com/).
3. Install [Go extension](https://code.visualstudio.com/docs/languages/go).
4. Clone and open this repository.
5. `F1` -> `Go: Install/Update Tools` -> (select all) -> OK.

## Build

### Terminal

* `make`: execute the build pipeline.
* `make help`: print help for the [Make targets](Makefile).

### Visual Studio Code

`F1` → `Tasks: Run Build Task (Ctrl+Shift+B or ⇧⌘B)` to execute the build pipeline.

## Release

The release workflow is triggered each time a tag with `v` prefix is pushed.

_CAUTION_: Make sure to understand the consequences before you bump the major version.
More info: [Go Wiki](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher),
[Go Blog](https://blog.golang.org/v2-go-modules).

## Maintenance

Notable files:

* [.github/workflows](.github/workflows) - GitHub Actions workflows,
* [.github/dependabot.yml](.github/dependabot.yml) - Dependabot configuration,
* [.vscode](.vscode) - Visual Studio Code configuration files,
* [.golangci.yml](.golangci.yml) - golangci-lint configuration,
* [.goreleaser.yml](.goreleaser.yml) - GoReleaser configuration,
* [Dockerfile](Dockerfile) - Dockerfile used by GoReleaser to create a container image,
* [Makefile](Makefile) - Make targets used for development, [CI build](.github/workflows) and [.vscode/tasks.json](.vscode/tasks.json),
* [go.mod](go.mod) - [Go module definition](https://github.com/golang/go/wiki/Modules#gomod),
* [tools.go](tools.go) - [build tools](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module).

## Contributing

Feel free to create an issue or propose a pull request.

Follow the [Code of Conduct](CODE_OF_CONDUCT.md).

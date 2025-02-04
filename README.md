# volgo

[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/elliot40404/volgo/release.yml)](https://github.com/elliot40404/volgo/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/elliot40404/volgo)](https://goreportcard.com/report/github.com/elliot40404/volgo)
[![Go Reference](https://pkg.go.dev/badge/github.com/elliot40404/volgo.svg)](https://pkg.go.dev/github.com/elliot40404/volgo)


![logo](./images/logo.png)
![demo](./images/demo.gif)

volgo is a simple cross platform cli app that can be used to manage the volume of your system audio.

## Installation

```bash
go install github.com/elliot40404/volgo/cmd/volgo@latest
```

## Usage 

```bash
volgo # to start in interactive mode

volgo help # to see all available commands
```

## Features

- [x] Interactive mode
- [x] Non interactive mode

## Build From Source with alternative engines

```bash
git clone https://github.com/elliot40404/volgo.git
cd volgo
go build -o volgo cmd/volgo/
```

## License

MIT

## Support My Work

<a href="https://ko-fi.com/elliot40404">
<img src="https://storage.ko-fi.com/cdn/brandasset/v2/support_me_on_kofi_red.png" alt="Support Me on Ko-fi" width="200">
</a>

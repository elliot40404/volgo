set windows-shell := ["pwsh.exe", "-NoLogo", "-Command"]

default: build

build_cmd := if os() == "windows" { "go build -o ./bin/volgo.exe ./cmd/volgo/" } else { "go build -o ./bin/volgo ./cmd/volgo/" }

build: clean lint
    {{build_cmd}}

run iter='' cron='':
    go run ./cmd/volgo/ {{iter}} "{{cron}}"

exec iter='' cron='':
    ./bin/volgo {{iter}} "{{cron}}"


install:
    go install ./cmd/volgo/

build-run: build exec

rmcmd := if os() == "windows" { "mkdir ./bin -Force; Remove-Item -Recurse -Force ./bin" } else { "rm -rf ./bin" }

clean:
    {{rmcmd}}

lint:
    golangci-lint run

lint-fix:
    golangci-lint run --fix

vendor:
    go mod tidy
    go mod vendor
    go mod tidy

release:
    goreleaser release --snapshot --clean

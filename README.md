# ☕ Coffee Beans

A small file storage server that supports a simple GET/PUT based protocol

Originally created to provide a tiny private maven repository.

## Usage

1. Create a basic `coffee-beans.yaml` with
    ```yaml
    repositories:
      - libs-release
    ```
2. Start `coffee-beans` to start an instance with a single repo called `libs-release`
serving files from `data` and listening at http://localhost:8080/repo/libs-release

You can use Gradle or Maven publish tasks to upload to this repository.

Configure systemd to run coffee-beans as a service.

## Complete Configuration

```yaml
content_root:  'data' # File serving area

server: # Server configuration
  listen_address: ''
  listen_port: 8080

repositories:   # Add your repositories here
  - libs-release

```

## Building

Requires Go 1.13+

- Run `go build` to build the project
- Run `go test` to run the tests

##  TODO

- [x] Multiple repositories
- [ ] Simple authentication
- [ ] Remote backends
- [ ] Proxying
- [ ] Virtual Repositories
- [ ] API
- [ ] UI




# ☕ Coffee Beans

A small file storage server that supports a simple GET/PUT based protocol

Originally created to provide a tiny private maven repository.

## Usage

Start `coffee-beans` to start an instance with a single repo called `repository`
serving files from `data` and listening on http://localhost:8080/repository

You can use Gradle or Maven publish tasks to upload to this repository.

Configure systemd to run coffee-beans as a service.

## Building

Requires Go 1.13+

- Run `go build` to build the project
- Run `go test` to run the tests

##  TODO

- [ ] Multiple repositories
- [ ] Simple authentication
- [ ] Remote backends
- [ ] Proxying
- [ ] Virtual Repositories
- [ ] API
- [ ] UI




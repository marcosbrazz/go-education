# go-education

## Starting a new app

**Create module**

`go mod init <module name>`

**Import external packages**

`go get <package name>`

*Example:* `go get github.com/urfave/cli`

**Build app to a compiled file**

`go build`

## Tests

- Files must end with `_test`.
- Functions must start with `Test` in name.
- Package may have `_test` in the end.
  
| Command                               | Description                                     |
| ------------------------------------- | ----------------------------------------------- |
| `go test`                             | Look for test files in current dir              |
| `go test ./...`                       | Look for test files in all project packages     |
| `go test -v`                          | Verbose                                         |
| `go test --cover`                     | Coverage summary                                |
| `go test --coverprofile coverage.txt` | Generate coverage report file                   |
| `go tool cover --func=coverage.txt`   | Read coverage report file and print in terminal |
| `go tool cover --html=coverage.txt`   | Generate an HTML files with lines not covered   |

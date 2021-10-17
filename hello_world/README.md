# Hello World

## Modules
It's required that your code live in some kind of module as of Go 1.16 and later. To generate your module, you need to type `go mod init hello`, (in this case, `hello` is the module name).

## Running Tests
Instead of being kept in a separate folder, unit tests tend to live right next to the source file to which they apply. You'll want (minimally) to `import "testing"`.

When you're green, running tests looks like:
```shell
❯ go test
PASS
ok      hello   0.607s
```

## Writing Tests
- files must be named `xxx_test.go`
- Test function must start with the word Test and take exactly one argument, `t *testing.T`
- requires that you `import "testing"`
  
## Other fun stuff
- browse a living document server by running `godoc -http :8000` (to run it on port 8000). Needs a manual install lately - but you get it with `go get golang.org/x/tools/cmd/godoc` anymore. Try `go install golang.org/x/tools/cmd/godoc@latest` instead.  **This was absolutetly no working on my install and it's killing me, figure it out later.**



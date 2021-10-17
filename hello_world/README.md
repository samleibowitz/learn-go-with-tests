# 1.0 Hello World

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
- browse a living document server by running `godoc -http=localhost:6060` (to run it on port 6060). Needs a manual install lately - but you can't get it with `go get golang.org/x/tools/cmd/godoc` anymore. Try `go install golang.org/x/tools/cmd/godoc@latest` instead.
  
IMPORTANT NOTE: The official Go installer doesn't set `GOPATH` or add it to your command path. I edited my `.zshrc` as follows:
```shell
# Add GOPATH
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

## 1.1 Hello, YOU

Note the way that `go test` reports errors in this case.

```
❯ go test
# hello [hello.test]
./hello_test.go:6:14: too many arguments in call to Hello
        have (string)
        want ()
FAIL    hello [build failed]
```

## 1.2 Hello world... again

Notice that we're not making use of a default parameter here, just an empty string.

Tests in Go can contain sub-tests. Since the subtests can vary on what parameters are passed, there may be a way to express something similar to RSpec's contexts here.

## 1.3 Keep Going! More requirements

Nothing too complicated here, just adding another language...
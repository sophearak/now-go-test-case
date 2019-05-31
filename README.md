# now-go-test-case
A test case for Now's Go builder

## To test

1. Clone this repo
2. run `now dev` inside the directory
3. visit `localhost:3000/test-1`

## Expected result

You should recieve a plain text response of `test-1 response`

## Actual result

```
Command failed: go build -ldflags -s -w -o /var/folders/81/r813w3c177lg_37bld5vb35m0000gn/T/f59b73f/handler /Users/danielerickson/dev/now-go-test-case/.now/cache/index/src/lambda/main__mod__.go
```

And in the console you'll see this output:

```
~/dev/now-go-test-case/ master * > now dev
> Now CLI 15.3.0 dev (beta) — https://zeit.co/feedback/dev
> Ready! Available at http://localhost:3000
> GET /test-1
> Building @now/go@canary:api/test-1/index.go
Downloading user files...
Parsing AST for "api/test-1/index.go"
Found exported function "Handler" in "api/test-1/index.go"
Tidy `go.mod` file...
go: finding github.com/techwraith/now-go-test-case/api latest
go: finding github.com/zeit/now-builders/utils/go/bridge latest
go: finding github.com/zeit/now-builders/utils/go latest
go: finding github.com/zeit/now-builders/utils latest
go: finding github.com/zeit/now-builders latest
go: finding github.com/aws/aws-lambda-go/events latest
go: finding github.com/aws/aws-lambda-go/lambda latest
Running `go build`...
go: finding github.com/techwraith/now-go-test-case/api latest
build command-line-arguments: cannot load github.com/techwraith/now-go-test-case/api: cannot find module providing package github.com/techwraith/now-go-test-case/api
failed to `go build`
{ Error: Command failed: go build -ldflags -s -w -o /var/folders/81/r813w3c177lg_37bld5vb35m0000gn/T/1df1c4e1/handler /Users/danielerickson/dev/now-go-test-case/.now/cache/index/src/lambda/main__mod__
.go
    at makeError (/Users/danielerickson/Library/Caches/co.zeit.now/dev/builders/node_modules/execa/index.js:174:9)
    at Promise.all.then.arr (/Users/danielerickson/Library/Caches/co.zeit.now/dev/builders/node_modules/execa/index.js:278:16)
    at <anonymous>
    at process._tickCallback (internal/process/next_tick.js:188:7)
  message:
   'Command failed: go build -ldflags -s -w -o /var/folders/81/r813w3c177lg_37bld5vb35m0000gn/T/1df1c4e1/handler /Users/danielerickson/dev/now-go-test-case/.now/cache/index/src/lambda/main__mod__.go',
  code: 1,
  stdout: null,
  stderr: null,
  failed: true,
  signal: null,
  cmd:
   'go build -ldflags -s -w -o /var/folders/81/r813w3c177lg_37bld5vb35m0000gn/T/1df1c4e1/handler /Users/danielerickson/dev/now-go-test-case/.now/cache/index/src/lambda/main__mod__.go',
  timedOut: false,
  killed: false }
```

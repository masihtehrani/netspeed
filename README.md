# speedtest
this project provide two different implementation of two net speed test services.

## Run Project

just you can run this command `go run ./cmd/...`

## Sample HTTP request & response

send request with `Fast.com`:

`curl --location --request GET http://localhost:6010/fastcom`

send request with `speedtest`:

`curl --location --request GET http://localhost:6010/speedtest`

response:

`{"download":"10.787452 Mbps","upload":"0.000000 Mbps"}`

## structure

it project have several part

### CMD

run project contain `main` function

### Server

provide HTTP server and routing system

### Contracts

structure providers just `interface` in root project

### Providers

contain `new` file for make various provider.

### Fastcom

it's one of the providers it made with external libraries. I add a extra interface in `fastcom.go` for I can write test for my package `fastcom`.

### Speedtest

it's one of the providers it made with external libraries

## Lint
code linted with Golangci_lint version v1.44.0

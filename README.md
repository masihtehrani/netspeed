# speedtest

## Run Project

just you can run this command `go run ./cmd/...`

## Sample HTTP request & response

send request with `Fast.com`:

`curl --location --request GET http://localhost:6010/fastcom`

send request with `speedtest`:

`curl --location --request GET http://localhost:6010/speedtest`

response:

`{"download":"10.787452 Mbps","upload":"0.000000 Mbps"}`

## Lint
code linted with Golangci_lint version v1.44.0

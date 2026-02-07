test:
	go clean -testcache
	go test ./...

lint:
	golangci-lint run

run_all_tests:
	go test ./...

lint:
	golangci-lint run
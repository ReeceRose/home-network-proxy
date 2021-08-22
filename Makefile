test:
	go test -coverprofile=coverage.out -race -covermode=atomic ./...
	go tool cover -html=coverage.out

test-go:
	go test -v ./...

run-db:
	docker-compose -f docker-compose-db.yml up

run-reporting-tool:
	go run cmd/reporting-tool/main.go

run-api:
	go run cmd/api/main.go

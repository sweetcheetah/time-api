test:
	@go test -v -cover ./...
build:
	@go build -o time-api main.go
run:
	@PORT=8080 go run main.go
docker-build:
	@docker build -t time-api:latest .
docker-run:
	@docker run -p 8080:8080 --rm time-api:latest

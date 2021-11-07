GOLANG_CI_LINT_IMAGE=golangci/golangci-lint:latest-alpine

go-build:
	GOOS=linux CGO_ENABLED=0 go build -o users_bin ./cmd/main.go

docker-build:
	docker build --no-cache -t lambda-users:0.0.1 .

run:
	docker run --env-file local.env --rm -p 9000:8080 lambda-users:0.0.1

up: go-build docker-build run

test-call:
	curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'

go-lint:
	docker run -v ${PWD}:/app -w /app $(GOLANG_CI_LINT_IMAGE) golangci-lint run -v --timeout 600m --fix

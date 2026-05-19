.PHONY: web api test lint build format

web:
	cd apps/web && npm run dev

api:
	cd apps/api && go run ./cmd/server

test:
	cd apps/api && go test ./...
	cd apps/web && npm run test -- --run

lint:
	cd apps/web && npm run lint
	cd apps/api && go vet ./...

format:
	cd apps/api && gofmt -w .
	cd apps/web && npm run lint -- --fix

build:
	cd apps/web && npm run build
	cd apps/api && go build ./cmd/server

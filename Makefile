build:
	docker-compose build

run-app:
	docker-compose up dev-app-orders

run-payment-status-update-worker:
	docker-compose up dev-app-payment-status-update-worker

run-order-status-update-worker:
	docker-compose up dev-app-order-status-update-worker

run-migration:
	docker-compose run dev-app-orders go run ./cmd/migration/main.go

run-prod:
	docker-compose up prod-app-orders

start-infra:
	docker-compose -f docker-compose-infra.yaml up

swagger:
	docker-compose run dev-app-orders swag init -g internal/external/handler/http_server/app.go -o doc/swagger/

mocks:
	docker-compose run dev-app-orders go generate ./...

test:
	docker-compose run dev-app-orders go test ./...

test-coverage:
	docker-compose run dev-app go test -coverprofile cover.out `go list ./... | egrep -v '(/cmd|/src/external/database|src/external/handler|/api|/doc|/infra|/src/pkg/uuid/mock|/src/application/contract/mock|/src/external/handler/http_server|/api)$\'` && go tool cover -html=cover.out

get-coverage:
	docker-compose run -d dev-app go test -coverprofile cover.out `go list ./... | egrep -v '(/cmd|/src/external/database|src/external/handler|/api|/doc|/infra|/src/pkg/uuid/mock|/src/application/contract/mock|/src/external/handler/http_server|/api)$\')` && go tool cover -func cover.out | fgrep total | awk '{print substr($$3, 1, length($$3)-1)}'

run-test:
	$(MAKE) mocks && $(MAKE) test

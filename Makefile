build:
	docker-compose build

run-dev:
	docker-compose up

run-prod:
	docker-compose -f docker-compose.prod.yaml up

mocks:
	docker-compose run dev-app go generate ./...

test:
	docker-compose run dev-app go test ./...

test-coverage:
	docker-compose run dev-app go test -coverprofile cover.out `go list ./... | egrep -v '(/cmd|/src/external/database|src/external/handler|/api|/doc|/infra|/src/pkg/uuid/mock|/src/application/contract/mock|/src/external/handler/http_server|/api)$\'` && go tool cover -html=cover.out

get-coverage:
	docker-compose run -d dev-app go test -coverprofile cover.out `go list ./... | egrep -v '(/cmd|/src/external/database|src/external/handler|/api|/doc|/infra|/src/pkg/uuid/mock|/src/application/contract/mock|/src/external/handler/http_server|/api)$\')` && go tool cover -func cover.out | fgrep total | awk '{print substr($$3, 1, length($$3)-1)}'

run-test:
	$(MAKE) mocks && $(MAKE) test

.PHONY: fmt clean deep-clean test test-update test-race help
.DEFAULT: server

TEST := CONFIG_ENV=test go test ./...

VERSION ?= $(shell git describe --abbrev=0 --tags)

#==============================
# App hygiene
#==============================
update: ## go mod tidy, then go get -u -d
	go mod tidy
	go get -u -d ./...

clean: ## go mod tidy, lint, go generate
	go mod tidy
	go generate ./...

deep-clean: clean ## Run clean, then purge modcache
	go clean -modcache -cache -i -r -x


#==============================
# Tests
#==============================
test: ## go vet, then run tests
	go vet ./...
	$(TEST)

test-update: ## test and update snapshots
	UPDATE_SNAPSHOTS=true $(TEST)

test-race: ## Run go vet, then run tests trying to catch race conditions
	go vet ./...
	CONFIG_ENV=test go test -race ./...


#==============================
# Meta
#==============================
help: ## Print help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

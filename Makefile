.PHONY: fmt clean deep-clean test test-update test-race help
.DEFAULT: server

BINARIES := $(find cmd -maxdepth 1 -mindepth 1 -type d -exec basename {} \;)
TEST := CONFIG_ENV=test go test ./...
CONTAINER := insider

VERSION ?= $(shell git describe --abbrev=0 --tags)
BUILD_FLAGS := 

ifneq (,$(wildcard ./vendor))
	$(warning Found vendor directory; setting "-mod vendor" to any "go build" commands)
	BUILD_FLAGS += -mod vendor
endif

PORT ?= 7071

#==============================
# Builds
#==============================
$(BINARIES): ## Create the server's binary -> ./bin/server
	$(BASE_FLAGS) go build $(BUILD_FLAGS) -ldflags="-X 'github.com/AnthonyHewins/insider/cmd/root.version=$(VERSION)'" -o ./bin/$@ ./...

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

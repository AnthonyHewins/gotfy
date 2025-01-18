.PHONY: fmt clean deep-clean test help

#==============================
# App hygiene
#==============================
update: clean ## go mod tidy, then go get -u -d
	go get -u ./...

clean: ## go mod tidy
	go mod tidy

deep-clean: clean ## Run clean, then purge modcache
	go clean -modcache -cache -i -r -x


#==============================
# Tests
#==============================
test: ## go vet, then run tests
	go vet ./...
	go test ./...


#==============================
# Meta
#==============================
help: ## Print help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

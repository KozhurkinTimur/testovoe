SHELL            := /bin/sh
GOBIN            ?= $(GOPATH)/bin
PATH             := $(GOBIN):$(PATH)
GO               = go

M                = $(shell printf "\033[34;1m>>\033[0m")
TARGET_DIR       ?= $(PWD)/.build
MIGRATIONS_DIR   = ./cmd/migrations/

.PHONY: start
start: ## Build and run docker-compose
	docker-compose build
	docker-compose up

.PHONY: upgrade
upgrade: ## Dependencies upgrade
	GOWORK=off go-mod-upgrade
	go mod tidy
	go mod vendor

.PHONY: build
build:
	$(info $(M) building entity center...)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GCFLAGS) $(LDFLAGS) -o $(TARGET_DIR)/service ./cmd/server/main.go

.PHONY: watch
watch: ## Run binaries that rebuild themselves on changes
	$(info $(M) run...)
	@$(GOBIN)/air -c $(PWD)/.air.conf

.PHONY: install-tools
install-tools: $(GOBIN) # Install tools needed for development
	@GOBIN=$(GOBIN) $(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@GOBIN=$(GOBIN) $(GO) install github.com/cosmtrek/air@latest

.PHONY: db-create-migration
db-create-migration: ## Create a new database migration file: db-create-migration {migration-name}
	$(info $(M) creating DB migration...)
	migrate create -ext sql -dir "$(MIGRATIONS_DIR)" $(filter-out $@,$(MAKECMDGOALS))

%:
	@:
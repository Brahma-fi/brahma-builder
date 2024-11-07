.EXPORT_ALL_VARIABLES:

include .env.local

PLUGIN_NAME := vault-eth-signer
VAULT_PATH := ./_scripts/vault
GO := go
service_name := brahma-builder
GOBIN	:= $(PWD)/_bin
MIGRATIONS_DIR = ./migrations/
GOPRIVATE      =github.com/Brahma-fi

OS="$(shell go env var GOOS | xargs)"

path :=$(if $(path), $(path), "./")
ecr='488892008441.dkr.ecr.us-east-2.amazonaws.com'


.PHONY: gen-abis setup-local-vault setup-local-plugin start-local-vault download-plugin

gen-abis:
	mkdir -p ./pkg/utils/abis/executorplugin
	abigen --abi=./pkg/utils/abis/json/executor_plugin.json \
		--pkg=utils \
		--type=executorplugin \
		--out=./pkg/utils/abis/executorplugin/binding.go
	mkdir -p ./pkg/utils/abis/bundler
	abigen --abi=./pkg/utils/abis/json/bundler.json \
		--pkg=utils \
		--type=bundler \
		--out=./pkg/utils/abis/bundler/binding.go
	mkdir -p ./pkg/utils/abis/permit2
	abigen --abi=./pkg/utils/abis/json/permit2.json \
		--pkg=utils \
		--type=permit2 \
		--out=./pkg/utils/abis/permit2/binding.go
	mkdir -p ./pkg/utils/abis/weth
	abigen --abi=./pkg/utils/abis/json/weth.json \
		--pkg=utils \
		--type=weth \
		--out=./pkg/utils/abis/weth/binding.go

setup-local-vault:
	@sh ./_scripts/vault/setup_vault.sh $(VAULT_PATH)

setup-local-plugin:
	@sh ./_scripts/vault/setup_plugin.sh $(VAULT_PATH)/plugins $(PLUGIN_NAME) $(EXECUTOR_KEY)

start-local-vault:
	vault server -dev \
		-dev-listen-address="0.0.0.0:8200" \
		-dev-plugin-dir=$(VAULT_PATH)/plugins \
		-dev-root-token-id="root"

download-plugin:
	@mkdir -p $(VAULT_PATH)/plugins
	@if [ ! -f $(VAULT_PATH)/plugins/$(PLUGIN_NAME) ] && [ ! -f $(VAULT_PATH)/$(PLUGIN_NAME)/_bin/debug/$(OS)/$(PLUGIN_NAME) ]; then \
		if [ ! -d $(VAULT_PATH)/$(PLUGIN_NAME) ]; then \
			git clone git@github.com:Brahma-fi/$(PLUGIN_NAME).git $(VAULT_PATH)/$(PLUGIN_NAME); \
		fi; \
		cd $(VAULT_PATH)/$(PLUGIN_NAME) && make build; \
	else \
		echo "$(PLUGIN_NAME) plugin already exists. Skipping download and build."; \
	fi
	@if [ -f $(VAULT_PATH)/$(PLUGIN_NAME)/_bin/debug/$(OS)/$(PLUGIN_NAME) ]; then \
		mv $(VAULT_PATH)/$(PLUGIN_NAME)/_bin/debug/$(OS)/$(PLUGIN_NAME) $(VAULT_PATH)/plugins/$(PLUGIN_NAME); \
	fi

.PHONY: build-common
build-common: ## - execute build common tasks clean and mod tidy
	@ $(GO) version
	@ $(GO) clean
	@ $(GO) mod tidy && $(GO) mod download
	@ $(GO) mod verify

build: build-common ## - build a debug binary to the current platform (windows, linux or darwin(mac))
	@ echo cleaning...
	@ rm -f $(GOBIN)/debug/$(OS)/$(service_name)
	@ echo building...
	@ $(GO) build -tags dev -o "$(GOBIN)/debug/$(OS)/$(service_name)" cmd/main.go
	@ ls -lah $(GOBIN)/debug/$(OS)/$(service_name)

.PHONY: ci
ci: gen-mocks build test ## - execute all ci tasks

.PHONY: build-linux-release
build-linux-release: build-common ## - build a static release linux elf(binary)
	@ CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -tags $(tags) -ldflags='-w -s -extldflags "-static"' -a -o "_bin/release/linux/$(service_name)" cmd/main.go

.PHONY: build-local-release
build-local-release: build-common ## - build a release binary to the current platform (windows, linux or darwin(mac))
	@ echo building
	@ @ CGO_ENABLED=0 go build -tags $(tags) -ldflags='-w -s -extldflags "-static"' -a -o "_bin/release/$(OS)/$(service_name)" cmd/main.go
	@ echo "_bin/release/$(OS)/"
	@ ls -lah _bin/release/$(OS)/$(service_name)
	@ echo "done"

.PHONY: test
test: build-common ## - execute go test command
	@ go test -v -cover ./...

.PHONY: scan
scan: ## - execute static code analysis
	@ gosec -fmt=sarif -out=$(service_name).sarif -exclude-dir=test -exclude-dir=bin -severity=medium ./... | 2>&1
	@ echo ""
	@ cat $(path)/$(service_name).sarif

.PHONY: ci-lint
ci-lint: ## - runs golangci-lint
	@ golangci-lint run -v ${LINT_FLAGS}

.PHONY: ci-lint-docker
ci-lint-docker: ## - runs golangci-lint with docker container
	@ docker run --rm -v "$(shell pwd)":/app -w /app ${LINT_IMAGE} golangci-lint run ${LINT_FLAGS}

.PHONY: docker-build
docker-build: ## - build docker image
	@ docker build -f deployment/$(service_name)/Dockerfile -t $(ecr)/$(service_name):$(version) --build-arg TOKEN=$(token) .

.PHONY: docker-build-arm
docker-build-arm: ## - build docker image
	@ DOCKER_DEFAULT_PLATFORM=linux/arm64 docker build -f deployment/$(service_name)/Dockerfile -t $(ecr)/$(service_name):$(version) --build-arg DEBUG=$(debug) .

.PHONY: docker-push
docker-push:
	@ docker push $(ecr)/$(service_name):$(version)

.PHONY: docker-scan
docker-scan: ## - Scan for known vulnerabilities
	@ printf "\033[32m\xE2\x9c\x93 Scan for known vulnerabilities the smallest and secured golang docker image based on scratch\n\033[0m"
	@ $(SHELL) .scripts/docker-scan.sh $(version)

test-coverage: ## - execute go test command with coverage
	@ mkdir -p _coverage && mkdir -p _report
	@ go test -json -v -cover -covermode=atomic -coverprofile=_coverage/cover.out ./pkg/... > _report/report.out

.PHONY: sonar-scan-local
sonar-scan-local: test-coverage ## - start sonar qube locally with docker (you will need docker installed in your machine)
	@ $(SHELL) _scripts/sonar-start.sh
	@ echo "login with user: admin pwd: 1234"

.PHONY: sonar-scan
sonar-scan: test-coverage ## - execute sonar scan
	@ sonar-scanner -D sonar.projectKey="$(service_name)" \
                    -D sonar.projectName="$(service_name)" \
                    -D sonar.scm.provider=git \
                    -D sonar.sources=. \
                    -D sonar.exclusions="swagger/**,pkg/bindings/**,_repl/**,_mocks/**,pkg/database/postgresSQL_test.go,_bin/**" \
                    -D sonar.tests=pkg \
                    -D sonar.test.inclusions=pkg/**/*_test.go \
                    -D sonar.go.tests.reportPaths=_report/report.out \
                    -D sonar.go.coverage.reportPaths=_coverage/*.out \
                    -D sonar.host.url="$(sonarUrl)" \
                    -D sonar.github.repository='https://github.com/Brahma-fi/$(service_name)' \
                    -D sonar.token="$(sonarToken)"

.PHONY: sonar-stop
sonar-stop: ## - stop sonar qube docker container
	@ docker stop sonarqube

.PHONY: db-migrate
db-migrate: ## Run migrate command
	$(info $(M) running DB migrations...)
	@$(GOBIN)/migrate -path "$(MIGRATIONS_DIR)" -database "$(STORAGE_DSN)" $(filter-out $@,$(MAKECMDGOALS))

.PHONY: db-create-migration
db-create-migration: ## Create a new database migration file
	$(info $(M) creating DB migration...)
	@$(GOBIN)/migrate create -ext sql -dir "$(MIGRATIONS_DIR)" $(filter-out $@,$(MAKECMDGOALS))

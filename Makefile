.PHONY: *

all: generate

.PHONY: generate
generate: generate-plan-modifiers speakeasy

speakeasy: check-speakeasy
	speakeasy run --skip-versioning --output console --minimal
	@go mod tidy
	@go generate .
	@git clean -fd docs > /dev/null
	@git checkout -- README.md
	@rm USAGE.md

FILES=$(shell find internal/provider -type f | grep data_source | grep -v portallist | grep -v cloudgatewayprovideraccountlist)
remove-data-sources:
	@if [ -n "$(FILES)" ]; then rm $(FILES); fi
	@rm -r examples/data-sources

check-speakeasy:
	@command -v speakeasy >/dev/null 2>&1 || { echo >&2 "speakeasy CLI is not installed. Please install before continuing."; exit 1; }

OS=$(shell uname | tr "[:upper:]" "[:lower:]")
ARCH=$(shell uname -m | sed 's/aarch64/arm64/' | sed 's/x86_64/amd64/')
test:
	@cd tests/e2e; rm -rf .terraform.lock.hcl terraform.tfstate terraform.tfstate.backup .terraform local-plugins
	mkdir -p tests/e2e/local-plugins/registry.terraform.io/kong/kong-mesh/999.99.9/$(OS)_$(ARCH)
	go mod tidy
	go build -o tests/e2e/local-plugins/registry.terraform.io/kong/kong-mesh/999.99.9/$(OS)_$(ARCH)/terraform-provider-kong-mesh_v999.99.9
	@cd tests/e2e; terraform providers mirror ./local-plugins || true
	@cd tests/e2e; ls -R local-plugins; terraform init -plugin-dir ./local-plugins; terraform apply -auto-approve; terraform destroy -auto-approve


test-cleanup:
	@cd tests/e2e; rm -rf local-plugins .terraform .terraform.lock.hcl terraform.tfstate terraform.tfstate.backup

unit-test:
	go test -v ./internal/sdk/internal/...

dev/use-local-shared-speakeasy:
	go mod edit -replace=github.com/Kong/shared-speakeasy=../shared-speakeasy
	go mod tidy

acceptance:
	@TF_ACC=1 go test -count=1 -v ./tests/resources

# renovate: datasource=go depName=Kong/shared-speakeasy/resource_plan_modifier packageName=github.com/Kong/shared-speakeasy/generators/resource_plan_modifier
RESOURCE_PLAN_MODIFIER_VERSION := v0.0.15

PLAN_MOD_CMD = $(if $(wildcard go.work),go run ../shared-speakeasy/generators/resource_plan_modifier,go run github.com/Kong/shared-speakeasy/generators/resource_plan_modifier@$(RESOURCE_PLAN_MODIFIER_VERSION))

.PHONY: generate-plan-modifiers
generate-plan-modifiers:
	mkdir -p "resouce-plan-modifiers"
	cat internal/provider/mesh*_resource.go \
	| grep "Resource struct" \
	| cut -d ' ' -f 2 \
	| sed 's/Resource$$//' \
	| while read RESOURCE; do \
		LOWER=$$(echo $$RESOURCE | tr A-Z a-z); \
		SDK_NAME=$$(grep -oE 'r\.client\.[A-Za-z]+' internal/provider/$${LOWER}_resource.go | head -1 | sed 's/r\.client\.//'); \
		if grep -q 'Mesh string.*json:"mesh"' internal/provider/$${LOWER}_resource.go; then MESH_SCOPED=true; else MESH_SCOPED=false; fi; \
		$(PLAN_MOD_CMD) \
		internal/provider/$${LOWER}_resource_plan_modify.go $$RESOURCE terraform-provider-kong-mesh $$SDK_NAME $$MESH_SCOPED; \
	done

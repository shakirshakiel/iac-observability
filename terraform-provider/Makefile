TF_PROVIDERS_DIR=$(HOME)/terraform/providers
TF_PROVIDERS_PATH=$(TF_PROVIDERS_DIR)/shakiel.com/providers/observability/

deps:
	@go get ./...
	@go mod tidy

deps/update:
	@go get -u ./...

goreleaser/release:
	@goreleaser release --clean

goreleaser/snapshot:
	@goreleaser release --clean --snapshot

terraformrc-dev:
	@envsubst < .terraformrc.dev > $(HOME)/.terraformrc

terraformrc-prod:
	@envsubst < .terraformrc.prod > $(HOME)/.terraformrc

create-tf-dir:
	@mkdir -p $(TF_PROVIDERS_DIR);
	@mkdir -p $(TF_PROVIDERS_PATH);

deploy_dev: create-tf-dir terraformrc-dev
	@go build -o $(TF_PROVIDERS_DIR)/terraform-provider-observability

deploy: create-tf-dir terraformrc-prod goreleaser/snapshot
	@cp -r dist/*.zip $(TF_PROVIDERS_PATH)/;

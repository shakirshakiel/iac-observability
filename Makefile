TF_PROVIDERS_DIR=$(HOME)/terraform/providers
TF_PROVIDERS_PATH=$(TF_PROVIDERS_DIR)/shakiel.com/providers/observability/

deps:
	cd terraform-provider && go get ./...
	cd terraform-provider && go mod tidy

deps/update:
	cd terraform-provider && go get -u ./...

goreleaser/release:
	cd terraform-provider && goreleaser release --clean

goreleaser/snapshot:
	cd terraform-provider && goreleaser release --clean --snapshot

terraformrc-prod:
	cd terraform-provider && envsubst < .terraformrc.prod > $(HOME)/.terraformrc

create-tf-dir:
	@mkdir -p $(TF_PROVIDERS_DIR);
	@mkdir -p $(TF_PROVIDERS_PATH);

deploy: create-tf-dir terraformrc-prod goreleaser/snapshot
	cd terraform-provider && cp -r dist/*.zip $(TF_PROVIDERS_PATH)/;

clean:
	rm -rf terraform-provider/dist;
	rm -rf $(TF_PROVIDERS_PATH);
	rm -rf .terraform.lock.hcl;
	rm -rf .terraform;
	rm -rf terraform.tfstate;
	rm -rf terraform.tfstate.backup;

run: clean deploy
	terraform init;
	terraform apply -auto-approve;

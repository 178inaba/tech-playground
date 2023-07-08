provider_macos_path = registry.terraform.io/178inaba/miro/99.0.0/darwin_amd64/

.PHONY: install-macos

# Temp version 99.0.0
install-macos:
	@go build -o terraform-provider-miro_99.0.0
	@mkdir -p ~/Library/Application\ Support/io.terraform/plugins/$(provider_macos_path)
	@mv terraform-provider-miro_99.0.0 ~/Library/Application\ Support/io.terraform/plugins/$(provider_macos_path)

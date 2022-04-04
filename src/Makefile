check_install:
	which swagger || { echo "Installing swagger..." >&2; exit 1; } || go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger: check_install
	swagger generate spec -o ../swaggerAPI-Products.yaml --scan-models

check_install:
	which swagger || { echo "Installing swagger..." >&2; exit 1; } || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	GO111MODULE=off swagger generate spec -o ../swaggerAPI-Products.yaml --scan-models
.PHONY: start build build-frontend build-all

NOW = $(shell date -u '+%Y%m%d%I%M%S')

RELEASE_VERSION = v1.0.0

APP 		= joylivedashboard
SERVER_BIN  	= ${APP}
GIT_COUNT 		= $(shell git rev-list --all --count)
GIT_HASH        = $(shell git rev-parse --short HEAD)
RELEASE_TAG     = $(RELEASE_VERSION).$(GIT_COUNT).$(GIT_HASH)

CONFIG_DIR       = ./configs
CONFIG_FILES     = dev
STATIC_DIR       = ./frontend/dist
START_ARGS       = -d $(CONFIG_DIR) -c $(CONFIG_FILES) -s $(STATIC_DIR)

all: start

start:
	@go run -ldflags "-X main.VERSION=$(RELEASE_TAG)" main.go start $(START_ARGS)

build:
	@CGO_ENABLED=1 go build -ldflags "-w -s -X main.VERSION=$(RELEASE_TAG)" -o $(SERVER_BIN)

build-frontend:
	@echo "Building frontend..."
	@cd frontend && npm ci && npm run build:prod
	@echo "Frontend build completed."

build-all: build-frontend build
	@echo "Full build completed. Frontend: $(STATIC_DIR), Backend: $(SERVER_BIN)"

build-linux:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC="zig cc -target x86_64-linux-musl" CXX="zig c++ -target x86_64-linux-musl" CGO_CFLAGS="-D_LARGEFILE64_SOURCE" go build -ldflags "-w -s -X main.VERSION=$(RELEASE_TAG)" -o $(SERVER_BIN)_linux_amd64

# go install github.com/google/wire/cmd/wire@latest
wire:
	@wire gen ./internal/wirex

# go install github.com/swaggo/swag/cmd/swag@latest
swagger:
	@swag init --parseDependency --generalInfo ./main.go --output ./internal/swagger

# https://github.com/OpenAPITools/openapi-generator
openapi:
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/internal/swagger/swagger.yaml -g openapi -o /local/internal/swagger/v3

clean:
	rm -rf data $(SERVER_BIN) frontend/dist

serve: build-all
	./$(SERVER_BIN) start $(START_ARGS)

serve-d: build-all
	./$(SERVER_BIN) start $(START_ARGS) --daemon

stop:
	./$(SERVER_BIN) stop

gen-application-code:
	gin-admin-cli gen -d ./ -m Resource -c gen/prototype/application.yaml

docker-build:
	docker build -t $(APP):$(RELEASE_TAG) .
	docker tag $(APP):$(RELEASE_TAG) $(APP):latest

docker-push: docker-build
	docker push $(APP):$(RELEASE_TAG)
	docker push $(APP):latest
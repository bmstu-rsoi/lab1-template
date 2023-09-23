SERVICE_NAME=person
ENV=.env

.PHONY: run
run:
	docker compose -f ./docker-compose.yaml --env-file $(ENV) up -d #--build #--remove-orphans

.PHONY: run-svc
run-svc: #  make run-svc svc=redis
	docker compose -f ./docker-compose.yaml --env-file $(ENV) up -d $(svc)

.PHONY: stop
stop:
	docker compose -f ./docker-compose.yaml --env-file $(ENV) down

.PHONY: migrate-up
migrate-up:
	goose -dir "./migrations/sql/" postgres "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	goose -dir "./migrations/sql/" postgres "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable" down

.PHONY: lint
lint:
	go vet ./...
	golangci-lint run --fix

.PHONY: test
test:
	go test -v -race -timeout 90s -count=1 -shuffle=on  -coverprofile cover.out ./...
	@go tool cover -func cover.out | grep total | awk '{print $3}'
	go tool cover -html="cover.out" -o coverage.html


.PHONY: image-build
image-build:
	docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .

.PHONY: image-push
image-push:
	docker push ${IMAGE_NAME}:${IMAGE_TAG}

.PHONY: docker-login
docker-login:
	docker login -u ${REGISTRY_USER} -p ${REGISTRY_PASS}


.PHONY: mocks
mocks:
	cd internal/handler; go generate;

.PHONY: .deps
deps: .deps
.deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
	go mod download

.PHONY: build
build:
	CGO_ENABLED=0 go build -o ./bin/${SERVICE_NAME} ./cmd/main.go

.PHONY: clean
clean:
	rm bin/${SERVICE_NAME}

.PHONY: clean-all
clean-all:
	sudo docker system prune --all --volumes -f





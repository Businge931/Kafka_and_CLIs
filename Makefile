APP = kafka_and_CLIs
GOBASE = $(shell pwd)
GOBIN = $(GOBASE)/build/bin
LINT_PATH = $(GOBASE)/build/lint
TEST_PATH = $(GOBASE)/service

###  mockgen -source=producer/producer.go -destination=producer/mocks/mock_producer.go -package=mocks

####  mockgen -source=consumer/consumer.go -destination=consumer/mocks/mock_consumer.go -package=mocks


help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

deps: ## Fetch required dependencies
	go mod tidy -compat=1.22
	go mod download

send: ## Start producer instance to send messages
	go run main.go send --channel kafka-project --server localhost:9092

receive: ## Start consumer instance to receive messages
	go run main.go receive --channel kafka-project --from start --server localhost:9092 --group tests


# # # # FOR THE COMMANDS BELOW, ENSURE TO BE IN THR KAFKA DIRECTORY

partition: ## edit the number of partitions in a topic to 4
	bin/kafka-topics.sh --alter --topic jim-test --partitions 4 --bootstrap-server localhost:9092

describe-topic: ##see details about a topic
	bin/kafka-topics.sh --describe --topic kafka-project --bootstrap-server localhost:9092

verify-consumer-assignment:
	bin/kafka-consumer-groups.sh --describe --group tests --bootstrap-server localhost:9092

check-kafka-log:
	tail -f logs/server.log


run: build ## Build and run program
	$(GOBIN)/$(APP)

lint: install-golangci ## Linter for developers
	$(LINT_PATH)/golangci-lint run --timeout=5m -c .golangci.yml

lint-fix:
	$(LINT_PATH)/golangci-lint run --timeout=5m -c .golangci.yml --fix

install-golangci: ## Install the correct version of lint
	@GOBIN=$(LINT_PATH) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.1
	
run-tests: ## Run tests 
	cd $(TEST_PATH) && go test .
test-cover: ## Run tests with coverage
	cd $(TEST_PATH) && go test -cover

test-coverage: ## Run tests and generate coverage profile
	cd $(TEST_PATH) && go test -coverprofile=coverage.out

test-coverage-browser: ## Check the test coverage in the browser
	cd $(TEST_PATH) && go tool cover -html=coverage.out -o /tmp/coverage.html && wslview /tmp/coverage.html


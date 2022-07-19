BINARY_NAME=laragoApp

build:
	@go mod vendor
	@echo "Building Application..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Application Built"
run: build
	@echo "Running Application..."
	@./tmp/${BINARY_NAME} &
	@echo "Application started"

clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf tmp/${BINARY_NAME}
	@echo "Cleaned"

test:
	@echo "Testing..."
	@go test -v ./...
	@echo "Done"

start: run

stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f tmp/${BINARY_NAME}
	@echo "Stopped"

restart: stop start
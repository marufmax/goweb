BINARY_NAME=laragoApp

build:
#	@go mod vendor
	@echo "Building Application..."
	@go build -o build/${BINARY_NAME} .
	@echo "Application Built"
run: build
	@echo "Running Application..."
	@./build/${BINARY_NAME} &
	@echo "Application started"

clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf build/${BINARY_NAME}
	@echo "Cleaned"

test:
	@echo "Testing..."
	@go test -v ./...
	@echo "Done"

start: run

stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f build/${BINARY_NAME}
	@echo "Stopped"

restart: stop start
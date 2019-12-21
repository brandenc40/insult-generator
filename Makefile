build:
	@echo "Formatting.."
	@gofmt -s -w .
	@echo "Building.."
	@go build

run: build
	@echo "Starting app.."
	@go run ./main.go

clean:
	@echo "Cleaning up..."
	@rm insult-generator
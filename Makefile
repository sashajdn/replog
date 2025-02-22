.PHONY: start stop clean help test

run:
	@echo "Running the replog binary..."
	docker-compose up -d

stop:
	@echo "Stopping the replog binary..."
	docker-compose down

clean:
	@echo "Cleaning up replog docker env..."
	docker-compose down -v --remove-orphans
	-docker network prune -f

test:
	go test -v ./...

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@echo "  start  - Start the Docker Compose environment."
	@echo "  stop   - Stop the Docker Compose environment."
	@echo "  clean  - Purge Docker setup (volumes, orphan containers, networks, and build artifacts)."
	@echo "  test   - Run Go unit tests."
	@echo "  help   - Display this help message."


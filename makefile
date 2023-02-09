help:
	@echo
	@echo "Commands     Descriptions"
	@echo
	@echo "help         List all commands."
	@echo "dc-down      Stop and Remove resources."
	@echo "dc-restart   Restart services."
	@echo "dc-stop      Stop services."
	@echo "dc-up        Create and Start containers."
	@echo "go-build     Compile packages and dependencies."
	@echo "go-get       Add dependencies to current module and install them."
	@echo "go-init      Initialize new module in current directory."
	@echo "go-run       Compile and Run Go program."
	@echo "go-test      Test packages."
	@echo "go-tidy      Add missing and remove unused modules."
	@echo "go-update    Remove files go.mod and go.sum."
	@echo "             Initialize new module in current directory."
	@echo "             Add missing and remove unused modules."
	@echo "rm-deps      Remove files go.mod and go.sum."
	@echo "rm-logs      Remove all files from logs directory."
	@echo "swagger      Generate swagger documentation."
	@echo

dc-down:
	@docker-compose --file docker-compose.yml down --volumes

dc-restart:
	@docker-compose --file /docker-compose.yml restart

dc-stop:
	@docker-compose --file docker-compose.yml stop

dc-up:
	@docker-compose --file docker-compose.yml up --detach

go-build:
	@go build -o kyp-api .

go-get:
	@go get ./...

go-init:
	@go mod init

go-run:
	@go run main.go

go-test:
	@go test ./... --coverprofile=coverage.txt | grep -v "no test files" && go tool cover --func=coverage.txt

go-tidy:
	@go mod tidy

go-update: rm-deps go-init go-tidy

rm-deps:
	@rm -f go.mod go.sum

rm-logs:
	@rm -f logs/*

swagger:
	@swag init --parseDependency --parseInternal
	@sed "1,3d" docs/docs.go > docs/docs.go.tmp
	@mv docs/docs.go.tmp docs/docs.go
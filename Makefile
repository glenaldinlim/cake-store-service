.DEFAULT_GOAL := welcome
.PHONY: compose-up compose-down migration-up migration-down

welcome:
	@echo "available command: "

compose-up:
	go mod tidy && go mod vendor
	docker-compose -f docker-compose.yml up --build -d

compose-down:
	docker-compose kill
	docker-compose rm -f

migration-up:
	$(GOPATH)/bin/goose -dir ./database/migration mysql "cake:secret@/cake_store?parseTime=true" up

migration-down:
	$(GOPATH)/bin/goose -dir ./database/migration mysql "cake:secret@/cake_store?parseTime=true" down

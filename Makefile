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
	@echo "migration up"

migration-down:
	@echo "migration down"

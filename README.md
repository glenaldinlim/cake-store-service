# Cake Store Service
Library List
- Routing: [HttpRouter](https://github.com/julienschmidt/httprouter)
- Database
  - Driver: [MySQL Driver](https://github.com/go-sql-driver/mysql)
  - Migration Tool: [Goose](https://github.com/pressly/goose)
- Validation: [Validator](https://github.com/go-playground/validator)
- Logger: [Logrus](https://github.com/sirupsen/logrus)

## Running App
### 1. Start Compose
- With make command use `make compose-up` to compose image or without make command use `go mod tidy && go mod vendor && docker-compose -f docker-compose.yml up --build -d`
### 2. Migration
- Install goose: `go install github.com/pressly/goose/v3/cmd/goose@latest`
- Run `make migration-up` or `make migration-down` on project root directory
  - `make migration-up` will migrate the DB that has been define in `./database/migration` directory
  - `make migration-down` will rollback the migration before
### 3. Run Unit Test
TBA
### 4. Stop Compose
- With make command use `make compose-down` or without make command use `docker-compose kill && docker-compose rm -f`

## Additional Information
- Run `$GOPATH/bin/goose -dir ./database/migration mysql "cake:secret@/cake_store?parseTime=true" create <name> sql` to add migration, change `<name>` with what you will do, example `create_cakes_table` or `alter_cakes_table`
- This app is development in Windows via WSL2 with Ubuntu, so maybe there will any failure in MySQL DSN.
# Cake Store Service
Library List
- Routing: [HttpRouter](https://github.com/julienschmidt/httprouter)
- Database
  - Driver: [MySQL Driver](https://github.com/go-sql-driver/mysql)
  - Migration Tool: [Goose](https://github.com/pressly/goose)
- Validation: [Validator](https://github.com/go-playground/validator)
- Logger

## Running App
### 1. Compose Image
### 2. Migration
- Install goose: `go install github.com/pressly/goose/v3/cmd/goose@latest`
- Run `make migration-up` or `make migration-down` on project root directory
  - `make migration-up` will migrate the DB that has been define in `./database/migration` directory
  - `make migration-down` will rollback the migration before
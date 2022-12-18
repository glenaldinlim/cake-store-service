# Cake Store Service
A simple RESTful API created with GoLang and MySQL as Database Server
Postman Collection: https://www.postman.com/galactic-star-21684/workspace/cake-store-service-api  

URI: http://localhost:8090/api  
Available Route:
| Name    | Method | Endpoint   | Request Body |
| ------- | ------ | ---------- | ------------ |
| Index   | GET    | /cakes     | No           |
| Store   | POST   | /cakes     | Yes          |
| Show    | GET    | /cakes/:id | No           |
| Update  | PATCH  | /cakes/:id | Yes          |
| Destroy | DELETE | /cakes/:id | No           |

Example Request Body Payload:
```json
{
  "title": "Tiramisu Oreo Cake",
  "description": "Cake with tiramisu flavour and oreo topping",
  "rating": 8.1,
  "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
}
```

Library List
- Routing: [HttpRouter](https://github.com/julienschmidt/httprouter)
- Database
  - Driver: [MySQL Driver](https://github.com/go-sql-driver/mysql)
  - Migration Tool: [Goose](https://github.com/pressly/goose)
- Validation: [Validator](https://github.com/go-playground/validator)
- Logger: [Logrus](https://github.com/sirupsen/logrus)
- Unit Test: [Testify](https://github.com/stretchr/testify)

## Running App
### 1. Start Compose
- With make command use `make compose-up` to compose image or without make command use `go mod tidy && go mod vendor && docker-compose -f docker-compose.yml up --build -d`
### 2. Migration
- Install goose: `go install github.com/pressly/goose/v3/cmd/goose@latest`
- Run `make migration-up` or `make migration-down` on project root directory
  - `make migration-up` will migrate the DB that has been define in `./database/migration` directory
  - `make migration-down` will rollback the migration before
### 3. Run Unit Test
Change directory to test directory and run selected command
- Run `go test -v` to run all test
- Run `go test -v -run=TestCakeRepo` to run cake repository test
- Run `go test -v -run=TestCakeSrv` to run cake service test
- Run `go test -v -run=TestCakeController` to run cake controller test

The `go test` command can in project root directory
### 4. Stop Compose
- With make command use `make compose-down` or without make command use `docker-compose kill && docker-compose rm -f`

## Additional Information
- Run `$GOPATH/bin/goose -dir ./database/migration mysql "cake:secret@/cake_store?parseTime=true" create <name> sql` to add migration, change `<name>` with what you will do, example `create_cakes_table` or `alter_cakes_table`
- This app is development in Windows via WSL2 with Ubuntu, so maybe there will any failure in MySQL DSN.
# Go JWT Auth

Simple Rest API with JWT Authentication written in Go.

---

## Dependencies

- [Air](https://github.com/cosmtrek/air)
- [GoDotEnv](https://github.com/joho/godotenv)
- [Gin Web Framework](https://gin-gonic.com/)
- [Gorm](https://gorm.io/)
- [Bcrypt](https://github.com/golang/crypto)
- [JWT](https://github.com/golang-jwt/jwt)

---

## Database

[PostgreSQL](https://www.elephantsql.com/)

---

## Setting up the project

Follow the steps below to set up the project.

### Environment Variables

create a `.env` file in the root directory and add the following:

```sh
PORT=8080
SECRET=YOUR_JWT_SECRET
DB_URL="host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable"
```

`PORT` is the port where the API will run (default: 8080). You can change it if you want.

`SECRET` is the JWT secret, replace it with your own.

`DB_URL` is the connection string for the database. Change it according to your database credentials.

---

## Running with Air (optional)

Air is a tool for running Go applications in the background, refreshing whenever it is modified.

If you got `air` installed, run the following command

```sh
air
```

## Running the API

Run the following command

```sh
go run ./cmd/main.go
```

---

## Testing

You can use [Postman](https://www.postman.com/) to test the API.

Base address: `http://localhost:{port}/`

ps: Check the port in `.env` file and feel free to change it.

---

## Endpoints

- [x] [POST] `~/signup`
- [x] [POST] `~/login`
- [x] [GET] `~/validate`

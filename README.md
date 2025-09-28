# gRPC Authentication Service

_Service for authentication and authorization made on gRPC_

## Functional

* Registration
* Log-in
* JWT token generation

## Stack of technologies

* Go 1.24.2
* gRPC
* Protocol Buffers(proto3)
* PostgreSQL
* bcrypt

## Installation

1. _Clone the repository_

```
https://github.com/RenSafary/grpc-auth.git
cd grpc-auth
```
2. _Install dependencies_
```
go mod tidy
```
3. _Сonfigure database_
```
CREATE DATABASE authgrpc;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username CHARACTER VARYING(50) NOT NULL,
    password text NOT NULL
);
```
4. _Compile the proto file_
```
protoc --go_out=. --go-grpc_out=. proto/auth.proto
```
5. _Start the server_
```
go run server/main.go
```
6. _Start the client_
```
go run client/main.go
```

## Environment variables

### Database
```
DB_USER="YOUR_USERNAME"
DB_PASSWORD="YOUR_PASSWORD"
DB_NAME="YOUR_DB_NAME"
DB_HOST="YOUR_HOST" 
DB_PORT="YOUR_PORT" 
DB_SSLMODE="disable"
```

### Json Web Token
```
SECRET_KEY="YOUR_SECRET_KEY"
```
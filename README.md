# gRPC Authentication Service

_Service for authentication and authorization made on gRPC_

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
    email CHARACTER VARYING(100),
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

## Functional

* Registration
* Log-in
* JWT token generation
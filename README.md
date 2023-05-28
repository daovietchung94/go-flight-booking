# grpc with go pratice

Pratice grpc microservices with go lang

## Setup env

**Go**

Ref: https://go.dev/doc/install

**Docker**

```sh
brew install docker --cask
```

**Goose**

```sh
brew install goose
```

## Development

**Starting database**

```sh
docker-compose up
```

**Proto gen**

***Install*** `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
```sh
protoc -I=proto proto/*.proto --go_out=:pb --go-grpc_out=:pb
```

## Testing

**1.Running the grpc service**

***Customer***
```sh
go run grpc/customer/main.go
```

**2.Running the api**

***Customer***
```sh
go run clients/rest/customer/main.go
```

go run github.com/99designs/gqlgen generate


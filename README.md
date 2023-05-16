```sh
mkdir basic-gqlgen
cd basic-gqlgen
go mod init basic-gqlgen
go get github.com/99designs/gqlgen
```

If not done yet, install `gqlgen` command
```
go install github.com/99designs/gqlgen@v0.17.27
```

Start using `gqlgen`
```sh
gqlgen init
```
Write following code (remove existing example lines) on `schema.graphql`
```
type Query {
  hello: String!
}
```

Run `generate` (ignore error messages, for now)
```sh
gqlgen generate
```
This will generate a file `generated.go` (See `gqlgen.yml` for configuration)

Write code and update `resolver.go`.
Remove some code from `schema.resolvers.go` file and then update `Hello` function
```go
return "world", nil
```

There might be an error message
```
validation failed: packages.Load: -: package basic-gqlgen/graph/model is not in GOROOT (/usr/local/go/src/basic-gqlgen/graph/model)
```

`gqlgen generate` command removes `graph/model/models_gen.go` file because we do not have any generated model (I think). Add a file NOT named `graph/model/models_gen.go` and put this line
```go
package model
```

Then the error messages on `gqlgen generate` will go away.
```
pqlgen validate
pqlgen generate
```

Write code on `server.go` file 
```sh
go run server.go
```

Open http://localhost:8080/ and run request with
```
query {
  hello
}
```

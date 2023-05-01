# simple-graph-ql


```shell

go mod init github.com/xcheng85/simple-graphql-go

go get github.com/99designs/gqlgen

go run github.com/99designs/gqlgen init

go get github.com/99designs/gqlgen/graphql/handler

# after graphql schema is updated
go get github.com/99designs/gqlgen@v0.17.30
go run github.com/99designs/gqlgen generate


query GetAllPlayers{
  GetAllPlayers{
    id
  }
}


```
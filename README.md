# simple-graph-ql


```shell

go mod init github.com/xcheng85/simple-graphql-go

go get github.com/99designs/gqlgen

go run github.com/99designs/gqlgen init

go get github.com/99designs/gqlgen/graphql/handler

# after graphql schema is updated
go get github.com/99designs/gqlgen@v0.17.30
go run github.com/99designs/gqlgen generate

# for advanced graphql schema
https://github.com/99designs/gqlgen/tree/master/_examples


query{
  GetAllPlayers{
    id
  }
}

//enum input has no "" 

mutation {
  addPlayer(player: {First: "Roger", Last: "Federer", Country: SWISS, Gender: MALE}){
    id
  }
}

mutation {
  addPlayer(player: {First: "Rafael", Last: "Nadal", Country: SPAIN, Gender: MALE}){
    id
  }
}

mutation ($first: String!){
  addPlayer(player: {First: $first, Last: "Federer", Country: SWISS, Gender: MALE}){
    id
  }
}

variables:

{
  "first": "Roger"
}

query{
  player(id:0){
    id
    First
    Last
    Country
    Gender
  }
}

```
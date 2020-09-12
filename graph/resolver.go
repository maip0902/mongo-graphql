package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
    "github.com/globalsign/mgo"
    "github.com/maip0902/mongo-graphql/graph/generated"
    "github.com/maip0902/mongo-graphql/db"
)

type Resolver struct{
    users *mgo.Collection
}

func New() *generated.Config {
    return &generated.Config {
        Resolvers: &Resolver{
            users: db.GetCollection("users"),
        },
    }
}

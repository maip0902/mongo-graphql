package resolver

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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver {
    r.users = db.GetCollection("users")
    return &mutationResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
    r.users = db.GetCollection("users")
    return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

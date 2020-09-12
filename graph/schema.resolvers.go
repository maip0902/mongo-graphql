package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/maip0902/mongo-graphql/graph/generated"
	"github.com/maip0902/mongo-graphql/graph/model"
	"github.com/maip0902/mongo-graphql/db"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		Email: input.Email,
	}

	// 	if err != nil {
	// 		return model.User{}, err
	// 	} else if count > 0 {
	// 		return model.User{}, errors.New("user with that email already exists")
	// 	}

	r.users.Insert(bson.M{"email": input.Email})
	// 	if err != nil {
	// 		return *model.User{}, err
	// 	}

	// 	r.users.Find(bson.M{"email": input.Email}).One(&user)
	// 	if err != nil {
	// 		return User{}, err
	// 	}

	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	var fields = bson.M{}
	var user *model.User
	update := false
	if input.First != nil && *input.First != "" {
		fields["first"] = *input.First
		update = true
	}
	if input.Last != nil && *input.Last != "" {
		fields["last"] = *input.Last
		update = true
	}
	if input.Email != nil && *input.Email != "" {
		fields["email"] = *input.Email
		update = true
	}
	if !update {
		return user, errors.New("no fields present for updating data")
	}
	err := r.users.UpdateId(bson.ObjectIdHex(input.ID), bson.M{"$set":fields})
	if err != nil {
		return user, err
	}
	err = r.users.Find(bson.M{"_id": bson.ObjectIdHex(input.ID)}).One(&user)
	if err != nil {
		return user, err
	}
	user.ID = bson.ObjectId(user.ID).Hex()
	fmt.Printf("%v", user)
	return user, nil
}

func (r *mutationResolver) UpdateNotification(ctx context.Context, input *model.UpdateNotification) (*model.User, error) {
	var user *model.User
	var oid = bson.ObjectIdHex(input.UserID)
	if err := r.users.Find(bson.M{"_id": oid}).One(&user); err != nil {
		return user, err
	}
	for index, val := range user.Notifications {
		if bson.ObjectId(val.ID).Hex() == input.ID {
			val.Seen = input.Seen
			user.Notifications[index] = val
			break
		}
	}
	if err := r.users.UpdateId(oid, user); err != nil {
		return user, err
	}
	return user, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	var user *model.User
	if err := r.users.FindId(bson.ObjectIdHex(id)).One(&user); err != nil {
		return user, err
	}
	user.ID = bson.ObjectId(user.ID).Hex()
	return user, nil
}

func (r *subscriptionResolver) NotificationAdded(ctx context.Context, id string) (<-chan *model.User, error) {
	panic(fmt.Errorf("not implemented"))
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

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

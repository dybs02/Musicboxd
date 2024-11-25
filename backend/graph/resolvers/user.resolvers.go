package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"
	"musicboxd/database"
	"musicboxd/graph"
	"musicboxd/graph/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *queryResolver) UserByDisplayName(ctx context.Context, displayName string) (*model.UserResponse, error) {
	// TODO check jwt and permissions
	coll := database.GetDB().GetCollection("users")
	user := coll.FindOne(ctx, bson.M{"displayname": displayName})
	if user.Err() != nil {
		return nil, user.Err()
	}

	res := model.UserResponse{}
	err := user.Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

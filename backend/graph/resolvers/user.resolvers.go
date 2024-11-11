package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"
	"fmt"
	"musicboxd/graph"
	"musicboxd/graph/model"
)

func (r *queryResolver) UserByID(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserByID - userById"))
}

func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

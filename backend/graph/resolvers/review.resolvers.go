package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"
	"errors"
	"musicboxd/database"
	"musicboxd/graph"
	"musicboxd/graph/model"
	"musicboxd/hlp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mutationResolver) CreateOrUpdateReview(ctx context.Context, itemID string, itemType string, title *string, description *string, value *int) (*model.Review, error) {
	cc, err := ValidateJWT(ctx)
	if err != nil {
		return nil, err
	}

	accessToken, err := GetUserAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	var album *model.Album

	if itemType == "track" {
		track, err := hlp.SpotifyGetTrack(itemID, accessToken)
		if err != nil {
			return nil, err
		}
		album = track.Album
	} else if itemType == "album" {
		album, err = hlp.SpotifyGetAlbum(itemID, accessToken)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("invalid item type")
	}

	dbAlbum := map[string]interface{}{
		"name":    album.Name,
		"images":  album.Images,
		"artists": album.Artists,
	}

	coll := database.GetDB().GetCollection("reviews")
	review := coll.FindOneAndUpdate(
		ctx,
		bson.M{"itemId": itemID, "userId": cc.UserID},
		bson.M{
			"$setOnInsert": bson.M{
				"comments": []interface{}{},
			},
			"$set": bson.M{
				"title":       title,
				"description": description,
				"value":       value,
				"album":       dbAlbum,
			}},
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After),
	)
	if review.Err() != nil {
		return nil, review.Err()
	}

	res := model.Review{}
	err = review.Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *mutationResolver) AddComment(ctx context.Context, itemID string, reviewID string, text string) ([]*model.Comment, error) {
	cc, err := ValidateJWT(ctx)
	if err != nil {
		return nil, err
	}

	convertedReviewID, err := primitive.ObjectIDFromHex(reviewID)
	if err != nil {
		return nil, err
	}

	coll := database.GetDB().GetCollection("reviews")
	comment := coll.FindOneAndUpdate(
		ctx,
		bson.M{"itemId": itemID, "userId": convertedReviewID},
		bson.M{
			"$push": bson.M{
				"comments": bson.M{
					"_id":       primitive.NewObjectID(),
					"reviewId":  reviewID, // why would i need it here?
					"userId":    cc.UserID,
					"text":      text,
					"createdAt": time.Now(), // TODO check if timezone independent - utc+0?
					"updatedAt": time.Now(),
				},
			},
		},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	if comment.Err() != nil {
		return nil, comment.Err()
	}

	res := model.Review{}
	err = comment.Decode(&res)
	if err != nil {
		return nil, err
	}

	if isFieldRequested(ctx, "user") {
		for _, comment := range res.Comments {
			coll := database.GetDB().GetCollection("users")

			convertedID, err := primitive.ObjectIDFromHex(*comment.UserID)
			if err != nil {
				return nil, err
			}

			user := coll.FindOne(ctx, bson.M{"_id": convertedID})
			if user.Err() != nil {
				return nil, user.Err()
			}

			u := model.UserResponse{}
			err = user.Decode(&u)
			if err != nil {
				return nil, err
			}

			comment.User = &u
		}
	}

	// TODO return only the added comment?
	return res.Comments, nil
}

func (r *queryResolver) Review(ctx context.Context, itemID string, userID string) (*model.Review, error) {
	_, err := ValidateJWT(ctx)
	if err != nil {
		return nil, err
	}

	convertedID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	coll := database.GetDB().GetCollection("reviews")
	review := coll.FindOne(ctx, bson.M{"itemId": itemID, "userId": convertedID})
	if review.Err() != nil {
		return nil, review.Err()
	}

	res := model.Review{}
	err = review.Decode(&res)
	if err != nil {
		return nil, err
	}

	if isFieldRequested(ctx, "comments.user") {
		for _, comment := range res.Comments {
			coll := database.GetDB().GetCollection("users")

			convertedID, err := primitive.ObjectIDFromHex(*comment.UserID)
			if err != nil {
				return nil, err
			}

			user := coll.FindOne(ctx, bson.M{"_id": convertedID})
			if user.Err() != nil {
				return nil, user.Err()
			}

			u := model.UserResponse{}
			err = user.Decode(&u)
			if err != nil {
				return nil, err
			}

			comment.User = &u
		}
	}

	return &res, nil
}

func (r *queryResolver) RecentReviews(ctx context.Context, number *int) ([]*model.Review, error) {
	coll := database.GetDB().GetCollection("reviews")
	cursor, err := coll.Find(
		ctx,
		bson.M{},
		options.Find().SetSort(bson.M{"createdAt": -1}).SetLimit(int64(min(*number, 20))),
	)
	if err != nil {
		return nil, err
	}

	res := []*model.Review{}
	for cursor.Next(ctx) {
		r := model.Review{}
		err = cursor.Decode(&r)
		if err != nil {
			return nil, err
		}

		// comments.user field is not added

		res = append(res, &r)
	}

	return res, nil
}

func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

package database

import (
	"context"
	"musicboxd/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetReviews(ctx context.Context, reviewIDMap map[string]bool, userID primitive.ObjectID) (map[string]*model.Review, error) {
	reviewIDs := make([]primitive.ObjectID, 0, len(reviewIDMap))
	for id := range reviewIDMap {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		reviewIDs = append(reviewIDs, objID)
	}

	reviewsColl := GetDB().GetCollection("reviews")
	reviewsCursor, err := reviewsColl.Find(
		ctx,
		bson.M{"_id": bson.M{"$in": reviewIDs}},
		options.Find().SetProjection(GetReviewProjection(userID)),
	)
	if err != nil {
		return nil, err
	}
	defer reviewsCursor.Close(ctx)

	var reviews []*model.Review
	if err = reviewsCursor.All(ctx, &reviews); err != nil {
		return nil, err
	}

	reviewMap := make(map[string]*model.Review)
	for _, review := range reviews {
		reviewMap[*review.ID] = review
	}

	return reviewMap, nil
}

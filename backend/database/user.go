package database

import (
	"context"
	"musicboxd/graph/model"
	"musicboxd/hlp/jwt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO replace all ocurences of getting user by ID with this function
// TODO add user projection to not query following and followers
func GetUserByPrimitiveID(ctx context.Context, userID primitive.ObjectID) (*model.UserResponse, error) {
	cc, err := jwt.ValidateJWTFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	coll := GetDB().GetCollection("users")
	user := coll.FindOne(
		ctx,
		bson.M{"_id": userID},
		options.FindOne().SetProjection(GetUserProjection(cc.UserID)),
	)
	if user.Err() != nil {
		return nil, user.Err()
	}

	u := model.UserResponse{}
	err = user.Decode(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func GetUserByID(ctx context.Context, userID string) (*model.UserResponse, error) {
	convertedID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	return GetUserByPrimitiveID(ctx, convertedID)
}

func GetFollowingIDs(ctx context.Context, userID primitive.ObjectID) ([]primitive.ObjectID, error) {
	res, err := GetUserByPrimitiveID(ctx, userID)
	if err != nil {
		return nil, err
	}

	followingIDs := make([]primitive.ObjectID, len(res.FollowingUsers))
	for i, userIDStr := range res.FollowingUsers {
		objID, err := primitive.ObjectIDFromHex(*userIDStr)
		if err != nil {
			return nil, err
		}
		followingIDs[i] = objID
	}

	return followingIDs, nil
}

func GetUsers(ctx context.Context, userIDMap map[string]bool) (map[string]*model.UserResponse, error) {
	userIDs := make([]primitive.ObjectID, 0, len(userIDMap))
	for id := range userIDMap {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		userIDs = append(userIDs, objID)
	}

	usersColl := GetDB().GetCollection("users")
	usersCursor, err := usersColl.Find(
		ctx,
		bson.M{"_id": bson.M{"$in": userIDs}},
	)
	if err != nil {
		return nil, err
	}
	defer usersCursor.Close(ctx)

	var users []*model.UserResponse
	if err = usersCursor.All(ctx, &users); err != nil {
		return nil, err
	}

	userMap := make(map[string]*model.UserResponse)
	for _, user := range users {
		userMap[user.ID] = user
	}

	return userMap, nil
}

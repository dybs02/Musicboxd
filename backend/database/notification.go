package database

import (
	"context"
	"fmt"
	"musicboxd/graph/model"
	"musicboxd/hlp/jwt"
	"time"

	"github.com/99designs/gqlgen/graphql/handler/transport"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddNotification(ctx context.Context, message string, nType string, notifiedUserID string) error {
	convertedNotifiedUserID, err := primitive.ObjectIDFromHex(notifiedUserID)
	if err != nil {
		return fmt.Errorf("invalid notified user ID: %w", err)
	}

	return AddNotificationPrimitiveID(ctx, message, nType, convertedNotifiedUserID)
}

func AddNotificationPrimitiveID(ctx context.Context, message string, nType string, notifiedUserID primitive.ObjectID) error {
	connParams := transport.GetInitPayload(ctx)
	jwtString := connParams["Authorization"].(string)
	cc, err := jwt.ValidateJWT(jwtString)
	if err != nil {
		return err
	}

	if nType != "chat" {
		return fmt.Errorf("invalid notification type: %s", nType)
	}

	coll := GetDB().GetCollection("notifications")

	res := coll.FindOneAndUpdate(
		ctx,
		bson.M{
			"type":            nType,
			"notifyingUserId": cc.UserID,
			"notifiedUserId":  notifiedUserID,
			"isRead":          false,
		},
		bson.M{
			"$set": bson.M{
				"createdAt": time.Now(),
			},
			"$setOnInsert": bson.M{
				"message":         message,
				"type":            nType,
				"notifyingUserId": cc.UserID,
				"notifiedUserId":  notifiedUserID,
				"isRead":          false,
			},
		},
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After),
	)
	if res.Err() != nil {
		return fmt.Errorf("failed to upsert notification: %w", res.Err())
	}

	var result model.Notification
	err = res.Decode(&result)
	if err != nil {
		return fmt.Errorf("failed to decode notification: %w", err)
	}

	return nil
}

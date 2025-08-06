package database

import (
	"context"
	"fmt"
	"musicboxd/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetChat(ctx context.Context, chatID string) (*model.Chat, error) {
	convertedChatID, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		return nil, fmt.Errorf("invalid chat ID: %w", err)
	}

	coll := GetDB().GetCollection("chats")
	var chat model.Chat
	err = coll.FindOne(
		ctx,
		bson.M{"_id": convertedChatID},
		options.FindOne().SetProjection(GetChatProjection()),
	).Decode(&chat)
	if err != nil {
		return nil, fmt.Errorf("failed to find chat: %w", err)
	}

	return &chat, nil
}

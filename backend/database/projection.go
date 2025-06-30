package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetReviewProjection(userID primitive.ObjectID) *bson.M {
	return &bson.M{
		// Calculate fields
		"likesCount":    bson.M{"$size": bson.M{"$ifNull": bson.A{"$likes", bson.A{}}}},
		"dislikesCount": bson.M{"$size": bson.M{"$ifNull": bson.A{"$dislikes", bson.A{}}}},
		"userReaction": bson.M{
			"$cond": bson.A{
				bson.M{"$in": bson.A{userID, bson.M{"$ifNull": bson.A{"$likes", bson.A{}}}}},
				"like",
				bson.M{
					"$cond": bson.A{
						bson.M{"$in": bson.A{userID, bson.M{"$ifNull": bson.A{"$dislikes", bson.A{}}}}},
						"dislike",
						"",
					},
				},
			},
		},
		// Include most fields
		"_id":         1,
		"value":       1,
		"itemId":      1,
		"itemType":    1,
		"title":       1,
		"description": 1,
		"userId":      1,
		"user":        1,
		"createdAt":   1,
		"updatedAt":   1,
		// TODO Add this as a function parameter to not query if not needed
		"commentIds": 1,
		"album":      1,
		// Exclude likes and dislikes
		"likes":    bson.M{"$cond": bson.A{true, bson.A{}, "$likes"}},    // Empty array
		"dislikes": bson.M{"$cond": bson.A{true, bson.A{}, "$dislikes"}}, // Empty array
	}
}

func GetCommentProjection(userID primitive.ObjectID) *bson.M {
	return &bson.M{
		// Calculate fields
		"repliesCount":  bson.M{"$size": bson.M{"$ifNull": bson.A{"$repliesIds", bson.A{}}}},
		"likesCount":    bson.M{"$size": bson.M{"$ifNull": bson.A{"$likes", bson.A{}}}},
		"dislikesCount": bson.M{"$size": bson.M{"$ifNull": bson.A{"$dislikes", bson.A{}}}},
		"userReaction": bson.M{
			"$cond": bson.A{
				bson.M{"$in": bson.A{userID, bson.M{"$ifNull": bson.A{"$likes", bson.A{}}}}},
				"like",
				bson.M{
					"$cond": bson.A{
						bson.M{"$in": bson.A{userID, bson.M{"$ifNull": bson.A{"$dislikes", bson.A{}}}}},
						"dislike",
						"",
					},
				},
			},
		},
		// Include most fields
		"_id":       1,
		"itemId":    1,
		"userId":    1,
		"text":      1,
		"createdAt": 1,
		"updatedAt": 1,
		// Exclude likes and dislikes
		"likes":    bson.M{"$cond": bson.A{true, bson.A{}, "$likes"}},    // Empty array
		"dislikes": bson.M{"$cond": bson.A{true, bson.A{}, "$dislikes"}}, // Empty array
	}
}

func GetPostProjection(userID primitive.ObjectID) *bson.M {
	return &bson.M{
		// Calculate fields
		"commentsNumber": bson.M{"$size": bson.M{"$ifNull": bson.A{"$commentIds", bson.A{}}}},
		"likesCount":     bson.M{"$size": bson.M{"$ifNull": bson.A{"$likes", bson.A{}}}},
		"dislikesCount":  bson.M{"$size": bson.M{"$ifNull": bson.A{"$dislikes", bson.A{}}}},
		"userReaction": bson.M{
			"$cond": bson.A{
				bson.M{"$in": bson.A{userID, bson.M{"$ifNull": bson.A{"$likes", bson.A{}}}}},
				"like",
				bson.M{
					"$cond": bson.A{
						bson.M{"$in": bson.A{userID, bson.M{"$ifNull": bson.A{"$dislikes", bson.A{}}}}},
						"dislike",
						"",
					},
				},
			},
		},
		// Include most fields
		"_id":       1,
		"content":   1,
		"userId":    1,
		"user":      1,
		"createdAt": 1,
		"updatedAt": 1,
		// TODO Add this as a function parameter to not query if not needed
		"commentIds":     1,
		"linkedReviewId": 1,
		// Exclude likes and dislikes
		"likes":    bson.M{"$cond": bson.A{true, bson.A{}, "$likes"}},    // Empty array
		"dislikes": bson.M{"$cond": bson.A{true, bson.A{}, "$dislikes"}}, // Empty array
	}
}

func GetUserProjection(userID primitive.ObjectID) *bson.M {
	return &bson.M{
		// Calculate fields
		"isFollowing": bson.M{"$cond": bson.A{
			bson.M{"$in": bson.A{userID, bson.M{"$ifNull": bson.A{"$followingUsers", bson.A{}}}}},
			true,
			false,
		}},
		"isFollower": bson.M{"$cond": bson.A{
			bson.M{"$in": bson.A{userID, bson.M{"$ifNull": bson.A{"$followerUsers", bson.A{}}}}},
			true,
			false,
		}},
		// Include most fields
		"_id":                1,
		"country":            1,
		"displayName":        1,
		"email":              1,
		"explicitContent":    1,
		"externalUrls":       1,
		"followers":          1,
		"href":               1,
		"spotifyId":          1,
		"images":             1,
		"product":            1,
		"type":               1,
		"uri":                1,
		"role":               1,
		"favouriteAlbums":    1,
		"trackReviewsNumber": 1,
		"albumReviewsNumber": 1,
		"followingUsers":     1,
		"followerUsers":      1,
	}
}

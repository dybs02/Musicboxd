// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Album struct {
	AlbumType            string              `json:"album_type" bson:"album_type"`
	TotalTracks          int                 `json:"total_tracks" bson:"total_tracks"`
	AvailableMarkets     []string            `json:"available_markets" bson:"available_markets"`
	ExternalUrls         *ExternalUrls       `json:"external_urls" bson:"external_urls"`
	Href                 string              `json:"href" bson:"href"`
	ID                   string              `json:"id" bson:"_id"`
	Images               []*Image            `json:"images" bson:"images"`
	Name                 string              `json:"name" bson:"name"`
	ReleaseDate          string              `json:"release_date" bson:"release_date"`
	ReleaseDatePrecision string              `json:"release_date_precision" bson:"release_date_precision"`
	Type                 string              `json:"type" bson:"type"`
	URI                  string              `json:"uri" bson:"uri"`
	Artists              []*SimplifiedArtist `json:"artists" bson:"artists"`
	Tracks               *Tracks             `json:"tracks" bson:"tracks"`
	AverageRating        *float64            `json:"averageRating,omitempty" bson:"averageRating,omitempty"`
	RatingCount          *int                `json:"ratingCount,omitempty" bson:"ratingCount,omitempty"`
}

type Albums struct {
	Href     string   `json:"href" bson:"href"`
	Limit    int      `json:"limit" bson:"limit"`
	Next     string   `json:"next" bson:"next"`
	Offset   int      `json:"offset" bson:"offset"`
	Previous string   `json:"previous" bson:"previous"`
	Total    int      `json:"total" bson:"total"`
	Items    []*Album `json:"items" bson:"items"`
}

type Chat struct {
	ID              string          `json:"_id" bson:"_id"`
	Name            *string         `json:"name,omitempty" bson:"name,omitempty"`
	ParticipantsIds []string        `json:"participantsIds" bson:"participantsIds"`
	Participants    []*UserResponse `json:"participants" bson:"participants"`
	ParticipantID   string          `json:"participantId" bson:"participantId"`
	Participant     *UserResponse   `json:"participant" bson:"participant"`
	Messages        []*Message      `json:"messages" bson:"messages"`
	CreatedAt       time.Time       `json:"createdAt" bson:"createdAt"`
}

type Comment struct {
	ID            *string       `json:"_id,omitempty" bson:"_id,omitempty"`
	ItemID        *string       `json:"itemId,omitempty" bson:"itemId,omitempty"`
	UserID        *string       `json:"userId,omitempty" bson:"userId,omitempty"`
	User          *UserResponse `json:"user,omitempty" bson:"user,omitempty"`
	Text          *string       `json:"text,omitempty" bson:"text,omitempty"`
	CreatedAt     *time.Time    `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt     *time.Time    `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	LikesCount    int           `json:"likesCount" bson:"likesCount"`
	DislikesCount int           `json:"dislikesCount" bson:"dislikesCount"`
	UserReaction  *string       `json:"userReaction,omitempty" bson:"userReaction,omitempty"`
	ReplyingToID  *string       `json:"replyingToId,omitempty" bson:"replyingToId,omitempty"`
	RepliesCount  int           `json:"repliesCount" bson:"repliesCount"`
	Replies       []*Comment    `json:"replies" bson:"replies"`
	Likes         []*string     `json:"likes,omitempty" bson:"likes,omitempty"`
	Dislikes      []*string     `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
	RepliesIds    []*string     `json:"repliesIds,omitempty" bson:"repliesIds,omitempty"`
}

type CommentsPage struct {
	TotalComments   int        `json:"totalComments" bson:"totalComments"`
	TotalPages      int        `json:"totalPages" bson:"totalPages"`
	HasPreviousPage bool       `json:"hasPreviousPage" bson:"hasPreviousPage"`
	HasNextPage     bool       `json:"hasNextPage" bson:"hasNextPage"`
	Comments        []*Comment `json:"comments" bson:"comments"`
}

type ExplicitContent struct {
	FilterEnabled bool `json:"filterEnabled" bson:"filterEnabled"`
	FilterLocked  bool `json:"filterLocked" bson:"filterLocked"`
}

type ExternalUrls struct {
	Spotify string `json:"spotify" bson:"spotify"`
}

type FavouriteAlbumEntry struct {
	Key   int          `json:"key" bson:"key"`
	Album *ReviewAlbum `json:"album,omitempty" bson:"album,omitempty"`
}

type FavouriteAlbumEntryInput struct {
	Key     int    `json:"key" bson:"key"`
	AlbumID string `json:"albumId" bson:"albumId"`
}

type Followers struct {
	Href  *string `json:"href,omitempty" bson:"href,omitempty"`
	Total int     `json:"total" bson:"total"`
}

type Image struct {
	URL    string `json:"url" bson:"url"`
	Height *int   `json:"height,omitempty" bson:"height,omitempty"`
	Width  *int   `json:"width,omitempty" bson:"width,omitempty"`
}

type Message struct {
	ID        string        `json:"_id" bson:"_id"`
	Content   string        `json:"content" bson:"content"`
	SenderID  string        `json:"senderId" bson:"senderId"`
	Sender    *UserResponse `json:"sender,omitempty" bson:"sender,omitempty"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
}

type MessagesPage struct {
	TotalMessages   int        `json:"totalMessages" bson:"totalMessages"`
	TotalPages      int        `json:"totalPages" bson:"totalPages"`
	HasPreviousPage bool       `json:"hasPreviousPage" bson:"hasPreviousPage"`
	HasNextPage     bool       `json:"hasNextPage" bson:"hasNextPage"`
	Messages        []*Message `json:"messages" bson:"messages"`
}

type Mutation struct {
}

type Post struct {
	ID             *string       `json:"_id,omitempty" bson:"_id,omitempty"`
	Content        string        `json:"content" bson:"content"`
	UserID         string        `json:"userId" bson:"userId"`
	User           *UserResponse `json:"user,omitempty" bson:"user,omitempty"`
	CreatedAt      time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt      *time.Time    `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	CommentIds     []*string     `json:"commentIds,omitempty" bson:"commentIds,omitempty"`
	Comments       []*Comment    `json:"comments,omitempty" bson:"comments,omitempty"`
	CommentsNumber int           `json:"commentsNumber" bson:"commentsNumber"`
	LinkedReviewID *string       `json:"linkedReviewId,omitempty" bson:"linkedReviewId,omitempty"`
	LinkedReview   *Review       `json:"linkedReview,omitempty" bson:"linkedReview,omitempty"`
	Likes          []*string     `json:"likes,omitempty" bson:"likes,omitempty"`
	LikesCount     int           `json:"likesCount" bson:"likesCount"`
	Dislikes       []*string     `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
	DislikesCount  int           `json:"dislikesCount" bson:"dislikesCount"`
	UserReaction   *string       `json:"userReaction,omitempty" bson:"userReaction,omitempty"`
}

type Query struct {
}

type RecentPosts struct {
	TotalPosts      int     `json:"totalPosts" bson:"totalPosts"`
	TotalPages      int     `json:"totalPages" bson:"totalPages"`
	HasPreviousPage bool    `json:"hasPreviousPage" bson:"hasPreviousPage"`
	HasNextPage     bool    `json:"hasNextPage" bson:"hasNextPage"`
	Posts           []*Post `json:"posts" bson:"posts"`
}

type RecentUserReviews struct {
	TotalReviews    int       `json:"totalReviews" bson:"totalReviews"`
	TotalPages      int       `json:"totalPages" bson:"totalPages"`
	HasPreviousPage bool      `json:"hasPreviousPage" bson:"hasPreviousPage"`
	HasNextPage     bool      `json:"hasNextPage" bson:"hasNextPage"`
	Reviews         []*Review `json:"reviews" bson:"reviews"`
}

type ReportedComment struct {
	ID             *string       `json:"_id,omitempty" bson:"_id,omitempty"`
	CommentID      string        `json:"commentId" bson:"commentId"`
	Comment        *Comment      `json:"comment,omitempty" bson:"comment,omitempty"`
	ReportedBy     string        `json:"reportedBy" bson:"reportedBy"`
	ReportedByUser *UserResponse `json:"reportedByUser,omitempty" bson:"reportedByUser,omitempty"`
	Status         string        `json:"status" bson:"status"`
	CreatedAt      time.Time     `json:"createdAt" bson:"createdAt"`
	ResolvedAt     *time.Time    `json:"resolvedAt,omitempty" bson:"resolvedAt,omitempty"`
	ModeratorID    *string       `json:"moderatorId,omitempty" bson:"moderatorId,omitempty"`
	ModeratorNotes *string       `json:"moderatorNotes,omitempty" bson:"moderatorNotes,omitempty"`
}

type Review struct {
	ID            *string       `json:"_id,omitempty" bson:"_id,omitempty"`
	Value         int           `json:"value" bson:"value"`
	ItemID        string        `json:"itemId" bson:"itemId"`
	ItemType      string        `json:"itemType" bson:"itemType"`
	Title         string        `json:"title" bson:"title"`
	Description   string        `json:"description" bson:"description"`
	UserID        string        `json:"userId" bson:"userId"`
	User          *UserResponse `json:"user,omitempty" bson:"user,omitempty"`
	CreatedAt     time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt     *time.Time    `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	CommentIds    []*string     `json:"commentIds,omitempty" bson:"commentIds,omitempty"`
	Comments      []*Comment    `json:"comments,omitempty" bson:"comments,omitempty"`
	Track         *ReviewTrack  `json:"track,omitempty" bson:"track,omitempty"`
	Album         *ReviewAlbum  `json:"album,omitempty" bson:"album,omitempty"`
	Likes         []*string     `json:"likes,omitempty" bson:"likes,omitempty"`
	LikesCount    int           `json:"likesCount" bson:"likesCount"`
	Dislikes      []*string     `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
	DislikesCount int           `json:"dislikesCount" bson:"dislikesCount"`
	UserReaction  *string       `json:"userReaction,omitempty" bson:"userReaction,omitempty"`
}

type ReviewAlbum struct {
	AlbumID string              `json:"albumId" bson:"albumId"`
	Name    string              `json:"name" bson:"name"`
	Images  []*Image            `json:"images" bson:"images"`
	Artists []*SimplifiedArtist `json:"artists" bson:"artists"`
}

type ReviewTrack struct {
	TrackID string              `json:"trackId" bson:"trackId"`
	Name    string              `json:"name" bson:"name"`
	Artists []*SimplifiedArtist `json:"artists" bson:"artists"`
}

type SearchResponse struct {
	Tracks *Tracks `json:"tracks,omitempty" bson:"tracks,omitempty"`
	Albums *Albums `json:"albums,omitempty" bson:"albums,omitempty"`
}

type SimplifiedArtist struct {
	ExternalUrls *ExternalUrls `json:"external_urls" bson:"external_urls"`
	Href         string        `json:"href" bson:"href"`
	ID           string        `json:"id" bson:"_id"`
	Name         string        `json:"name" bson:"name"`
	Type         string        `json:"type" bson:"type"`
	URI          string        `json:"uri" bson:"uri"`
}

type Subscription struct {
}

type Tokens struct {
	AccessToken  *string `json:"AccessToken,omitempty" bson:"AccessToken,omitempty"`
	TokenType    *string `json:"TokenType,omitempty" bson:"TokenType,omitempty"`
	Scope        *string `json:"Scope,omitempty" bson:"Scope,omitempty"`
	ExpiresIn    *int    `json:"ExpiresIn,omitempty" bson:"ExpiresIn,omitempty"`
	RefreshToken *string `json:"RefreshToken,omitempty" bson:"RefreshToken,omitempty"`
}

type Track struct {
	Album            *Album              `json:"album" bson:"album"`
	Artists          []*SimplifiedArtist `json:"artists" bson:"artists"`
	AvailableMarkets []string            `json:"available_markets" bson:"available_markets"`
	DiscNumber       int                 `json:"disc_number" bson:"disc_number"`
	DurationMs       int                 `json:"duration_ms" bson:"duration_ms"`
	Explicit         bool                `json:"explicit" bson:"explicit"`
	ExternalUrls     *ExternalUrls       `json:"external_urls" bson:"external_urls"`
	Href             string              `json:"href" bson:"href"`
	ID               string              `json:"id" bson:"_id"`
	IsPlayable       bool                `json:"is_playable" bson:"is_playable"`
	Name             string              `json:"name" bson:"name"`
	Popularity       int                 `json:"popularity" bson:"popularity"`
	PreviewURL       *string             `json:"preview_url,omitempty" bson:"preview_url,omitempty"`
	TrackNumber      int                 `json:"track_number" bson:"track_number"`
	Type             string              `json:"type" bson:"type"`
	URI              string              `json:"uri" bson:"uri"`
	IsLocal          bool                `json:"is_local" bson:"is_local"`
	AverageRating    *float64            `json:"averageRating,omitempty" bson:"averageRating,omitempty"`
	RatingCount      *int                `json:"ratingCount,omitempty" bson:"ratingCount,omitempty"`
}

type Tracks struct {
	Href     *string  `json:"href,omitempty" bson:"href,omitempty"`
	Limit    *int     `json:"limit,omitempty" bson:"limit,omitempty"`
	Next     *string  `json:"next,omitempty" bson:"next,omitempty"`
	Offset   int      `json:"offset" bson:"offset"`
	Previous string   `json:"previous" bson:"previous"`
	Total    int      `json:"total" bson:"total"`
	Items    []*Track `json:"items" bson:"items"`
}

type User struct {
	ID              *string                `json:"_id,omitempty" bson:"_id,omitempty"`
	Country         string                 `json:"country" bson:"country"`
	DisplayName     string                 `json:"displayName" bson:"displayName"`
	Email           string                 `json:"email" bson:"email"`
	ExplicitContent *ExplicitContent       `json:"explicitContent" bson:"explicitContent"`
	ExternalUrls    *ExternalUrls          `json:"externalUrls" bson:"externalUrls"`
	Followers       *Followers             `json:"followers" bson:"followers"`
	Href            string                 `json:"href" bson:"href"`
	SpotifyID       string                 `json:"spotifyId" bson:"spotifyId"`
	Images          []*Image               `json:"images" bson:"images"`
	Product         string                 `json:"product" bson:"product"`
	Type            string                 `json:"type" bson:"type"`
	URI             string                 `json:"uri" bson:"uri"`
	Tokens          *Tokens                `json:"tokens,omitempty" bson:"tokens,omitempty"`
	Role            string                 `json:"role" bson:"role"`
	FavouriteAlbums []*FavouriteAlbumEntry `json:"favouriteAlbums" bson:"favouriteAlbums"`
	FollowingUsers  []*string              `json:"followingUsers" bson:"followingUsers"`
	FollowerUsers   []*string              `json:"followerUsers" bson:"followerUsers"`
	IsFollowing     *bool                  `json:"isFollowing,omitempty" bson:"isFollowing,omitempty"`
	IsFollower      *bool                  `json:"isFollower,omitempty" bson:"isFollower,omitempty"`
}

type UserResponse struct {
	ID                 string                 `json:"_id" bson:"_id"`
	Country            string                 `json:"country" bson:"country"`
	DisplayName        string                 `json:"displayName" bson:"displayName"`
	Email              string                 `json:"email" bson:"email"`
	ExplicitContent    *ExplicitContent       `json:"explicitContent" bson:"explicitContent"`
	ExternalUrls       *ExternalUrls          `json:"externalUrls" bson:"externalUrls"`
	Followers          *Followers             `json:"followers" bson:"followers"`
	Href               string                 `json:"href" bson:"href"`
	SpotifyID          string                 `json:"spotifyId" bson:"spotifyId"`
	Images             []*Image               `json:"images" bson:"images"`
	Product            string                 `json:"product" bson:"product"`
	Type               string                 `json:"type" bson:"type"`
	URI                string                 `json:"uri" bson:"uri"`
	Role               string                 `json:"role" bson:"role"`
	FavouriteAlbums    []*FavouriteAlbumEntry `json:"favouriteAlbums" bson:"favouriteAlbums"`
	TrackReviewsNumber *int                   `json:"trackReviewsNumber,omitempty" bson:"trackReviewsNumber,omitempty"`
	AlbumReviewsNumber *int                   `json:"albumReviewsNumber,omitempty" bson:"albumReviewsNumber,omitempty"`
	FollowingUsers     []*string              `json:"followingUsers" bson:"followingUsers"`
	FollowerUsers      []*string              `json:"followerUsers" bson:"followerUsers"`
	IsFollowing        *bool                  `json:"isFollowing,omitempty" bson:"isFollowing,omitempty"`
	IsFollower         *bool                  `json:"isFollower,omitempty" bson:"isFollower,omitempty"`
}

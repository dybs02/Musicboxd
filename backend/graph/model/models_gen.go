// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Album struct {
	AlbumType            string   `json:"album_type" bson:"album_type"`
	TotalTracks          int      `json:"total_tracks" bson:"total_tracks"`
	AvailableMarkets     []string `json:"available_markets" bson:"available_markets"`
	Href                 string   `json:"href" bson:"href"`
	ID                   string   `json:"id" bson:"_id"`
	Images               []*Image `json:"images" bson:"images"`
	Name                 string   `json:"name" bson:"name"`
	ReleaseDate          string   `json:"release_date" bson:"release_date"`
	ReleaseDatePrecision string   `json:"release_date_precision" bson:"release_date_precision"`
	Type                 string   `json:"type" bson:"type"`
	URI                  string   `json:"uri" bson:"uri"`
}

type ExplicitContent struct {
	FilterEnabled bool `json:"filterEnabled" bson:"filterEnabled"`
	FilterLocked  bool `json:"filterLocked" bson:"filterLocked"`
}

type ExternalUrls struct {
	Spotify string `json:"spotify" bson:"spotify"`
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

type Mutation struct {
}

type Query struct {
}

type SearchResponse struct {
	Tracks *Tracks `json:"tracks" bson:"tracks"`
}

type Tokens struct {
	AccessToken  *string `json:"AccessToken,omitempty" bson:"AccessToken,omitempty"`
	TokenType    *string `json:"TokenType,omitempty" bson:"TokenType,omitempty"`
	Scope        *string `json:"Scope,omitempty" bson:"Scope,omitempty"`
	ExpiresIn    *int    `json:"ExpiresIn,omitempty" bson:"ExpiresIn,omitempty"`
	RefreshToken *string `json:"RefreshToken,omitempty" bson:"RefreshToken,omitempty"`
}

type Track struct {
	Album            *Album   `json:"album" bson:"album"`
	AvailableMarkets []string `json:"available_markets" bson:"available_markets"`
	DiscNumber       int      `json:"disc_number" bson:"disc_number"`
	DurationMs       int      `json:"duration_ms" bson:"duration_ms"`
	Explicit         bool     `json:"explicit" bson:"explicit"`
	Href             string   `json:"href" bson:"href"`
	ID               string   `json:"id" bson:"_id"`
	IsPlayable       bool     `json:"is_playable" bson:"is_playable"`
	Name             string   `json:"name" bson:"name"`
	Popularity       int      `json:"popularity" bson:"popularity"`
	PreviewURL       *string  `json:"preview_url,omitempty" bson:"preview_url,omitempty"`
	TrackNumber      int      `json:"track_number" bson:"track_number"`
	Type             string   `json:"type" bson:"type"`
	URI              string   `json:"uri" bson:"uri"`
	IsLocal          bool     `json:"is_local" bson:"is_local"`
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
	ID              *string          `json:"_id,omitempty" bson:"_id,omitempty"`
	Country         string           `json:"country" bson:"country"`
	DisplayName     string           `json:"displayName" bson:"displayName"`
	Email           string           `json:"email" bson:"email"`
	ExplicitContent *ExplicitContent `json:"explicitContent" bson:"explicitContent"`
	ExternalUrls    *ExternalUrls    `json:"externalUrls" bson:"externalUrls"`
	Followers       *Followers       `json:"followers" bson:"followers"`
	Href            string           `json:"href" bson:"href"`
	SpotifyID       string           `json:"spotifyId" bson:"spotifyId"`
	Images          []*Image         `json:"images" bson:"images"`
	Product         string           `json:"product" bson:"product"`
	Type            string           `json:"type" bson:"type"`
	URI             string           `json:"uri" bson:"uri"`
	Tokens          *Tokens          `json:"tokens,omitempty" bson:"tokens,omitempty"`
}

type UserResponse struct {
	ID              string           `json:"_id" bson:"_id"`
	Country         string           `json:"country" bson:"country"`
	DisplayName     string           `json:"displayName" bson:"displayName"`
	Email           string           `json:"email" bson:"email"`
	ExplicitContent *ExplicitContent `json:"explicitContent" bson:"explicitContent"`
	ExternalUrls    *ExternalUrls    `json:"externalUrls" bson:"externalUrls"`
	Followers       *Followers       `json:"followers" bson:"followers"`
	Href            string           `json:"href" bson:"href"`
	SpotifyID       string           `json:"spotifyId" bson:"spotifyId"`
	Images          []*Image         `json:"images" bson:"images"`
	Product         string           `json:"product" bson:"product"`
	Type            string           `json:"type" bson:"type"`
	URI             string           `json:"uri" bson:"uri"`
}

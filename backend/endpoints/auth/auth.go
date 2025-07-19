package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"musicboxd/database"
	"musicboxd/graph/model"
	"musicboxd/hlp"
	"musicboxd/hlp/jwt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

func LoginEndpoint(c *gin.Context) {
	spotifyOauth2Config := &oauth2.Config{
		RedirectURL:  hlp.Envs["SPOTIFY_REDIRECT_URI"],
		ClientID:     hlp.Envs["SPOTIFY_CLIENT_ID"],
		ClientSecret: hlp.Envs["SPOTIFY_CLIENT_SECRET"],
		Scopes:       []string{"user-read-private", "user-read-email"},
		Endpoint:     spotify.Endpoint,
	}

	codeVerifier := hlp.GenerateRandomString(64)
	codeChallenge := oauth2.S256ChallengeOption(codeVerifier)
	state := hlp.GenerateRandomString(16)
	url := spotifyOauth2Config.AuthCodeURL(state, codeChallenge)

	// TODO
	// Is it a good idea to store code_verifier in cookie to keep it stateless,
	// or should it be stored in a backend session?

	// TODO enable secure cookie only in production
	c.SetCookie(CODE_VERIFIER_KEY, codeVerifier, 0, "/", "", false, true)
	c.SetCookie(STATE_KEY, state, 0, "/", "", false, true)
	c.SetSameSite(http.SameSiteStrictMode)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func CallbackEndpoint(c *gin.Context) {
	state, err := c.Cookie(STATE_KEY)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no state cookie"})
		return
	}

	if c.Query("state") != state {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "state mismatch"})
		return
	}

	code := c.Query("code")
	c.Set(AUTHORIZATION_CODE, code)
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no code"})
		return
	}

	codeVerifier, err := c.Cookie(CODE_VERIFIER_KEY)
	c.Set(CODE_VERIFIER_KEY, codeVerifier)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no code verifier cookie"})
		return
	}

	c.SetCookie(CODE_VERIFIER_KEY, "", -1, "/", "", false, true)
	c.SetCookie(STATE_KEY, "", -1, "/", "", false, true)

	tokens, err := requestTokens(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no tokens"})
		return
	}

	user, err := requestUserData(tokens)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no user data"})
		return
	}

	dbUser, err := saveToDB(tokens, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save to db"})
		return
	}

	ui := jwt.UserInfo{
		ID:    *dbUser.ID,
		Email: dbUser.Email,
		Role:  dbUser.Role,
	}
	token, err := jwt.GenerateJWT(ui)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate jwt"})
		return
	}

	q := url.Values{}
	q.Set(USERID_KEY, *dbUser.ID)
	q.Set(JWT_KEY, token)
	q.Set("role", dbUser.Role)
	location := url.URL{Path: hlp.Envs["FRONTEND_URL"] + "/loginredirect", RawQuery: q.Encode()}

	isSecureCookie := false
	sameSiteMode := http.SameSiteLaxMode
	if hlp.Envs["ENVIRONMENT"] == "prod" {
		isSecureCookie = true
		sameSiteMode = http.SameSiteNoneMode
	}

	c.SetSameSite(sameSiteMode)
	c.SetCookie(JWT_KEY, token, 0, "/", hlp.Envs["BACKEND_DOMAIN"], isSecureCookie, true)
	c.SetSameSite(sameSiteMode)
	c.SetCookie(USERID_KEY, *dbUser.ID, 0, "/", hlp.Envs["BACKEND_DOMAIN"], isSecureCookie, false)
	c.Redirect(http.StatusFound, location.RequestURI())
}

func requestTokens(c *gin.Context) (*tokenResponse, error) {
	const (
		spotifyTokenURL = "https://accounts.spotify.com/api/token"
		grantType       = "authorization_code"
		contentType     = "application/x-www-form-urlencoded"
	)

	code, _ := c.Get(AUTHORIZATION_CODE)
	codeVerifier, _ := c.Get(CODE_VERIFIER_KEY)

	formData := url.Values{}
	formData.Add("grant_type", grantType)
	formData.Add("code", code.(string))
	formData.Add("redirect_uri", hlp.Envs["SPOTIFY_REDIRECT_URI"])
	formData.Add("client_id", hlp.Envs["SPOTIFY_CLIENT_ID"])
	formData.Add("code_verifier", codeVerifier.(string))

	req, err := http.NewRequest(http.MethodPost, spotifyTokenURL, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data tokenResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func requestUserData(tokens *tokenResponse) (*userData, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+tokens.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err

	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user userData
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func saveToDB(tokensData *tokenResponse, userData *userData) (*model.User, error) {
	explicitContent := &model.ExplicitContent{
		FilterEnabled: userData.ExplicitContent["filter_enabled"],
		FilterLocked:  userData.ExplicitContent["filter_locked"],
	}

	externalUrls := &model.ExternalUrls{
		Spotify: userData.ExternalURLs["spotify"],
	}

	followers := &model.Followers{
		Href:  &userData.FollowerCount.Href,
		Total: userData.FollowerCount.Total,
	}

	images := make([]*model.Image, len(userData.Images))
	for i, img := range userData.Images {
		images[i] = &model.Image{
			URL:    img.URL,
			Height: &img.Height,
			Width:  &img.Width,
		}
	}
	if len(images) == 0 {
		// TODO host the default image on the server / genarete image on frontend
		images = make([]*model.Image, 1)
		imageSize := 300
		images[0] = &model.Image{
			URL:    "data:image/png;base64, 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",
			Height: &imageSize,
			Width:  &imageSize,
		}
	}

	tokens := &model.Tokens{
		AccessToken:  &tokensData.AccessToken,
		TokenType:    &tokensData.TokenType,
		Scope:        &tokensData.Scope,
		ExpiresIn:    &tokensData.ExpiresIn,
		RefreshToken: &tokensData.RefreshToken,
	}

	favouriteAlbumsAlbumsMap := []*model.FavouriteAlbumEntry{
		{Key: 1, Album: nil},
		{Key: 2, Album: nil},
		{Key: 3, Album: nil},
		{Key: 4, Album: nil},
	}

	user := &model.User{
		Country:         userData.Country,
		DisplayName:     userData.DisplayName,
		Email:           userData.Email,
		ExplicitContent: explicitContent,
		ExternalUrls:    externalUrls,
		Followers:       followers,
		Href:            userData.Href,
		SpotifyID:       userData.Id,
		Images:          images,
		Product:         userData.Product,
		Type:            userData.Type,
		URI:             userData.Uri,
		Tokens:          tokens,
		Role:            `user`,
		FavouriteAlbums: favouriteAlbumsAlbumsMap,
		FollowingUsers:  make([]*string, 0),
		FollowerUsers:   make([]*string, 0),
	}

	db := database.GetDB()
	res := db.GetCollection("users").FindOne(context.TODO(), bson.M{"spotifyId": userData.Id})
	var err error
	if res.Err() == nil {
		_, err = db.GetCollection("users").UpdateOne(context.TODO(), bson.M{"spotifyId": userData.Id}, bson.M{
			"$set": bson.M{
				"country":         userData.Country,
				"displayName":     userData.DisplayName,
				"email":           userData.Email,
				"explicitContent": explicitContent,
				"externalUrls":    externalUrls,
				"followers":       followers,
				"href":            userData.Href,
				"spotifyId":       userData.Id,
				"images":          images,
				"product":         userData.Product,
				"type":            userData.Type,
				"uri":             userData.Uri,
				"tokens":          tokens, // TODO update tokens or just refresh them?
			},
		})
	} else {
		_, err = db.GetCollection("users").InsertOne(context.TODO(), user)
	}

	if err != nil {
		return nil, err
	}

	res = db.GetCollection("users").FindOne(context.TODO(), bson.M{"spotifyId": userData.Id})
	if res.Err() != nil {
		return nil, res.Err()
	}

	dbUser := model.User{}
	err = res.Decode(&dbUser)
	if err != nil {
		return nil, err
	}

	return &dbUser, nil
}

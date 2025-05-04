package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"musicboxd/database"
	"musicboxd/graph/model"
	"musicboxd/hlp"
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

	ui := hlp.UserInfo{
		ID:    *dbUser.ID,
		Email: dbUser.Email,
		Role:  dbUser.Role,
	}
	token, err := hlp.GenerateJWT(ui)
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
	}

	db := database.GetDB()
	res := db.GetCollection("users").FindOne(context.TODO(), bson.M{"spotifyId": userData.Id})
	var err error
	if res.Err() == nil {
		_, err = db.GetCollection("users").UpdateOne(context.TODO(), bson.M{"spotifyId": userData.Id}, bson.M{
			// TODO add setOnInsert ??
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

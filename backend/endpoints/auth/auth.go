package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"musicboxd/hlp"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
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
	c.SetCookie(CODE_VERIFIER_KEY, codeVerifier, 0, "/", hlp.Envs["BACKEND_DOMAIN"], true, true)
	c.SetCookie(STATE_KEY, state, 0, "/", hlp.Envs["BACKEND_DOMAIN"], true, true)
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

	c.SetCookie(CODE_VERIFIER_KEY, "", -1, "/", hlp.Envs["BACKEND_DOMAIN"], true, true)
	c.SetCookie(STATE_KEY, "", -1, "/", hlp.Envs["BACKEND_DOMAIN"], true, true)

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

	err = saveToDB(tokens, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save to db"})
		return
	}

	// TODO set JWT cookie

	c.Redirect(http.StatusFound, hlp.Envs["FRONTEND_URL"])
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

func saveToDB(tokens *tokenResponse, user *userData) error {
	// TODO check if user exists in database
	// if not, create user
	return nil
}

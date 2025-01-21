package auth

const STATE_KEY string = "spotify_auth_state"
const CODE_VERIFIER_KEY string = "pkce_code_verifier"
const AUTHORIZATION_CODE string = "authorization_code"
const JWT_KEY string = "jwt"
const USERID_KEY string = "userId"

type (
	tokenResponse struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		Scope        string `json:"scope"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	}

	userData struct {
		Country         string            `json:"country"`
		DisplayName     string            `json:"display_name"`
		Email           string            `json:"email"`
		ExplicitContent map[string]bool   `json:"explicit_content"`
		ExternalURLs    map[string]string `json:"external_urls"`
		FollowerCount   FollowerCount     `json:"followers"`
		Href            string            `json:"href"`
		Id              string            `json:"id"`
		Images          []Images          `json:"images"`
		Product         string            `json:"product"`
		Type            string            `json:"type"`
		Uri             string            `json:"uri"`
	}

	Images struct {
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	}

	FollowerCount struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
	}
)

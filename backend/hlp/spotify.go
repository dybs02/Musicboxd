package hlp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"musicboxd/graph/model"
	"net/http"
	"net/url"
)

func makeSpotifyRequest(url string, token string) ([]byte, int, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("couldn't create request: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("couldn't send request: %v", err)
	}
	defer resp.Body.Close()

	respCode := resp.StatusCode
	if respCode == 429 {
		return nil, respCode, fmt.Errorf("rate limit exceeded")
	}
	if respCode == 403 {
		return nil, respCode, fmt.Errorf("bad OAtuh")
	}
	if respCode != 200 {
		return nil, respCode, fmt.Errorf("response code: %d", respCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, respCode, err
	}

	return body, respCode, nil
}

// TODO implement searchType as enum
func SpotifySearch(query string, searchType string, accessToken string) (*model.SearchResponse, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=%s&limit=10", url.QueryEscape(query), url.QueryEscape(searchType))

	reqBody, code, err := makeSpotifyRequest(url, accessToken)
	if err != nil {
		return nil, err
	}
	if code == 401 {
		// TODO refresh token
		return nil, fmt.Errorf("unauthorized spotify request - bad token")
	}

	res := model.SearchResponse{}
	err = json.Unmarshal(reqBody, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func SpotifyGetTrack(id string, accessToken string) (*model.Track, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/tracks/%s", url.QueryEscape(id))

	reqBody, code, err := makeSpotifyRequest(url, accessToken)
	if err != nil {
		return nil, err
	}
	if code == 401 {
		// TODO refresh token
		return nil, fmt.Errorf("unauthorized spotify request - bad token")
	}

	res := model.Track{}
	err = json.Unmarshal(reqBody, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func SpotifyGetAlbum(id string, accessToken string) (*model.Album, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/albums/%s", url.QueryEscape(id))

	reqBody, code, err := makeSpotifyRequest(url, accessToken)
	if err != nil {
		return nil, err
	}
	if code == 401 {
		// TODO refresh token
		return nil, fmt.Errorf("unauthorized spotify request - bad token")
	}

	res := model.Album{}
	err = json.Unmarshal(reqBody, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func SpotifyGetAlbumTracks(url string, accessToken string) ([]*model.Track, error) {
	reqBody, code, err := makeSpotifyRequest(url, accessToken)
	if err != nil {
		return nil, err
	}
	if code == 401 {
		// TODO refresh token
		return nil, fmt.Errorf("unauthorized spotify request - bad token")
	}

	response := struct {
		Items []*model.Track `json:"items"`
	}{}
	err = json.Unmarshal(reqBody, &response)
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}

func SpotifyGetAlbumsById(ids []string, accessToken string) ([]*model.Album, error) {
	idsParam := ids[0]
	for _, id := range ids[1:] {
		idsParam += "," + id
	}

	url := fmt.Sprintf("https://api.spotify.com/v1/albums?ids=%s", idsParam)
	reqBody, code, err := makeSpotifyRequest(url, accessToken)
	if err != nil {
		return nil, err
	}
	if code == 401 {
		// TODO refresh token
		return nil, fmt.Errorf("unauthorized spotify request - bad token")
	}

	response := struct {
		Albums []*model.Album `json:"albums"`
	}{}
	err = json.Unmarshal(reqBody, &response)
	if err != nil {
		return nil, err
	}

	return response.Albums, nil
}

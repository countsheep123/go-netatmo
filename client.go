package netatmo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

const (
	authURL  = "https://api.netatmo.com/oauth2/authorize"
	tokenURL = "https://api.netatmo.com/oauth2/token"
)

type Config struct {
	Username     string
	Password     string
	ClientID     string
	ClientSecret string
	Scopes       []string
}

type Client struct {
	httpClient *http.Client
	token      string
}

func NewClient(cnf *Config) (*Client, error) {

	httpClient, token, err := auth(cnf.Username, cnf.Password, cnf.ClientID, cnf.ClientSecret, cnf.Scopes)
	if err != nil {
		return nil, err
	}

	return &Client{
		httpClient: httpClient,
		token:      token,
	}, nil
}

// https://dev.netatmo.com/resources/technical/guides/authentication/clientcredentials
// Resource Owner Password Credentials Grant
func auth(username, password, clientID, clientSecret string, scopes []string) (*http.Client, string, error) {
	oauth := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
	}

	token, err := oauth.PasswordCredentialsToken(context.TODO(), username, password)
	if err != nil {
		return nil, "", nil
	}

	httpClient := oauth.Client(context.TODO(), token)

	return httpClient, token.AccessToken, nil
}

func (c *Client) get(u url.URL, v interface{}) error {
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		if err := json.NewDecoder(res.Body).Decode(v); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("invalid request: status_code = %d", res.StatusCode)
	}
}

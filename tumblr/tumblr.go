package tumblr

import (
	"github.com/kurrik/oauth1a"
	"net/http"
	"os/exec"
)

var Signer = new(oauth1a.HmacSha1Signer)

var Service = &oauth1a.Service{
	RequestURL:   "http://www.tumblr.com/oauth/request_token",
	AuthorizeURL: "http://www.tumblr.com/oauth/authorize",
	AccessURL:    "http://www.tumblr.com/oauth/access_token",
	ClientConfig: Client,
	Signer:       Signer,
}

var Client = &oauth1a.ClientConfig{
	ConsumerKey:    "ifcG9GefmbDkhmrQwyAWGBee7IwSiNAxZnFhQh3SQOlLqbUFZI",
	ConsumerSecret: "Z2wh2xFaDstIsAFQJf5MMZEo47dofiqzF6KX1heR6ke2ZAIdIY",
	CallbackURL:    "http://localhost:51234",
}

type AccessToken struct {
	Token  string
	Secret string
	user   *oauth1a.UserConfig
}

func (a *AccessToken) User() *oauth1a.UserConfig {
	if a.user == nil {
		a.user = oauth1a.NewAuthorizedConfig(a.Token, a.Secret)
	}
	return a.user
}

func (a *AccessToken) Sign(req *http.Request) {
	Service.Sign(req, a.User())
}

func Authorize(cli *oauth1a.ClientConfig) (*AccessToken, error) {
	a := &AccessToken{
		user: &oauth1a.UserConfig{},
	}
	httpClient := new(http.Client)
	if err := a.user.GetRequestToken(Service, httpClient); err != nil {
		return nil, err
	}
	authURL, err := a.user.GetAuthorizeURL(Service)
	if err != nil {
		return nil, err
	}

	// Redirect the user to authURL and parse out a.Token and a.Verifier from the response
	if err := exec.Command("open", authURL).Run(); err != nil {
		return nil, err
	}
	return nil, nil
	/*
	if err := a.user.GetAccessToken(a.Token, a.Verifier, Service, httpClient); err != nil {
		return nil, err
	}

	return a, nil
	*/
}

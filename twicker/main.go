package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mrjones/oauth"
	"strconv"
)

var (
	flagURL =                   flag.String("u", "https://api.twitter.com/1.1/statuses/user_timeline.json", "URL")
	flagAccessTokenFile =       flag.String("a", "access-token", "Access token file")
	flagConsumerKeySecretFile = flag.String("c", "consumer-key-secret", "Consumer key and secret file")
)

type ConsumerKeySecret struct {
	Key    string
	Secret string
}

func main() {
	flag.Parse()

	var (
		err               error
		raw               []byte
		consumerKeySecret *ConsumerKeySecret
		accessToken       *oauth.AccessToken
	)

	// Consumer key and secret
	raw, err = ioutil.ReadFile(*flagConsumerKeySecretFile)
	if err != nil {
		log.Fatal(err)
	}
	consumerKeySecret = &ConsumerKeySecret{}
	if err = json.Unmarshal(raw, consumerKeySecret); err != nil {
		log.Fatal(err)
	}

	// Access tokens
	if *flagAccessTokenFile != "" {
		raw, err = ioutil.ReadFile(*flagAccessTokenFile)
		if err != nil {
			log.Fatal(err)
		}
		accessToken = &oauth.AccessToken{}
		if err = json.Unmarshal(raw, accessToken); err != nil {
			log.Fatal(err)
		}
	}

	// Make consumer client
	c := oauth.NewConsumer(
		consumerKeySecret.Key,
		consumerKeySecret.Secret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)
	c.Debug(false)

	if accessToken == nil {
		requestToken, url, err := c.GetRequestTokenAndUrl("oob")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("(1) Go to: " + url)
		fmt.Println("(2) Grant access, you should get back a verification code.")
		fmt.Print("(3) Enter that verification code here:\n--> ")

		verificationCode := ""
		fmt.Scanln(&verificationCode)

		accessToken, err = c.AuthorizeToken(requestToken, verificationCode)
		if err != nil {
			log.Fatal(err)
		}
		raw, err := json.Marshal(accessToken)
		if err != nil {
			log.Fatal(err)
		}
		println(string(raw))
	}

	// Useful work

	tweets, err := FetchTimeline(c, accessToken)
	if err != nil {
		log.Fatalf("error (%s)\n", err)
	}

	log.Printf("Received %d records\n", len(tweets))

	// Print indented
	b, err := json.MarshalIndent(tweets, "", "  ")
	if err != nil {
		log.Fatalf("error indenting (%s)\n", err)
	}
	fmt.Println(string(b))
}

func FetchTimeline(c *oauth.Consumer, accessToken *oauth.AccessToken) (tweets []interface{}, err error) {
	for i := 1; ; i++ {
		log.Printf("%d·", i)
		dt, err := FetchTimelinePage(c, accessToken, i)
		if err != nil {
			return nil, err
		}
		if len(dt) == 0 {
			break
		}
		tweets = append(tweets, dt...)
	}
	return tweets, nil
}

func FetchTimelinePage(c *oauth.Consumer, accessToken *oauth.AccessToken, page int) (tweets []interface{}, err error) {
	resp, err := c.Get(
		*flagURL,
		map[string]string{
			"skip_user":   "1",
			"count":       "100",
			"include_rts": "1",
			"page":        strconv.FormatInt(int64(page), 10),
		},
		accessToken,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	tweets = make([]interface{}, 0)
	err = json.NewDecoder(resp.Body).Decode(&tweets)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}

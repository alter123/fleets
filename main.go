package main

import (
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

const maxTweets = "200"

var (
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")
	maxTweetAge       = getenv("MAX_TWEET_AGE")
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("environ variable not found " + name)
	}
	return v
}

func getTimeline(api *anaconda.TwitterApi) ([]anaconda.Tweet, error) {
	args := url.Values{}
	args.Add("count", maxTweets)    // Define count of tweets to process
	args.Add("include_rts", "true") // Optionally include rt's also to delete
	timeline, err := api.GetUserTimeline(args)
	if err != nil {
		return make([]anaconda.Tweet, 0), err
	}
	return timeline, nil
}

func fleets(api *anaconda.TwitterApi, ageLimit time.Duration) (err error) {
	timeline, err := getTimeline(api)
	if err != nil {
		log.Println("Unable to fetch timeline ", err.Error())
		return
	}

	for _, t := range timeline {
		createdTime, err := t.CreatedAtTime()
		if err != nil {
			log.Println("could not parse time ", err.Error())
			return err
		}

		if time.Since(createdTime) > ageLimit && strings.Contains(t.Text, "#Fleet") {
			_, err := api.DeleteTweet(t.Id, true)
			log.Printf("DELETED TWEET (was %vh old): %v #%d - %s\n", time.Since(createdTime).Hours(), createdTime, t.Id, t.Text)
			if err != nil {
				log.Println("failed to delete: ", err.Error())
			}
		}
	}

	log.Println("finished processing tweets")
	return
}

func main() {
	log.SetFlags(log.LstdFlags | log.LstdFlags)

	api := anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, consumerKey, consumerSecret)
	api.SetLogger(anaconda.BasicLogger)

	h, err := time.ParseDuration(maxTweetAge)
	if err != nil {
		log.Fatalln("Unable to parse max tweet age")
		return
	}
	fleets(api, h)
}

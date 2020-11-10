# fleets

Utility to optionally delete tweets & retweets from the timeline which are older than a certain time and is tagged with `#Fleet`.

### fleets exists because

- Real-time nature of Twitter makes tweets irrelevant after a certain period of time
- Content posted years ago becomes mundane to represent your ideas


## setting up

fleets run on scheduled jobs in GitHub actions. To run your own fleet [fork this repo](https://github.com/alter123/fleets/fork) and [create a new Twitter application and generate API keys](https://apps.twitter.com/). The program assumes the following [environment variables](https://docs.github.com/en/free-pro-team@latest/actions/reference/encrypted-secrets) are set:

```sh
TWITTER_CONSUMER_KEY
TWITTER_CONSUMER_SECRET
TWITTER_ACCESS_TOKEN
TWITTER_ACCESS_TOKEN_SECRET
MAX_TWEET_AGE
```
`MAX_TWEET_AGE` expects a value of hours, such as: `MAX_TWEET_AGE = 72h`

##

Runs upon [github actions](https://github.com/features/actions) ![github actions](https://github.githubassets.com/images/modules/site/features/actions-icon-actions.svg)

# gator
boot dev blog aggregator project RSS feed

## Requirements

* Go 1.22.3+
* PostgreSQL v16.4

## How to install

Run the following command to install the go package
`go install github.com/remcous/bootdev_gator`

## Config file

Create a Config file name `.gatorconfig.json` in your home directory `~/.gatorconfig.json`

The file should have the following structure:
```
{   
"db_url":<database_connection_string>,
"current_user_name": <username>
}
```

## Valid Commands
`./gator register <username>` - Create a new user

`./gator login <username>` - Change currently active user

`./gator reset` - Reset the database

`./gator users` - Get a list of users

`./gator agg <time_between_requests>` - Aggregate the RSS Feed with a duration between requests

`./gator addfeed <feed_name> <feed_url>` - Add a new feed, automatically follow for current user

`./gator feeds` - Get a list of active feeds

`./gator follow <url>` - Follow the feed by url for the current user

`./gator following` - List followed feeds for the current user

`./gator unfollow <url>` - Unfollow feed by URL for the current user

`./gator browse <num_posts>` - Browse posts from the current users feed(s)
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/remcous/bootdev_gator/internal/database"
)

/*******************************************************************************
*	Handler
*******************************************************************************/

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("unable to parse [%s] into a time duration, %w", cmd.Args[0], err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("Unable to get next feed to update")
		return
	}
	fmt.Println("Found a feed to fetch!")

	scrapeFeed(s.db, nextFeed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(
		context.Background(),
		feed.ID,
	)
	if err != nil {
		fmt.Printf("failed to mark feed [%s] fetched, %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("Unable to fetch updated feed [%s], %w", feed.Name, err)
		return
	}

	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}
	fmt.Printf("Feed [%s] collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}

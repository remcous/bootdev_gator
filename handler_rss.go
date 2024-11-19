package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
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

	scrapeFeed(s, nextFeed)
}

func scrapeFeed(s *state, feed database.Feed) {
	_, err := s.db.MarkFeedFetched(
		context.Background(),
		feed.ID,
	)
	if err != nil {
		fmt.Printf("failed to mark feed [%s] fetched, %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("Unable to fetch updated feed [%s], %v", feed.Name, err)
		return
	}

	for _, item := range feedData.Channel.Item {
		publishedAt := sql.NullTime{}
		if pubTime, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  pubTime,
				Valid: true,
			}
		}

		_, err = s.db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				ID:        uuid.New(),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
				Title:     item.Title,
				Url:       item.Link,
				Description: sql.NullString{
					String: item.Description,
					Valid:  true,
				},
				PublishedAt: publishedAt,
				FeedID:      feed.ID,
			},
		)
		if err != nil {
			if !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				fmt.Println(err)
			}

			fmt.Printf("Couldn't create post: %v", err)
		}
	}
	fmt.Printf("Feed [%s] collected, %v posts found\n", feed.Name, len(feedData.Channel.Item))
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	var numPosts int
	var err error
	if len(cmd.Args) < 1 {
		numPosts = 2
	} else {
		numPosts, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			numPosts = 2
		}
	}

	posts, err := s.db.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  int32(numPosts),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to get posts for user [%s]", user.Name)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		printPost(post)
	}

	return nil
}

func printPost(post database.Post) {
	fmt.Printf("--- %s ---\n", post.Title)
	fmt.Printf("Published: %s\n", post.PublishedAt.Time)
	fmt.Printf("Description: %s\n", post.Description.String)
	fmt.Printf("Link: %s\n", post.Url)
	fmt.Println("==========================================")
}

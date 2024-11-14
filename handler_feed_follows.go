package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/remcous/bootdev_gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to get user [%s] from database, %w", s.cfg.CurrentUserName, err)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't find feed with url [%s], %w", cmd.Args[0], err)
	}

	feedFollow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			FeedID:    feed.ID,
			UserID:    user.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create a feed follow, %w", err)
	}

	fmt.Println("Feed follow created:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to get user [%s] from database, %w", s.cfg.CurrentUserName, err)
	}

	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("unable to get followed feeds from database, %w", err)
	}

	if len(followedFeeds) == 0 {
		fmt.Println("No feed follows found for user [%s]", currentUser.Name)
		return nil
	}

	fmt.Printf("%s followed feeds:\n", currentUser.Name)
	for _, feed := range followedFeeds {
		fmt.Printf("* %s\n", feed.FeedName)
	}

	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/remcous/bootdev_gator/internal/database"
)

/*******************************************************************************
*	Handler
*******************************************************************************/

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed to retrieve feed, %w", err)
	}

	fmt.Printf("Feed: %+v\n", feed)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	currentUserName := s.cfg.CurrentUserName
	currentUser, err := s.db.GetUser(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf("unable to access user [%s], %w", currentUserName, err)
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      feedName,
			Url:       feedUrl,
			UserID:    currentUser.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("unable to create feed, %w", err)
	}

	fmt.Printf("New feed record created: %+v\n", feed)

	return nil
}

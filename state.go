package main

import (
	"github.com/remcous/bootdev_gator/internal/config"
	"github.com/remcous/bootdev_gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

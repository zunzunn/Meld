package app

import (
	"fmt"

	"meld/internal/cli"
	"meld/internal/config"
)

// App represents the initialized application state.
type App struct {
	Config *config.Config
	CLI    *cli.RootCommand
	Error  error
}

// Init initializes the application with configuration and CLI.
func Init() (*App, error) {
	cfg, err := config.Init()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize configuration: %w", err)
	}

	rootCmd, err := cli.Init(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize CLI: %w", err)
	}

	return &App{
		Config: cfg,
		CLI:    rootCmd,
		Error:  nil,
	}, nil
}

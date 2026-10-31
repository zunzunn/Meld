package cli

import (
	"testing"

	"meld/internal/config"
)

func TestVersionString(t *testing.T) {
	cfg, err := config.Init()
	if err != nil {
		t.Fatalf("failed to init config: %v", err)
	}

	rootCmd, err := Init(cfg)
	if err != nil {
		t.Fatalf("failed to init root command: %v", err)
	}

	// Version should be set from config
	expectedVersion := cfg.Viper.GetString("version")
	if rootCmd.Version != expectedVersion {
		t.Fatalf("expected version %q, got %q", expectedVersion, rootCmd.Version)
	}
}

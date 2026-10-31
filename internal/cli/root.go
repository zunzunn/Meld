package cli

import (
	"github.com/spf13/cobra"
	"meld/internal/config"
)

// RootCommand represents the CLI root command.
type RootCommand struct {
	*cobra.Command
}

// Init initializes the Cobra root command with all subcommands.
func Init(cfg *config.Config) (*RootCommand, error) {
	cmd := &cobra.Command{
		Use:   "meld",
		Short: "MELD - a universal, local-first file utility",
		Long: `MELD is a universal, local-first file utility that combines
file conversion, compression, optimization, and other file operations
into one application.`,

		Version: cfg.Viper.GetString("version"),

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				// Delegate to the subcommand
				subCmd, _, err := cmd.Find(args)
				if err == nil && subCmd != nil {
					return subCmd.RunE(subCmd, args[1:])
				}
			}
			// Show help when no args specified
			return cmd.Help()
		},
	}

	cmd.InitDefaultVersionFlag()

	cmd.PersistentFlags().BoolP("verbose", "", false, "Enable verbose output")
	cmd.PersistentFlags().BoolP("quiet", "q", false, "Suppress non-essential output")
	cmd.PersistentFlags().BoolP("json", "j", false, "Output machine-readable JSON")

	// Add all subcommands
	cmd.AddCommand(newConvertCommand())
	cmd.AddCommand(newCompressCommand())
	cmd.AddCommand(newOptimizeCommand())
	cmd.AddCommand(newResizeCommand())
	cmd.AddCommand(newExtractCommand())
	cmd.AddCommand(newInfoCommand())
	cmd.AddCommand(newDoctorCommand())
	cmd.AddCommand(newConfigCommand())
	cmd.AddCommand(newTUICommand())
	cmd.AddCommand(newVersionCommand(cfg))
	cmd.AddCommand(newHelpCommand())

	return &RootCommand{Command: cmd}, nil
}

// newConvertCommand creates the meld convert command.
func newConvertCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "convert",
		Short: "Convert files between formats",
		Long: `MELD convert converts files between different formats.

Supported conversions will be available in future releases.

Use "meld convert --help" for more information about a command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
}

// newCompressCommand creates the meld compress command.
func newCompressCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "compress",
		Short: "Compress files to reduce size",
		Long: `MELD compress reduces file size using various compression algorithms.

Supported compression options will be available in future releases.

Use "meld compress --help" for more information about a command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
}

// newOptimizeCommand creates the meld optimize command.
func newOptimizeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "optimize",
		Short: "Optimize files for size and quality",
		Long: `MELD optimize improves file efficiency through optimization.

Available optimization techniques will be available in future releases.

Use "meld optimize --help" for more information about a command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
}

// newResizeCommand creates the meld resize command.
func newResizeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "resize",
		Short: "Resize images to specified dimensions",
		Long: `MELD resize adjusts image dimensions.

Supported image formats and resize options will be available in future releases.

Use "meld resize --help" for more information about a command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
}

// newExtractCommand creates the meld extract command.
func newExtractCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "extract",
		Short: "Extract archives to directory",
		Long: `MELD extract extracts archive files to a target directory.

Supported archive formats and extraction options will be available in future releases.

Use "meld extract --help" for more information about a command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
}

// newInfoCommand creates the meld info command.
func newInfoCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Display file or resource information",
		Long: `MELD info displays metadata and information about files or resources.

Available information fields will be available in future releases.

Use "meld info --help" for more information about a command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
}

// newDoctorCommand creates the meld doctor command.
func newDoctorCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "doctor",
		Short: "Check system health and configuration",
		Long: `MELD doctor checks the system for common issues and configuration problems.

Diagnostic results and recommendations will be available in future releases.

Use "meld doctor --help" for more information about a command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
}

// newConfigCommand creates the meld config command.
func newConfigCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Manage MELD configuration",
		Long: `MELD config manages application configuration settings.

Configuration management features will be available in future releases.

Use "meld config --help" for more information about a command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
}

// newTUICommand creates the meld tui command.
func newTUICommand() *cobra.Command {
	return &cobra.Command{
		Use:   "tui",
		Short: "Launch the TUI interface",
		Long: `MELD tui launches the keyboard-friendly text user interface.

TUI features will be available in future releases.

Use "meld tui --help" for more information about a command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
}

// newVersionCommand creates the meld version command.
func newVersionCommand(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version of meld",
		RunE: func(cmd *cobra.Command, args []string) error {
			version := cfg.Viper.GetString("version")
			cmd.Println("meld version " + version)
			return nil
		},
	}
}

// newHelpCommand creates the meld help command.
func newHelpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "help",
		Short: "Show help for meld",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				// Show help for the specified subcommand
				cmd.Help()
				return nil
			}
			return cmd.Help()
		},
	}
}
package cmd

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

// CommandRouter is the main commands router.
type CommandRouter struct {
	logger    errorLogger
	config    *Config
	formatter formatter
	extractor extractor
	rootCmd   *cobra.Command
}

func (r *CommandRouter) getFlag(
	cmd *cobra.Command, id string, options []string,
) []string {
	elements, err := cmd.Flags().GetStringSlice(id)
	if err != nil {
		panic(err)
	}
	for _, e := range elements {
		if !slices.Contains(options, e) {
			fmt.Printf("Error: invalid option %s=%s", id, e) // nolint // output
			os.Exit(0)
		}
	}

	return elements
}

// admin runs the server.
func (r *CommandRouter) extract(cmd *cobra.Command, args []string) {
	results := r.extractor.Extract(
		r.getFlag(cmd, "sources", r.config.Sources),
		r.getFlag(cmd, "contacts", r.config.Contacts),
	)

	fmt.Println(r.formatter.Format(results)) // nolint // output
}

// Run the router.
func (r *CommandRouter) Run() {
	extractCmd := &cobra.Command{
		Use:   "extract",
		Short: "extract contacts",
		Run:   r.extract,
	}
	extractCmd.Flags().StringSlice(
		"sources", r.config.Sources, "Sources of contacts.",
	)
	extractCmd.Flags().StringSlice(
		"contacts", r.config.Contacts, "Contacts to extract.",
	)
	r.rootCmd.AddCommand(extractCmd)
	err := r.rootCmd.Execute()
	if err != nil {
		r.logger.Error(errors.Wrap(err, "root command"))
	}
}

// NewCommandRouter creates a new CommandRouter.
func NewCommandRouter(
	log errorLogger, extractor extractor, formatter formatter,
) CommandRouter {
	config := NewConfig()

	return CommandRouter{
		config:    config,
		logger:    log,
		extractor: extractor,
		formatter: formatter,
		rootCmd:   &cobra.Command{Use: "contacts_extractor.app"},
	}
}

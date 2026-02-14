package command

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewModule(commands []*cobra.Command) *cobra.Command {
	rootCmd := &cobra.Command{}
	for _, c := range commands {
		rootCmd.AddCommand(c)
	}

	return rootCmd
}

func AsCommand(c any) any {
	return fx.Annotate(
		c,
		fx.ResultTags(`group:"commands"`),
	)
}

package main

import (
	"os"

	"github.com/andresmeireles/go-template/internal/core/command"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(command.NewModule),
		fx.Invoke(func(c *cobra.Command) {
			if err := c.Execute(); err != nil {
				os.Exit(1)
			}
		}),
		fx.NopLogger,
	)
}

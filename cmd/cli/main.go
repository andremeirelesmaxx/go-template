package main

import (
	"os"

	"github.com/andremeirelesmaxx/go-template/internal/core/command"
	"github.com/andremeirelesmaxx/go-template/internal/core/database"
	"github.com/andremeirelesmaxx/go-template/internal/core/env"
	"github.com/andremeirelesmaxx/go-template/internal/repositories"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			env.NewEnv,
			database.NewConn,
			database.NewRepository,
			database.NewMigrate,
			repositories.NewExample,
			command.AsCommand(command.NewAppModeCommand),
			command.AsCommand(command.NewAppRootPathCommand),
			command.AsCommand(command.NewMigrationUp),
			fx.Annotate(
				command.NewModule,
				fx.ParamTags(`group:"commands"`),
			),
		),
		fx.Invoke(func(c *cobra.Command) {
			if err := c.Execute(); err != nil {
				os.Exit(1)
			}
		}),
		// fx.NopLogger,
	)
}

package command

import (
	"fmt"

	"github.com/andremeirelesmaxx/go-template/internal/core/env"
	"github.com/spf13/cobra"
)

func NewAppRootPathCommand(env env.Env) *cobra.Command {
	return &cobra.Command{
		Use:   "app:root",
		Short: "show app root path",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println(env.App.Root)
		},
	}
}

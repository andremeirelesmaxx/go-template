package command

import (
	"fmt"

	"github.com/andremeirelesmaxx/go-template/internal/core/env"
	"github.com/spf13/cobra"
)

func NewAppModeCommand(env env.Env) *cobra.Command {
	return &cobra.Command{
		Use:   "app:mode",
		Short: "show app mode",
		Run: func(c *cobra.Command, args []string) {
			if env.App.IsLocal {
				fmt.Println("on dev/local mode")

				return
			}

			fmt.Println("on prod mode")
		},
	}
}

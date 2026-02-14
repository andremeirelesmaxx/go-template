package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewModule() *cobra.Command {
	a := &cobra.Command{
		Use: "boules",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("boules")
		},
	}

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(a)

	return rootCmd
}

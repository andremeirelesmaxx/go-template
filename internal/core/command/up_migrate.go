package command

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
)

func NewMigrationUp(m *migrate.Migrate) *cobra.Command {
	return &cobra.Command{
		Use: "migrate:up",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("start up migration")
			err := m.Up()

			if err != nil {
				fmt.Println("error on migration runner", err.Error())

				return
			}

			fmt.Println("up migration ran with success")
		},
	}
}

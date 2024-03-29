package deletecmd

import (
	"github.com/khulnasoft/orchard/pkg/client"
	"github.com/spf13/cobra"
)

func newDeleteServiceComandCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "service-account NAME",
		Short: "Delete a service account",
		Args:  cobra.ExactArgs(1),
		RunE:  runDeleteServiceAccountCommand,
	}
}

func runDeleteServiceAccountCommand(cmd *cobra.Command, args []string) error {
	name := args[0]

	client, err := client.New()
	if err != nil {
		return err
	}

	return client.ServiceAccounts().Delete(cmd.Context(), name, false)
}

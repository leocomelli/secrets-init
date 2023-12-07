package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type rootCmd struct {
	cmd *cobra.Command
}

func (r *rootCmd) execute() {
	if err := r.cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRootCmd() *rootCmd {
	root := &rootCmd{}
	cmd := &cobra.Command{
		Use:   "secret-init",
		Short: "Read external secrets from some providers",
		Long: `
This is a simple CLI that reads secrets from Secrets Manager, like:
  - AWS
  - GCP
It's a perfect "init" container in Kubernetes.
it can create a file on a shared volume so the other containers can use that file.
secrets-init can filter one or more secrets by name using a regular expression.
it also parses the secret content as plain text or json.
		`,
		SilenceUsage:      true,
		SilenceErrors:     true,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}

	cmd.AddCommand(
		newSyncCmd().cmd,
		newVersionCmd().cmd,
	)
	root.cmd = cmd
	return root
}

func Execute() { newRootCmd().execute() }

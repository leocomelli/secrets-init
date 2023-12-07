package cmd

import (
	"fmt"
	"github.com/leocomelli/secrets-init/internal/secrets"
	"github.com/leocomelli/secrets-init/pkg/provider/common"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

type syncCmd struct {
	cmd *cobra.Command
}

func (r *syncCmd) execute() {
	if err := r.cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newSyncCmd(data *data) *syncCmd {
	sync := &syncCmd{}
	secretOpts := &common.SecretsOpts{}

	cmd := &cobra.Command{
		Version:           data.version,
		Use:               "sync",
		Aliases:           []string{"s"},
		Short:             "Sync external secrets to a container init",
		SilenceUsage:      true,
		SilenceErrors:     true,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {

			fetch, err := secrets.New(secretOpts)
			if err != nil {
				common.Logger.Fatal("error getting the secrets", zap.Error(err))
				return err
			}

			return fetch.Secrets()
		},
	}

	sync.cmd = cmd

	sync.cmd.Flags().StringVarP(&secretOpts.Provider, "provider", "e", "gcp", "name of the provider that manages the secrets")
	sync.cmd.Flags().StringVarP(&secretOpts.AssumeRole, "assume-role", "a", "", "role to assume when using aws provider")
	sync.cmd.Flags().StringVarP(&secretOpts.Project, "project", "p", "", "gcp project that contains the secrets")
	sync.cmd.Flags().StringVarP(&secretOpts.Filter, "filter", "f", "", "regex to filter secrets by name")
	sync.cmd.Flags().StringVarP(&secretOpts.Parser, "data-parser", "d", "plaintext", "parse secret based on data type")
	sync.cmd.Flags().StringVarP(&secretOpts.Template, "template", "t", "", "template to render secret data")
	sync.cmd.Flags().StringVarP(&secretOpts.Output, "output", "o", "", "path to write output file to")

	return sync
}

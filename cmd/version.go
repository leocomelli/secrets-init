package cmd

import (
	"fmt"
	"github.com/leocomelli/secrets-init/version"
	"github.com/spf13/cobra"
	"os"
)

type versionCmd struct {
	cmd *cobra.Command
}

func (r *versionCmd) execute() {
	if err := r.cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newVersionCmd() *versionCmd {
	sync := &versionCmd{}
	cmd := &cobra.Command{
		Use:               "version",
		Aliases:           []string{"v"},
		Short:             "Return the current version of secret init",
		SilenceUsage:      true,
		SilenceErrors:     true,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Build Date:", version.BuildDate)
			fmt.Println("Git Commit:", version.GitCommit)
			fmt.Println("Version:", version.Version)
		},
	}

	sync.cmd = cmd
	return sync
}

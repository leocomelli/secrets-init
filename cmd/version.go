package cmd

import (
	"fmt"
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

func newVersionCmd(data *data) *versionCmd {
	sync := &versionCmd{}
	cmd := &cobra.Command{
		Version: data.version,
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Return the current version of secret init",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Build Date:", data.date)
			fmt.Println("Git Commit:", data.commit)
			fmt.Println("Version:", data.version)
		},
	}

	sync.cmd = cmd
	return sync
}

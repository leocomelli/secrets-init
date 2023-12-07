package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	// The git commit that was compiled. These will be filled in by the compiler.
	GitCommit string

	// The main version number that is being run at the moment.
	Version string

	// BuildDate contains the date and time of build process.
	BuildDate string
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
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Return the current version of secret init",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Build Date:", BuildDate)
			fmt.Println("Git Commit:", GitCommit)
			fmt.Println("Version:", Version)
		},
	}

	sync.cmd = cmd
	return sync
}

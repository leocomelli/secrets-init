package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

// GitCommit returns the git commit that was compiled. This will be filled in by the compiler.
var GitCommit string

// Version returns the main version number that is being run at the moment.
const Version = "0.1.0"

// BuildDate returns the date the binary was built
var BuildDate = ""

// GoVersion returns the version of the go runtime used to compile the binary
var GoVersion = runtime.Version()

// OsArch returns the os and arch used to build the binary
var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)

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
			fmt.Println("Build Date:", BuildDate)
			fmt.Println("Git Commit:", GitCommit)
			fmt.Println("Version:", Version)
			fmt.Println("Go Version:", GoVersion)
			fmt.Println("OS / Arch:", OsArch)
		},
	}

	sync.cmd = cmd
	return sync
}

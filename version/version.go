package version

var (
	// The git commit that was compiled. These will be filled in by the compiler.
	GitCommit string

	// The main version number that is being run at the moment.
	Version string

	// BuildDate contains the date and time of build process.
	BuildDate string
)

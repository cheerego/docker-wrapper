package cmd

import (
	"fmt"
	"runtime"
)

var (
	GitTag       string
	GitCommit    string
	GitTreeState string
	BuildDate    string
	GoVersion    string = runtime.Version()
	Compiler     string = runtime.Compiler
	Platform     string = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

package main

import (
	"os"

	"github.com/khulnasoft-lab/labeler/cmd"
	"github.com/khulnasoft-lab/labeler/logs"
	"github.com/tonglil/versioning"
)

var version string

func init() {
	versioning.Set(version)

	logs.Output = os.Stdout
}

func main() {
	cmd.Execute()
}

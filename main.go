package main

import (
	"os"
	// Import the pulse plugin library
	"github.com/intelsdi-x/pulse/control/plugin"
	// Import our collector plugin implementation
	"github.com/intelsdi-x/pulse-plugin-collector-psutil/psutil"
)

// plugin bootstrap
func main() {
	plugin.Start(
		psutil.Meta(),
		psutil.NewPsutilCollector(), // CollectorPlugin interface
		os.Args[1],
	)
}

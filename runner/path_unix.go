// +build linux freebsd netbsd openbsd

package runner

import "os/exec"

const (
	// DefaultChromePath is the default path to use for Chrome if the
	// executable is not in $PATH.
	DefaultChromePath = "/usr/bin/google-chrome"
)

// DefaultChromeNames are the default Chrome executable names to look for in
// $PATH.
var DefaultChromeNames = []string{
	"google-chrome-stable",
	"google-chrome",
	"chromium-browser",
	"chromium",
	"google-chrome-beta",
	"google-chrome-unstable",
}

func findChromePath() string {
	for _, p := range chromeNames {
		path, err := exec.LookPath(p)
		if err == nil {
			return path
		}
	}

	return DefaultChromePath
}

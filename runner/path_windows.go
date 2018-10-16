// +build windows

package runner

import "os/exec"

const (
	// DefaultChromePath is the default path to use for Google Chrome if the
	// executable is not in %PATH%.
	DefaultChromePath = `C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`
	DefaultCanaryPath = `C:\Users\TEST\AppData\Local\Google\Chrome SxS\Application\chrome.exe`

	// DefaultEdgeDiagnosticsAdapterPath is the default path to use for the
	// Microsoft Edge Diagnostics Adapter if the executable is not in %PATH%.
	DefaultEdgeDiagnosticsAdapterPath = `c:\Edge\EdgeDiagnosticsAdapter\x64\EdgeDiagnosticsAdapter.exe`
)

// DefaultChromeNames are the default Chrome executable names to look for in
// $PATH.
var DefaultChromeNames = []string{
	"google-chrome-stable",
	"google-chrome",
	"chrome",
	"chromium-browser",
	"chromium",
	"google-chrome-beta",
	"google-chrome-unstable",
}

func findChromePath() string {
	//path, err := exec.LookPath(`Application\chrome.exe`)
	//if err == nil {
	//	return path
	//}

	return DefaultCanaryPath
}

func findEdgePath() string {
	path, err := exec.LookPath(`EdgeDiagnosticsAdapter.exe`)
	if err == nil {
		return path
	}

	return DefaultEdgeDiagnosticsAdapterPath
}

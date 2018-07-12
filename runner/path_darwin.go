// +build darwin

package runner

const (
	// DefaultChromePath is the default path to the Chrome application.
	DefaultChromePath = `/Applications/Google Chrome Canary.app/Contents/MacOS/Google Chrome Canary`
)

func findChromePath() string {
	return DefaultChromePath
}

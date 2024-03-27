package main

// Constants, global variables and their initializing funcs.

import (
	"os"
	"path"
	"time"
)

const projectName = "cleanupca"

// logoutTimeout defines the time span an imap session will be alive.
// cf. [here](https://pkg.go.dev/github.com/mxk/go-imap/imap#Client.Logout)
const logoutTimeout = 1000 * time.Second

// variables initialized by initConstants func, see below
var (
	xdgConfigDir string // Standard config dir according to XDG.     Most likely ~/.config
	caFile       string // Name of file containing ca certificates.  Most likely ~/.config/ca.pem
	cabakFile    string // Name of backup file.                      Most likely ~/.config/ca.pem.bak
)

// initConstants initialize aforementioned global variables
// Call this function before any other action of goiko's source code!
func initConstants() (err error) {
	xdgConfigDir, err := os.UserConfigDir()
	if err != nil {
		return
	}

	caFile = path.Join(xdgConfigDir, "ca.pem")
	cabakFile = path.Join(xdgConfigDir, "ca.pem.bak")

	return
}
